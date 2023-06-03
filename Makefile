MAKEFLAGS += --no-builtin-rules
.DEFAULT_GOAL := help

include .build/.env
app := $(APP_NAME)
mainDir := $(MAIN_DIR)
binDir := $(BIN_DIR)
logDir := $(LOG_DIR)
logTestDir := $(LOG_TEST_DIR)
distDir := $(DIST_DIR)

gobuild_opts := $(shell cat .build/gobuild.opts | tr '\n' ' ')
gofiles := $(shell find . -type f -name '*.go' -print)
now = $(shell date '+%Y%m%d-%H%M%S')

windres := /usr/bin/x86_64-w64-mingw32-windres
winVersionInfo := $(WIN_VERSION_INFO)

.PHONY: help ## Show help.
help:
	@cat $(MAKEFILE_LIST) | grep '##' | grep -v 'MAKEFILE_LIST' | sed s/^.PHONY:// | awk -F \#\# '{ printf "%-20s%s\n", $$1, $$2 }'

##
.PHONY: gobuild ## Build go binary.
gobuild: $(binDir)/$(app)

$(binDir)/$(app): $(gofiles)
	CGO_ENABLED=0 go build $(gobuild_opts) -o $@ ./$(mainDir)
	!(ldd $(binDir)/$(app))

.PHONY: gotest ## Run go test.
gotest:
	mkdir -p $(logTestDir)
	go test -v ./... -coverpkg=./... -count=1 -coverprofile=$(logTestDir)/gocover-$(now).out | tee $(logTestDir)/gotest-$(now).log
	go tool cover -html=$(logTestDir)/gocover-$(now).out -o $(logTestDir)/gocover-$(now).html

.PHONY: gox ## Build go binary for multi platform.
gox: resc
	test $(VERSION) != ''
	mkdir -p $(distDir)
	sh gox.sh

.PHONY: resc
resc: cleansyso
	$(windres) .build/windows/$(winVersionInfo).rc _$(winVersionInfo).syso

##
.PHONY: build ## Build binary.
build: gobuild

.PHONY: test ## Run test.
test: gotest

.PHONY: release ## Build release.
release: gox

.PHONY: cicd ## Run CI/CD.
cicd: build test release

##
.PHONY: cleanbin ## Clean binary.
cleanbin:
	rm -rfv $(binDir)

.PHONY: cleanlog ## Clean log.
cleanlog:
	rm -rfv $(logDir)

.PHONY: cleandist ## Clean dist.
cleandist:
	rm -rfv $(distDir)

.PHONY: cleansyso ## Clean syso.
cleansyso:
	rm -fv *.syso

.PHONY: cleanall ## Clean all.
cleanall: cleanbin cleanlog cleandist cleansyso
