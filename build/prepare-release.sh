#!/bin/bash
# Copyright 2025 SolarWinds Worldwide, LLC. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

if [ "$#" -lt 1 ]; then
    echo "Usage: $0 <version> $1 <otel_version>"
    exit 1
fi

VERSION=$1
OTEL_VERSION=$2

# Update CHANGELOG.md
CHANGELOG_FILE="./CHANGELOG.md"
if [ ! -f "$CHANGELOG_FILE" ]; then
    echo "CHANGELOG.md not found!"
    exit 1
fi
if ! grep -q "## v$VERSION" "$CHANGELOG_FILE"; then
    perl -pi -e "s/^## vNext/## vNext\n\n## v$VERSION/" "$CHANGELOG_FILE"
    echo "CHANGELOG.md updated with version v$VERSION"
else
    echo "CHANGELOG.md already contains 'v$VERSION', no update made."
fi

# Update release manifest files
ALL_MANIFEST_YAML=$(find . -name "manifest.yaml" -type f | sort)
# Update distribution manifest yaml versions.
for f in $ALL_MANIFEST_YAML; do
    perl -pi -e "s/version: \d+\.\d+\.\d+/version: $VERSION/g" "$f"
    echo "Updated version in distribution yaml \`$f\` with version v$VERSION"
done

if [ -z "$OTEL_VERSION" ]; then
  echo "OTEL_VERSION not set, skipping."
else
  # update swi_contrib_version in Makefile
    MAKEFILE="./Makefile"
    if [ ! -f "$MAKEFILE" ]; then
        echo "Makefile not found!"
        exit 1
    fi
    perl -pi -e "s|otel_version := \d+\.\d+\.\d+|otel_version := $OTEL_VERSION|g" "$MAKEFILE"
    echo "Updated otel_version in Makefile to version $OTEL_VERSION"

  # update otel contrib references in distribution yaml files
  for f in $ALL_MANIFEST_YAML; do
      perl -pi -e "s|^(\s+- gomod: github.com/open-telemetry/opentelemetry-collector-contrib/[^ ]*) v[0-9]+\.[0-9]+\.[0-9]+$|\1 v$OTEL_VERSION|" "$f"
      echo "References to 'github.com/open-telemetry/opentelemetry-collector-contrib' in $f updated with version v$OTEL_VERSION"
      perl -pi -e "s|^(\s+- gomod: go.opentelemetry.io/[^ ]*) v[0-9]+\.[0-9]+\.[0-9]+$|\1 v$OTEL_VERSION|" "$f"
      echo "References to 'go.opentelemetry.io' in $f updated with version v$OTEL_VERSION"
  done
fi

# update solarwinds contrib references in distribution yaml files
for f in $ALL_MANIFEST_YAML; do
    perl -pi -e "s|^(\s+- gomod: github.com/solarwinds/solarwinds-otel-collector-contrib/[^ ]*) v[0-9]+\.[0-9]+\.[0-9]+$|\1 v$VERSION|" "$f"
    echo "References to 'github.com/solarwinds/solarwinds-otel-collector-contrib' in $f updated with version v$VERSION"
done

# We need to run go mod tidy after raising versions of solarwinds-otel-collector-contrib components
echo "Running go mod tidy"
find . -name "go.mod" -execdir sh -c 'go mod tidy' \;

# update pkg\version\version.go to set the actual release version
GO_VERSION_FILE="./pkg/version/version.go"
if [ ! -f "$GO_VERSION_FILE" ]; then
    echo "version.go not found!"
    exit 1
fi
perl -pi -e "s|^(const Version =) \"[0-9]+\.[0-9]+\.[0-9]+\"$|\1 \"$VERSION\"|" "$GO_VERSION_FILE"
echo "Version.go updated with version $VERSION"
