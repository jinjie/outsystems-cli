# Development

Requires a Go toolchain (see the `go` directive in
[`go.mod`](https://github.com/jinjie/outsystems-cli/blob/main/go.mod) for the
version). Common tasks:

```sh
go build ./...          # compile everything
go vet ./...            # static analysis
go test ./...           # run the unit tests
gofmt -w cmd internal   # format
```

Project layout, the HTTP/config/output packages, coding conventions, and a
step-by-step recipe for adding a new command live in
[`AGENTS.md`](https://github.com/jinjie/outsystems-cli/blob/main/AGENTS.md).
Read it before making changes so the codebase stays consistent.

## Command reference generation

The [command reference](commands/index.md) pages are generated from the CLI
command tree with cobra's `GenMarkdownTree`, so they always match the real
`--help` output. Regenerate them after changing commands:

```sh
go run ./tools/gendocs          # writes to docs/commands
```

Do not edit the generated pages by hand; change the command definitions in
`cmd/` instead.

## Keeping up with the LifeTime API

The command tree is written by hand from the LifeTime v2 OpenAPI spec. The spec
itself (`swagger.json`) is **git-ignored** and not committed. To add or update
endpoints:

1. Drop the latest `swagger.json` into the **root folder** of the project.
2. Ask your coding agent to reconcile the CLI against it. The agent picks up
   `swagger.json` from the root, diffs it against the implemented commands, and
   adds or updates any changed endpoints following the conventions in
   `AGENTS.md` (one file per resource, `--data` for bodies, gated query params,
   `url.PathEscape` for path params, etc.).
3. Run `go build ./...`, `go vet ./...`, and `go test ./...` to verify, then
   regenerate the command reference with `go run ./tools/gendocs`.

Because `swagger.json` is ignored, it is safe to leave a working copy in the
root while developing; it will not be committed.
