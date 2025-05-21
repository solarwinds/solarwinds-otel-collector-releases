include Makefile.Common

# Define compatible builder_version with the current version of the collector
builder_version := 0.123.0
ALL_SRC := $(shell find . \( -name "*.go" \) \
							-not -path '*generated*' \
							-type f | sort)

.PHONY: ci-check-licenses
ci-check-licenses: add-licenses check-licenses

.PHONY: add-licenses
add-licenses: $(ADDLICENSE)
	@addLicensesOutput=`$(ADDLICENSE) -y "2025" -c "SolarWinds Worldwide, LLC. All rights reserved." -l "apache" \
	    -v $(ALL_SRC) 2>&1`; \
		if [ "$$addLicensesOutput" ]; then \
			echo "Files modified:"; \
			echo "$$addLicensesOutput"; \
			exit 1; \
		else \
			echo "No files modified."; \
		fi

.PHONY: check-licenses
check-licenses:
	@build/check-licenses.sh ${EXPECTED_GO_LICENSE_HEADER} ${EXPECTED_SHELL_LICENSE_HEADER}

.PHONY: prepare-release
prepare-release:
	@build/prepare-release.sh $(version) $(swi_contrib_version) $(builder_version)

.PHONY: build
build:
	go install go.opentelemetry.io/collector/cmd/builder@v$(builder_version)
	CGO_ENABLED=0 GOEXPERIMENT=boringcrypto builder --config=./distributions/$(distribution)/manifest.yaml
