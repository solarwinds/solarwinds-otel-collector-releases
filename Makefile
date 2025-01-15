include Makefile.Common

ALL_GO_SRC := $(shell find . \( -name "*.go" \) \
			     -not -path '*generated*' \
			     -type f | sort)

ALL_SHELL_SRC := $(shell find . \( -name "*.sh" \) \
			     -type f | sort)

.PHONY: ci-check-licenses
ci-check-licenses: add-licenses check-licenses

.PHONY: add-licenses
add-licenses: $(ADDLICENSE)
	@addLicensesOutput=`$(ADDLICENSE) -y "2025" -c "SolarWinds Worldwide, LLC. All rights reserved." -l "apache" \
	    -v $(ALL_GO_SRC) $(ALL_SHELL_SRC) 2>&1`; \
		if [ "$$addLicensesOutput" ]; then \
			echo "Files modified:"; \
			echo "$$addLicensesOutput"; \
			exit 1; \
		else \
			echo "No files modified."; \
		fi

.PHONY: check-licenses
check-licenses:
	@checkResult=$$(for f in $(ALL_GO_SRC) ; do \
	         if [[ $$f == "*.go" ]] && ! diff -q <(head -n 13 $$f) $(EXPECTED_GO_LICENSE_HEADER) > /dev/null; then \
				  echo "Diff for $$f:"; \
				  diff --label $$f -u <(head -n 13 $$f) $(EXPECTED_GO_LICENSE_HEADER); \
			 elif [[ $$f == "*.sh" ]] && ! diff -q <(head -n 13 $$f) $(EXPECTED_SHELL_LICENSE_HEADER) > /dev/null; then \
               				  echo "Diff for $$f:"; \
               				  diff --label $$f -u <(head -n 13 $$f) $(EXPECTED_SHELL_LICENSE_HEADER); \
		     fi; \
	   done); \
	   if [ -n "$${checkResult}" ]; then \
	           echo "License header check failed:"; \
	           echo "$${checkResult}"; \
	           exit 1; \
	   fi
