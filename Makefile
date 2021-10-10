VERSION=$(shell grep Version version.go | cut -f2 -d\")
TAG_NAME=v$(VERSION)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: build
build:
	go build ./...

.PHONY: release
release:
	git tag $(TAG_NAME) && \
	git push origin $(TAG_NAME)
