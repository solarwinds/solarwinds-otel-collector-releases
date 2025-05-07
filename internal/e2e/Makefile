
tag_arg ?= latest-playground

.PHONY: e2e-tests
e2e-tests:
	IMAGE_TAG=$(tag_arg) go test -tags=e2e
