# Go's install commands

This is a helper tool to quickly bootstrap a new Go project with a handful of opinionated options.

Currently available options:

- [x] [chi router](https://github.com/go-chi/chi)
- [x] [air hot reload](https://github.com/air-verse/air)
- [x] [sqlc](https://github.com/sqlc-dev/sqlc)
- [x] [templ](https://github.com/a-h/templ)
- [x] [goose migrations](https://github.com/pressly/goose)
- [x] [htmx](https://github.com/bigskysoftware/htmx)
- [x] [tailwindcss](https://tailwindcss.com/blog/standalone-cli)

## Installation

```
go get -tool github.com/workflou/install
```

## Usage

```
go tool install <command>

# e.g.
go tool install all

# or separate tools:
go tool install chi
go tool install air
go tool install makefile
go tool install sqlc
go tool install templ
go tool install htmx
go tool install tailwind
```

## Contributing

Feel free to submit your own options. Add a new file with a tool of your choice, register it in `main.go` and provide stub files in `stub` directory. 