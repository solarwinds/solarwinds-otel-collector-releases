// Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/go-version"
)

type categoryToChangesMap = map[string][]string

const breakingChanges = "breaking_changes"
const deprecations = "deprecations"
const enhancements = "enhancements"

// parseVersion parses a version string by stripping the 'v' prefix and creating a version object.
func parseVersion(verStr string) (*version.Version, error) {
	verStr = strings.TrimPrefix(verStr, "v")
	return version.NewVersion(verStr)
}

// parseLinkHeader extracts pagination URLs from the GitHub API Link header.
func parseLinkHeader(header string) map[string]string {
	links := make(map[string]string)
	parts := strings.Split(header, ",")
	for _, part := range parts {
		sections := strings.Split(part, ";")
		if len(sections) < 2 {
			continue
		}
		url := strings.Trim(sections[0], "<> ")
		rel := strings.Trim(strings.TrimPrefix(sections[1], " rel="), "\"")
		links[rel] = url
	}
	return links
}

// getVersionsBetween retrieves all released versions between oldVersion and newVersion from GitHub.
func getVersionsBetween(oldVersion, newVersion string, opentelemetryRepo string) ([]*version.Version, error) {
	owner := "open-telemetry"
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases?per_page=100", owner, opentelemetryRepo)

	var allReleases []string
	for url != "" {
		response, err := getResponse(url)
		if err != nil {
			return nil, fmt.Errorf("get request failed for url %s: %v", url, err)
		}
		defer response.Body.Close()

		var releases []struct {
			TagName    string `json:"tag_name"`
			Prerelease bool   `json:"prerelease"`
		}
		if err := json.NewDecoder(response.Body).Decode(&releases); err != nil {
			return nil, fmt.Errorf("failed to decode releases: %v", err)
		}

		for _, rel := range releases {
			if !rel.Prerelease {
				allReleases = append(allReleases, rel.TagName)
			}
		}

		linkHeader := response.Header.Get("Link")
		if linkHeader != "" {
			links := parseLinkHeader(linkHeader)
			url = links["next"]
			break
		}
	}

	// Parse the boundary versions
	oldVer, err := parseVersion(oldVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid old version %s: %v", oldVersion, err)
	}
	newVer, err := parseVersion(newVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid new version %s: %v", newVersion, err)
	}

	var filtered []*version.Version
	for _, tag := range allReleases {
		ver, err := version.NewVersion(tag)
		if err != nil {
			return nil, fmt.Errorf("failed to parse version %s: %v", tag, err)
		}
		if ver.GreaterThanOrEqual(oldVer) && ver.LessThanOrEqual(newVer) {
			filtered = append(filtered, ver)
		}
	}

	// Sort versions in ascending order
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Compare(filtered[j]) < 0
	})

	return filtered, nil
}

// getResponse calls http.Get for given url and returns the response
func getResponse(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET response for url:  %s: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request status is %d for url %s", resp.StatusCode, url)
	}
	return resp, err
}

// fetchReleaseNotes retrieves the HTML content of release notes for a specific version.
func fetchReleaseNotes(version string, opentelemetryRepo string) (string, error) {
	url := fmt.Sprintf("https://github.com/open-telemetry/%s/releases/tag/v%s", opentelemetryRepo, version)
	response, err := getResponse(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read release notes body: %v", err)
	}
	return string(bodyBytes), nil
}

// extractReleaseSections extracts specified sections (e.g., Breaking changes, Deprecations, Enhancements) from HTML content.
func extractReleaseSections(htmlContent string) (map[string][]string, error) {
	// Define sections to extract and their corresponding categories
	sectionPhrases := map[string]string{
		"Breaking changes": breakingChanges,
		"Deprecations":     deprecations,
		"Enhancements":     enhancements,
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Printf("failed to parse HTML: %v", err)
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	sectionMap := make(map[string][]string)
	doc.Find("h1, h2, h3").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		// Skip the changelog type headers
		if strings.Contains(text, "End user changelog") || strings.Contains(text, "API changelog") {
			return // Continue to next element
		}

		// Process section headers regardless of which changelog they belong to
		for phrase, category := range sectionPhrases {
			if strings.Contains(text, phrase) {
				var changes []string
				node := s.Next() // Get the next sibling after <h3>
				if node.Length() > 0 && node.Is("ul") {
					for child := node.Children().First(); child.Length() > 0; child = child.Next() {
						if !child.Is("li") {
							continue
						}
						changeText := strings.TrimSpace(child.Text())
						if changeText == "" {
							continue
						}
						reformated := strings.ReplaceAll(changeText, "\n", "\n             ")
						changes = append(changes, reformated)
					}
				}
				sectionMap[category] = append(sectionMap[category], changes...)
			}
		}
	})
	return sectionMap, nil
}

