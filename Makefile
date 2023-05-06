MAKEFLAGS += --no-builtin-rules

.PHONY: test
test:
	docker compose up -d ci
	docker compose ps
