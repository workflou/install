.PHONY: watch
watch:
	go tool air -c .air.toml

.PHONY: build
build: sqlc templ
	go build -o ./tmp/main .

.PHONY: templ
templ:
	go tool templ generate

.PHONY: create-migration
create-migration:
	GOOSE_DRIVER=sqlite3 GOOSE_MIGRATION_DIR=./schema go tool goose -s create migration sql

.PHONY: sqlc
sqlc:
	go tool sqlc generate