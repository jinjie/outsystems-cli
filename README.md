# OutSystems LifeTime v2 CLI

A command-line interface for the **OutSystems LifeTime REST API v2**, written
in Go with [cobra](https://github.com/spf13/cobra). Output is JSON by default
and written to stdout so it pipes cleanly into tools like `jq`.

> Working on this project? Read [AGENTS.md](./AGENTS.md) for the coding
> conventions.

## Why this exists

I've never been a UI person — I live in the terminal. For me it's faster, and
most of my workflows can be driven by scripts and automated. With coding agents
in the mix, getting things done from the terminal has become a genuine game
changer. OutSystems doesn't ship a CLI, and that gap is what pushed me to build
one.

This is a deliberately simple CLI. The hope is that it saves fellow
administrators from clicking through the LifeTime UI for routine work, and lets
them script the parts of their day that are worth scripting.

It tracks the **latest LifeTime v2 API**. Upgrades may occasionally break
compatibility with older LifeTime installs — if that happens, pin and install a
specific version that matches your LifeTime version.

## Install / build

### Homebrew (macOS / Linux)

```sh
brew install jinjie/tap/outsystems-cli
```

This installs from the [`jinjie/homebrew-tap`](https://github.com/jinjie/homebrew-tap)
tap, which GoReleaser updates automatically on each release.

### From source

```sh
go build -o outsystems-cli .
```

Check the CLI's own build version (distinct from the `version` subcommand,
which reports the remote LifeTime server version):

```sh
./outsystems-cli --version
```

## Releases

Releases are built with [GoReleaser](https://goreleaser.com) and published by
the `release` GitHub Actions workflow when a `v*` tag is pushed:

```sh
git tag v0.1.0
git push origin v0.1.0
```

This produces cross-platform archives (Linux, macOS, Windows for amd64/arm64),
a `checksums.txt`, and a GitHub release with an auto-generated changelog. The
build injects the version, commit, and date into `--version` via `-ldflags`.

To test the pipeline locally without tagging or publishing:

```sh
goreleaser release --snapshot --clean
```

## Configure a profile

```sh
# Interactive (prompts for hostname and API key; key is read without echo):
./outsystems-cli init

# Non-interactive:
./outsystems-cli init --hostname lifetime.example.com --api-key <token>

# A named profile:
./outsystems-cli init --profile staging --hostname lt-staging.example.com --api-key <token>
```

Config is stored at `~/.outsystems-cli/config.yaml` (`%USERPROFILE%\.outsystems-cli\`
on Windows) with `0600` permissions.

## Selecting a profile

Precedence: `--profile` flag → `OUTSYSTEMS_CLI_PROFILE` env var →
`default_profile` in config → `default`.

```sh
./outsystems-cli applications list --profile staging
OUTSYSTEMS_CLI_PROFILE=staging ./outsystems-cli applications list
```

## Global flags

| Flag | Description |
| --- | --- |
| `--profile` | Profile to use. |
| `--format` | Output format. Only `json` is supported for now. |
| `--output`, `-o` | Write the payload to a file instead of stdout. |

## Request bodies

Write commands that take a JSON body accept `--data`:

```sh
--data '{"Name":"My App"}'      # inline JSON
--data @body.json               # read from a file
--data -                        # read from stdin (pipe-friendly)
```

## Examples

```sh
./outsystems-cli version
./outsystems-cli applications list --include-env-status | jq '.[].Name'
./outsystems-cli environments list
./outsystems-cli deployments create --data @deploy.json
./outsystems-cli deployments exec <deployment-key> start
./outsystems-cli users get <user-key> --include-teams

# Download a binary: fetch the link, then stream the file.
LINK=$(./outsystems-cli applications version-download <app-key> <ver-key> | jq -r .url)
./outsystems-cli downloads get --url "$LINK" -o app.oap
```

## Downloading binaries

The `*-download` endpoints (application/module source code, app versions,
solutions) return a `DownloadLink` JSON object: `{"url": "...", "expires": "..."}`.
That `url` points at the actual binary (served as `application/octet-stream`
with a `Content-Disposition` filename). Use `downloads get` to stream it:

```sh
# From the DownloadLink url (recommended — it may be pre-signed):
./outsystems-cli downloads get --url "<DownloadLink.url>" -o app.oap

# From a bare download key, hitting /downloads/{DownloadKey}/ on the profile host:
./outsystems-cli downloads get <download-key> -o app.oap

# Pipe to stdout (e.g. into another tool):
./outsystems-cli downloads get --url "<url>" -o - > app.oap
```

Destination precedence: `-o <file>` → `-o -` (stdout) → server-suggested
filename in the current directory → stdout. Existing files are not overwritten
unless `--force` is given. The Bearer token is sent only when the URL host
matches the profile host, so pre-signed links to other hosts never receive your
LifeTime credential.

---

## Commands

Every LifeTime v2 API endpoint is implemented. Run `--help` at any level to
explore the command tree:

```sh
./outsystems-cli --help
./outsystems-cli environments --help
./outsystems-cli environments dbconnections --help
```

Resource groups: `applications`, `environments` (including
`environments dbconnections`), `deployments`, `modules`, `roles`, `teams`,
`users`, `downloads`, `auth`, plus `init` and `version`. The API reference is
`swagger.json` in the repo root.
