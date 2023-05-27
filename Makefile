MAKEFLAGS += --no-builtin-rules
.DEFAULT_GOAL := help

app := gotmpl2html
mainDir := cmd/$(app)
binDir := bin
logDir := log
distDir := dist
gobuild_opts := $(shell cat .build/gobuild.opts | tr '\n' ' ')

gofiles := $(shell find . -type f -name '*.go' -print)
now = $(shell date '+%Y%m%d-%H%M%S')

.PHONY: help ## Show help.
help:
	@cat $(MAKEFILE_LIST) | grep '##' | grep -v 'MAKEFILE_LIST' | sed s/^.PHONY:// | awk -F \#\# '{ printf "%-15s%s\n", $$1, $$2 }'

.PHONY: gobuild ## Build go binary.
gobuild: $(binDir)/$(app)

$(binDir)/$(app): $(gofiles)
	go build $(gobuild_opts) -o $@ ./$(mainDir)

.PHONY: gobuildx ## Build go binary for multi platform.
gobuildx:

.PHONY: gotest ## Run go test.
gotest:
	mkdir -p $(logDir)
	go test -v ./... -coverpkg=./... -count=1 -coverprofile=$(logDir)/gocover-$(now).out | tee $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html

.PHONY: ci ## Run CI.
ci: gotest

.PHONY: cd ## Run CD.
cd: gobuildx

.PHONY: cleanbin ## Clean binary.
cleanbin:
	rm -rvf $(binDir)

.PHONY: cleanlog ## Clean log.
cleanlog:
	rm -rvf $(logDir)

.PHONY: cleandist ## Clean dist.
cleandist:
	rm -rvf $(distDir)

.PHONY: cleanall ## Clean all.
cleanall: cleanbin cleanlog cleandist
