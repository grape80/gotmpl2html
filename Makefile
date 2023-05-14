MAKEFLAGS += --no-builtin-rules
.DEFAULT_GOAL := help

logDir := log

now = $(shell date '+%Y%m%d-%H%M%S')

.PHONY: help ## Show help.
help:
	@cat $(MAKEFILE_LIST) | grep '##' | grep -v 'MAKEFILE_LIST' | sed s/^.PHONY:// | awk -F \#\# '{ printf "%-15s%s\n", $$1, $$2 }'

.PHONY: gobuild ## Build go binary.
gobuild:
	echo 'go build'

.PHONY: gotest ## Run go test.
gotest:
	mkdir $(logDir)
	go test -v -cover -count=1 -coverprofile=$(logDir)/gocover-$(now).out | tee $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html

.PHONY: ci ## Run CI.
ci: gotest

.PHONY: cleanlog ## Clean log.
cleanlog:
	rm -rvf $(logDir)
