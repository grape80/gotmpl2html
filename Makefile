MAKEFLAGS += --no-builtin-rules

logDir := log

now = $(shell date '+%Y%m%d-%H%M%S')

.PHONY: gobuild
gobuild:
	echo 'go build'

.PHONY: gotest
gotest:
	ls -la
	cat /etc/passwd | grep 1001 && cat /etc/group | grep 123
	id `whoami` && mkdir $(logDir)
	go test -v -cover -count=1 -coverprofile=$(logDir)/gocover-$(now).out | tee $(logDir)/gotest-$(now).log
	go tool cover -html=$(logDir)/gocover-$(now).out -o $(logDir)/gocover-$(now).html

.PHONY: ci
ci: gotest

.PHONY: cleanlog ## Remove log directory.
cleanlog:
	rm -rvf $(logDir)
