MAKEFLAGS += --no-builtin-rules

.PHONY: gobuild
gobuild:
	echo 'go build'

.PHONY: gotest
gotest:
	echo 'go test'

.PHONY: ci
ci: gotest
