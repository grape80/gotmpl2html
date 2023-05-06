MAKEFLAGS += --no-builtin-rules

.PHONY: gobuild
gobuild:
	go version

.PHONY: dc.up%
dc.up%:
	docker compose up -d ${@:dc.up-%:%}
	docker compose ps
