# AGENTS.md

Context and coding conventions for the `outsystems-cli` project. Read this
before making changes so the codebase stays consistent.

## What this is

A command-line interface for the **OutSystems LifeTime REST API v2**, written
in Go using [cobra](https://github.com/spf13/cobra). It wraps LifeTime
endpoints as resource/verb subcommands (e.g. `applications list`,
`deployments create`) and prints JSON suitable for piping.

- Module path: `github.com/jinjie/outsystems-cli`
- API reference: `swagger.json` in the repo root (Swagger 2.0, basePath
  `/lifetimeapi/rest/v2`).

## Project layout

```
main.go                     Entrypoint; calls cmd.Execute().
cmd/                        Cobra command tree (one file per resource).
  root.go                   Root command, global flags, shared resolvers.
  helpers.go                runAndPrint, readBodyInput, addBoolQuery.
  init.go                   `init` command (configure a profile).
  version.go                `version` command.
  applications.go           `applications ...`
  environments.go           `environments ...`
  deployments.go            `deployments ...`
  modules.go                `modules ...`
  roles.go                  `roles ...`
  teams.go                  `teams ...`
  users.go                  `users ...`
internal/config/            Profile config: load/save, path, precedence.
internal/client/            HTTP client: URL building, Bearer auth, verbs.
internal/output/            JSON rendering to stdout or file.
tools/gendocs/              Generates docs/commands/ from the cobra tree.
docs/                       MkDocs Material site (GitHub Pages). Hand-written
                            pages + generated commands/ reference.
mkdocs.yml                  Docs site config (nav, theme).
README.md                   Short landing page linking to the docs site.
```

## Configuration

- File: `~/.outsystems-cli/config.yaml` on all platforms (home resolved per-OS
  via `os.UserHomeDir()`). On Windows this is `%USERPROFILE%\.outsystems-cli\`.
- Written with `0600` perms (directory `0700`) because it holds API tokens.
- Shape:
  ```yaml
  default_profile: default
  profiles:
    default:
      hostname: lifetime.example.com
      api_key: <token>
  ```
- Profile precedence (see `Config.ResolveProfileName`):
  `--profile` flag → `OUTSYSTEMS_CLI_PROFILE` env → config `default_profile` →
  built-in `default`.

## Authentication

Every API request sends `Authorization: Bearer <api_key>`. The hostname is
stored without scheme or path; the client always uses HTTPS and prepends
`client.BasePath` (`/lifetimeapi/rest/v2`).

## Output and piping conventions

- **JSON is the default and only format for now.** `--format` exists and is
  validated, but only `json` is accepted. Add new formats in
  `internal/output` without changing call sites.
- **stdout carries only the payload.** All diagnostics, prompts, and progress
  go to **stderr**. This keeps commands pipe-friendly (`... | jq ...`).
- `--output <file>` / `-o` writes the payload to a file and keeps stdout quiet
  on success.
- Successful JSON is re-indented for readability but stays valid. Non-JSON
  responses (some endpoints return a bare key as text) pass through unchanged.
- **Binary downloads are a separate path.** The `*-download` endpoints return a
  `DownloadLink` (JSON `{url, expires}`), which is printed like any other JSON.
  Fetching the actual bytes is done by `downloads get` (`cmd/downloads.go`),
  which bypasses `runAndPrint`: it calls `client.DownloadKey` / `DownloadURL`
  (streaming, never buffering the whole file) and `output.StreamBinary` to
  write to `-o <file>`, the server-suggested `Content-Disposition` filename, or
  stdout (`-o -`). The Bearer token is attached only when the target host
  matches the profile host, so pre-signed cross-host links are not sent the
  credential. Existing files are not overwritten without `--force`.

## Global flags (persistent, defined in root.go)

- `--profile` — profile name override.
- `--format` — output format (`json` only for now).
- `--output, -o` — write to a file instead of stdout.

## Coding conventions

Follow these so other agents do not steer away from the established style:

1. **One file per top-level resource** in `cmd/`. The file defines a parent
   `cobra.Command` (the noun) and its subcommands (the verbs), wired together
   in that file's `init()` which also calls `rootCmd.AddCommand(...)`.
2. **Every command is a thin wrapper.** The `RunE` body should:
   - call `runAndPrint(func() ([]byte, error) { ... })` for anything that
     returns a payload (including writes that return a key/record),
   - inside the closure, call `resolveClient()` first, then the client verb,
   - use `commandContext(cmd)` as the context so cancellation propagates.
3. **Path parameters** are escaped with `url.PathEscape(...)` and concatenated
   onto the API path. Keep the trailing slash exactly as the swagger path has
   it (some endpoints have it, some do not — match swagger precisely).
4. **Query parameters** use `url.Values`. For optional booleans, use
   `addBoolQuery(q, "ApiParamName", flagVal, cmd.Flags().Changed("flag-name"))`
   so the parameter is only sent when the user set it, letting the API apply
   its own defaults. For optional strings/ints, gate with
   `cmd.Flags().Changed(...)` before `q.Set(...)`.
5. **Request bodies** come from a `--data` flag resolved via
   `readBodyInput(...)`, which supports inline JSON, `@file`, and `-` (stdin).
   Do not invent per-field flags for complex bodies; keep bodies as JSON so the
   CLI mirrors the API and stays scriptable.
6. **Flag naming:** kebab-case flags mapping to the swagger PascalCase params
   (e.g. `--include-modules` → `IncludeModules`). Store flag values in a
   small package-level struct named `<cmd>Flags`.
7. **Errors:** return them; `cmd.Execute()` prints `Error: ...` to stderr and
   exits non-zero. `rootCmd` sets `SilenceUsage`/`SilenceErrors` so usage noise
   is not dumped on runtime failures. Non-2xx responses surface as
   `client.APIError` including the server body.
8. **No secrets on stdout.** Never echo the API key. `init` reads it without
   echo when stdin is a TTY.
9. **Keep comments purposeful** — doc comments on exported identifiers and
   brief rationale where behavior is non-obvious (e.g. why a param is gated).
10. **Verify before done:** `go build ./...` and `go vet ./...` must pass.

## Adding a new command (recipe)

1. Find the endpoint in `swagger.json` (operationId, path, params, body).
2. Add a subcommand in the matching resource file (or a new `cmd/<resource>.go`
   with a parent command if the resource is new).
3. Map query params with `addBoolQuery` / gated `q.Set`, body with `--data` +
   `readBodyInput`, path params with `url.PathEscape`.
4. Register it in the file's `init()`.
5. Run `go build ./...`, `go vet ./...`, and `go test ./...`.
6. Regenerate the command reference with `go run ./tools/gendocs` (writes
   `docs/commands/`). Add a content page under `docs/` only if the command
   introduces a new concept worth a narrative, not for every command.

## Syncing with the LifeTime API (swagger.json)

The command tree is written by hand from the LifeTime v2 OpenAPI spec. The spec
file (`swagger.json`) is **git-ignored** and not committed, so a working copy
can live in the repo root during development without being committed.

When asked to update the CLI against a new LifeTime version:

1. Expect `swagger.json` in the **repo root** (the user drops it there).
2. Diff it against the implemented commands: for each path + method in the
   spec, confirm there is a matching subcommand; note added, removed, or
   changed endpoints (new params, new bodies, changed paths).
3. Add or update commands following the conventions above — one file per
   resource, `--data` for bodies, gated query params, `url.PathEscape` for path
   params, trailing slash matching the spec exactly.
4. Verify with `go build ./...`, `go vet ./...`, `go test ./...`.
5. Regenerate docs with `go run ./tools/gendocs`.

## Documentation

- The public docs site is built with MkDocs Material from `docs/` and deployed
  to GitHub Pages by `.github/workflows/docs.yml` on push to `main`.
- `README.md` is a short landing page that links to the site; do not grow it
  back into full docs. Narrative docs belong under `docs/`.
- `docs/commands/` is **generated** by `go run ./tools/gendocs` (cobra
  `GenMarkdownTree`). Never hand-edit those files; change the command
  definitions in `cmd/` and regenerate.
- The root command is exposed to tooling via `cmd.RootCmd()`.

## Status

All LifeTime v2 API endpoints from `swagger.json` are implemented. The only
path intentionally not exposed as its own command is `/downloads/{DownloadKey}/`
(marked `@hide` in swagger as "not really an API method"); it is handled by the
`downloads get` command, which streams the binary behind the `DownloadLink`
that the `*-download` endpoints return.