// getComponentChanges retrieves breaking changes, deprecations, and enhancements for specified components across versions.
func getComponentChanges(versionOld, versionNew string, componentsOfInterest []string, opentelemetryRepo string) (map[string]categoryToChangesMap, error) {
	versions, err := getVersionsBetween(versionOld, versionNew, opentelemetryRepo)
	if err != nil {
		return nil, fmt.Errorf("failed to get versions: %v", err)
	}

	// Fetch release notes for each version
	releaseNotes := make(map[string]map[string][]string)
	for _, ver := range versions {
		htmlContent, err := fetchReleaseNotes(ver.String(), opentelemetryRepo)
		if err != nil {
			fmt.Printf("Skipping %s due to fetch error: %v\n", ver, err)
			continue
		}
		sectionChanges, err := extractReleaseSections(htmlContent)
		if err != nil {
			fmt.Printf("Skipping %s due to parsing error: %v\n", ver, err)
			continue
		}
		releaseNotes[ver.String()] = sectionChanges
	}

	// Initialize the component changes map
	componentChanges := make(map[string]categoryToChangesMap)
	for _, component := range componentsOfInterest {
		componentChanges[component] = categoryToChangesMap{
			breakingChanges: {},
			deprecations:    {},
			enhancements:    {},
		}
	}
	// Now filter only those changes that happened on components we care about
	for ver, sectionChanges := range releaseNotes {
		for category, changes := range sectionChanges {
			for _, change := range changes {
				for _, component := range componentsOfInterest {
					// Line has to contain 'component_name:'
					if strings.Contains(change, fmt.Sprintf("%s:", component)) {
						componentChanges[component][category] = append(
							componentChanges[component][category],
							fmt.Sprintf("%s: %s", ver, change),
						)
					}
				}
			}
		}
	}
	// Filter out components with no changes in any category
	result := make(map[string]categoryToChangesMap)
	for component, categories := range componentChanges {
		hasChanges := false
		for _, changes := range categories {
			if len(changes) > 0 {
				hasChanges = true
				break
			}
		}
		if hasChanges {
			result[component] = categories
		}
	}

	// Sort all categories for each component
	for _, changes := range componentChanges {
		sort.Strings(changes[breakingChanges])
		sort.Strings(changes[deprecations])
		sort.Strings(changes[enhancements])
	}

	return result, nil
}

// Regular expression to match words with dots (e.g., feature gates like receiver.prometheusreceiver.UseCollectorStartTimeFallback)
var codePattern = regexp.MustCompile(`\b[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+\b`)

// formatDescription formats a change description by wrapping words with dots in backticks for code formatting.
func formatDescription(desc string) string {
	return codePattern.ReplaceAllStringFunc(desc, func(match string) string {
		return "`" + match + "`"
	})
}

// formatComponentChanges formats the component changes into a Markdown string suitable for GitHub comments, skipping empty categories.
func formatComponentChanges(opentelemetryRepo string, componentChanges map[string]categoryToChangesMap) string {
	var blocks []string
	components := make([]string, 0, len(componentChanges))
	for component := range componentChanges {
		components = append(components, component)
	}
	sort.Strings(components) // Sort components alphabetically

	prPattern := regexp.MustCompile(`#(\d+)`) // Matches # followed by digits

	for _, component := range components {
		var componentBlock strings.Builder
		componentBlock.WriteString(fmt.Sprintf("#### %s\n", component))

		categories := componentChanges[component]
		for _, category := range []string{breakingChanges, deprecations, enhancements} {
			changes := categories[category]
			if len(changes) > 0 { // Only include categories with changes
				display := strings.Title(strings.ReplaceAll(category, "_", " "))
				componentBlock.WriteString(fmt.Sprintf("- **%s**:\n", display))
				for _, change := range changes {
					// v0.119.0: cumulativetodeltaprocessor: Add metric type filter for cumulativetodelta processor (#33673)
					parts := strings.SplitN(change, ": ", 2)
					if len(parts) == 2 {
						version, desc := parts[0], parts[1]
						// Format description with code and PR links
						formattedDesc := formatDescription(desc)
						formattedDesc = prPattern.ReplaceAllStringFunc(formattedDesc, func(match string) string {
							prNum := strings.TrimPrefix(match, "#")
							return fmt.Sprintf("[#%s](https://github.com/open-telemetry/%s/pull/%s)", prNum, opentelemetryRepo, prNum)
						})
						componentBlock.WriteString(fmt.Sprintf("  - %s: %s\n", version, formattedDesc))
					} else {
						// Fallback if splitting fails
						componentBlock.WriteString(fmt.Sprintf("  - %s\n", change))
					}
				}
			}
		}
		// Only append the block if it has content beyond the component header
		if componentBlock.Len() > len(fmt.Sprintf("#### %s\n", component)) {
			blocks = append(blocks, componentBlock.String())
		}
	}
	// separator between components
	return strings.Join(blocks, "\n---\n")
}

