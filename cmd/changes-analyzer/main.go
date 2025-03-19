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
	"flag"
	"fmt"
	"os"
	"strings"
)

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
