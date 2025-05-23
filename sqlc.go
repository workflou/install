package main

func sqlc() {
	ensureDir("database")
	ensureDir("database/schema")
	ensureDir("database/query")

	copyStub("sqlc/database/init.go", "database/init.go")
	copyStub("sqlc/database/schema/00001_init.sql", "database/schema/00001_init.sql")
	copyStub("sqlc/database/query/admin.sql", "database/query/admin.sql")
	copyStub("sqlc/sqlc.yaml", "sqlc.yaml")

	runCmd("go", "get", "-tool", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
	runCmd("go", "get", "github.com/pressly/goose/v3")
	runCmd("go", "get", "modernc.org/sqlite")
	runCmd("go", "mod", "tidy")
}
