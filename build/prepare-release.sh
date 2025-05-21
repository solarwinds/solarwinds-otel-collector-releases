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
    echo "Usage: $0 <version> $1 <swi_contrib_version> $2 <builder_version>"
    exit 1
fi

VERSION=$1
CONTRIB_VERSION=$2
BUILDER_VERSION=$3

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
# Update destribution manifest yaml versions.
for f in $ALL_MANIFEST_YAML; do
    perl -pi -e "s/version: \d+\.\d+\.\d+/version: $VERSION/g" "$f"
    echo "Updated version in distribution yaml \`$f\` with version v$VERSION"
done


if [ -z "$CONTRIB_VERSION" ]; then
  echo "CONTRIB_VERSION not set, skipping."
else
  for f in $ALL_MANIFEST_YAML; do
      # update solarwinds contrib references
      perl -pi -e "s|^(\s+- gomod: github.com/solarwinds/solarwinds-otel-collector-contrib/[^ ]*) v[0-9]+\.[0-9]+\.[0-9]+$|\1 v$CONTRIB_VERSION|" "$f"
      echo "References to 'github.com/solarwinds/solarwinds-otel-collector-contrib' in $f updated with version v$CONTRIB_VERSION"
  done

  # We need to run go mod tidy after raising versions of solarwinds-otel-collector-contrib components
  echo "Running go mod tidy"
  find . -name "go.mod" -execdir sh -c 'go mod tidy' \;
fi

if [ -z "$BUILDER_VERSION" ]; then
  echo "BUILDER_VERSION not set, skipping."
else
  YAML_FILE="./.github/workflows/buildAndTest.yml"
  if [ ! -f "$YAML_FILE" ]; then
      echo "buildAndTest.yml not found!"
      exit 1
  fi
  # update BUILDER_VERSION env var in ./.github/workflows/build.yml
  perl -pi -e "s|BUILDER_VERSION: \"v[0-9]+\.[0-9]+\.[0-9]+\"|BUILDER_VERSION: \"v$BUILDER_VERSION\"|g" "$YAML_FILE"
  echo "Updated BUILDER_VERSION in /.github/workflows/buildAndTest.yaml to version $BUILDER_VERSION"
  # update builder version in README.md
  README_FILE="./README.md"
  sed -i '' "/cmd\/builder@v/s|.*|\`go install go.opentelemetry.io/collector/cmd/builder@v$BUILDER_VERSION\`|" $README_FILE
  echo "Updated BUILDER_VERSION in README to version $BUILDER_VERSION"
fi

# update pkg\version\version.go to set the actual release version
GO_VERSION_FILE="./pkg/version/version.go"
if [ ! -f "$GO_VERSION_FILE" ]; then
    echo "version.go not found!"
    exit 1
fi
perl -pi -e "s|^(const Version =) \"[0-9]+\.[0-9]+\.[0-9]+\"$|\1 \"$VERSION\"|" "$GO_VERSION_FILE"
echo "Version.go updated with version $VERSION"
