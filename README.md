# Go's install commands

This is a helper tool to quickly bootstrap a new Go project with a handful of opinionated options.

Currently available options:

[x] chi router

## Installation

```
go get -tool github.com/workflou/install
```

## Usage

```
go tool install <command>

# e.g.
go tool install chi
```

## Roadmap

[ ] air hot reload
[ ] sqlc
[ ] templ
[ ] goose
[ ] htmx
[ ] tailwindcss

## Contributing

Feel free to submit your own options. Add a new file with a tool of your choice, register it in `main.go` and provide stub files in `stub` directory. 