// getMessage generates a formatted github formated message listing component changes between two versions. Optionally, encodes to base64.
func getMessage(oldTag, newTag string, componentsOfInterest []string, opentelemetryRepo string, encode bool) (string, error) {
	componentChanges, err := getComponentChanges(oldTag, newTag, componentsOfInterest, opentelemetryRepo)
	if err != nil {
		fmt.Printf("failed to get component changes: %v", err)
		return "", fmt.Errorf("failed to get component changes: %v", err)
	}
	compareURL := fmt.Sprintf("https://github.com/open-telemetry/%s/compare/v%s...v%s", opentelemetryRepo, oldTag, newTag)
	// Build the Markdown output
	markdown := strings.ToUpper(fmt.Sprintf("# %s changes\n", opentelemetryRepo))
	markdown += fmt.Sprintf("**Diff**: [%s to %s](%s)\n\n", oldTag, newTag, compareURL)
	markdown += formatComponentChanges(opentelemetryRepo, componentChanges)
	markdown += "\n\n"

	if encode {
		return base64.StdEncoding.EncodeToString([]byte(markdown)), nil
	}
	return markdown, nil
}

// getComponentsFromGoMod reads the go.mod file, filters lines containing dependencyFilter,
// extracts the component names, and returns them as a comma-separated string.
func getComponentsFromGoMod(goModPath, dependencyFilter string) (string, error) {
	file, err := os.Open(goModPath)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod file: %v", err)
	}
	defer file.Close()
	var components []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line contains the dependency filter, ignoring indirect dependencies
		if strings.Contains(line, dependencyFilter) && !strings.Contains(line, "indirect") {
			parts := strings.Fields(line)
			if len(parts) >= 2 { // Ensure there’s at least a path and version
				// Extract the component name from the path (last part after "/")
				pathParts := strings.Split(parts[0], "/")
				if len(pathParts) > 0 {
					component := pathParts[len(pathParts)-1]
					components = append(components, component)
				}
			}
		}
	}
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod file: %v", err)
	}
	// Return components as a comma-separated string
	return strings.Join(components, ","), nil
}

// Example: go run ./main.go --old v0.119.0 --new v0.121.0 --goModPath ./../../../cmd/solarwinds-otel-collector/go.mod --dependencyFilter opentelemetry-collector-contrib
func main() {
	var oldTag, newTag, componentsStr, repo, goModPath, dependencyFilter string
	var encode bool
	flag.StringVar(&oldTag, "old", "", "Old version tag (e.g., v0.119.0)")
	flag.StringVar(&newTag, "new", "", "New version tag (e.g., v0.121.0)")
	flag.StringVar(&componentsStr, "components", "", "Comma-separated list of components (e.g., prometheusreceiver,awss3exporter)")
	flag.StringVar(&repo, "repo", "", "GitHub repository name")
	flag.StringVar(&goModPath, "goModPath", "", "Path to the go.mod file (e.g., /app/go.mod)")
	flag.StringVar(&dependencyFilter, "dependencyFilter", "", "Filter for dependencies in go.mod (e.g., open-telemetry-contrib)")
	flag.BoolVar(&encode, "encode", false, "Whether to base64 encode the output")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if oldTag == "" || newTag == "" {
		fmt.Println("Error: old tag and new tag are required.")
		flag.Usage()
		os.Exit(1)
	}
	if repo == "" {
		fmt.Println("Error: repo is required.")
		flag.Usage()
		os.Exit(1)
	}

	var componentsOfInterest []string
	if componentsStr != "" {
		// Use provided components if available
		componentsOfInterest = strings.Split(componentsStr, ",")
	} else if goModPath != "" && dependencyFilter != "" {
		// Or extract components from go.mod
		components, err := getComponentsFromGoMod(goModPath, dependencyFilter)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		componentsOfInterest = strings.Split(components, ",")
	} else {
		fmt.Println("Error: Either components or both goModPath and dependencyFilter are required.")
		flag.Usage()
		os.Exit(1)
	}

	message, err := getMessage(oldTag, newTag, componentsOfInterest, repo, encode)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(message)
}
