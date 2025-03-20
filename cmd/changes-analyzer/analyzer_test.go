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

// analyzer_test.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

type mockTransport struct {
	responses map[string]*http.Response
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()
	if resp, ok := t.responses[url]; ok {
		return resp, nil // No logging here to avoid consuming body
	}
	return nil, fmt.Errorf("no mock response for URL: %s", url)
}

func TestGetMessage(t *testing.T) {
	releasesURL := "https://api.github.com/repos/open-telemetry/opentelemetry-collector-contrib/releases?per_page=100"
	releasesBody := `[{"tag_name":"v0.122.0","prerelease":false}]`
	releaseNotesURL := "https://github.com/open-telemetry/opentelemetry-collector-contrib/releases/tag/v0.122.0"
	nextReleasesURL := "https://api.github.com/repos/open-telemetry/opentelemetry-collector-contrib/releases?per_page=100&page=2"
	releaseNotesBody := `
        <h2>End user changelog</h2>
        <h3>🛑 Breaking changes 🛑</h3>
        <ul>
            <li>elasticsearchexporter: Dynamically route documents by default unless {logs,metrics,traces}_index is non-empty (#38361)
Overhaul in document routing.</li>
        </ul>
    `

	linkHeader := fmt.Sprintf(`<%s>; rel="next"`, nextReleasesURL)
	mockResponses := map[string]*http.Response{
		releasesURL: {
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(releasesBody)),
			Header: http.Header{
				"Content-Type": []string{"application/json"},
				"Link":         []string{linkHeader},
			},
		},
		releaseNotesURL: {
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(releaseNotesBody)),
			Header:     http.Header{"Content-Type": []string{"text/html"}},
		},
	}

	originalTransport := client.Transport
	client.Transport = &mockTransport{responses: mockResponses}
	defer func() { client.Transport = originalTransport }()

	oldTag := "v0.121.0"
	newTag := "v0.122.0"
	componentsOfInterest := []string{"elasticsearchexporter"}
	repo := "opentelemetry-collector-contrib"
	encode := false

	message, err := getMessage(oldTag, newTag, componentsOfInterest, repo, encode)
	if err != nil {
		t.Fatalf("getMessage failed: %v", err)
	}

	expected := `# OPENTELEMETRY-COLLECTOR-CONTRIB CHANGES
**Diff**: [v0.121.0 to v0.122.0](https://github.com/open-telemetry/opentelemetry-collector-contrib/compare/v0.121.0...v0.122.0)

#### elasticsearchexporter
- **Breaking Changes**:
  - 0.122.0: elasticsearchexporter: Dynamically route documents by default unless {logs,metrics,traces}_index is non-empty ([#38361](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/38361))
             Overhaul in document routing.


`
	if message != expected {
		t.Errorf("getMessage returned unexpected result:\nGot:\n'%s'\nExpected:\n'%s'", message, expected)
	}
}

// TestGetComponentsFromGoMod tests the getComponentsFromGoMod function with various scenarios.
func TestGetComponentsFromGoMod(t *testing.T) {
	tests := []struct {
		name             string
		goModContent     string
		dependencyFilter string
		want             string
		wantErr          bool
		errMsg           string
	}{
		{
			name: "basic case with single dependency",
			goModContent: `
                module example.com/myapp
                require github.com/example/mylib v1.0.0
            `,
			dependencyFilter: "github.com/example",
			want:             "mylib",
			wantErr:          false,
		},
		{
			name: "multiple dependencies with filter",
			goModContent: `
                module example.com/myapp
                require github.com/example/mylib v1.0.0
                require github.com/example/anotherlib v2.0.0
                require github.com/other/ignoreme v1.0.0
            `,
			dependencyFilter: "github.com/example",
			want:             "mylib,anotherlib",
			wantErr:          false,
		},
		{
			name: "ignore indirect dependencies",
			goModContent: `
                module example.com/myapp
                require github.com/example/mylib v1.0.0
                require github.com/example/anotherlib v2.0.0 // indirect
            `,
			dependencyFilter: "github.com/example",
			want:             "mylib",
			wantErr:          false,
		},
		{
			name: "no matching dependencies",
			goModContent: `
                module example.com/myapp
                require github.com/other/ignoreme v1.0.0
            `,
			dependencyFilter: "github.com/example",
			want:             "",
			wantErr:          false,
		},
		{
			name:             "empty file",
			goModContent:     "",
			dependencyFilter: "github.com/example",
			want:             "",
			wantErr:          false,
		},
		{
			name:             "file not found",
			goModContent:     "", // Irrelevant since file won’t exist
			dependencyFilter: "github.com/example",
			want:             "",
			wantErr:          true,
			errMsg:           "failed to open go.mod file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for all cases except "file not found"
			var goModPath string
			if tt.name == "file not found" {
				// For "file not found", use a nonexistent path
				goModPath = "/nonexistent/path/to/go.mod"
			} else {
				tmpFile, err := ioutil.TempFile("", "test-gomod-*.mod")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				defer os.Remove(tmpFile.Name()) // Clean up after test
				defer tmpFile.Close()

				// Write the go.mod content to the temp file
				if _, err := tmpFile.WriteString(tt.goModContent); err != nil {
					t.Fatalf("failed to write to temp file: %v", err)
				}
				goModPath = tmpFile.Name()
			}

			// Run the function
			got, err := getComponentsFromGoMod(goModPath, tt.dependencyFilter)

			// Check error conditions
			if tt.wantErr {
				if err == nil {
					t.Errorf("getComponentsFromGoMod() error = nil, but we expected error containing %q", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("getComponentsFromGoMod() error = %v, but we expected error containing %q", err, tt.errMsg)
				}
				return
			}
			if err != nil {
				t.Errorf("getComponentsFromGoMod() error = %v, want nil", err)
				return
			}

			// Check result
			if got != tt.want {
				t.Errorf("getComponentsFromGoMod() returned %q, but we expected %q", got, tt.want)
			}
		})
	}
}
