# OutSystems LifeTime v2 CLI

[![Docs](https://img.shields.io/badge/docs-online-blue)](https://jinjie.github.io/outsystems-cli/)
[![Docs build](https://github.com/jinjie/outsystems-cli/actions/workflows/docs.yml/badge.svg)](https://github.com/jinjie/outsystems-cli/actions/workflows/docs.yml)
[![Release](https://img.shields.io/github/v/release/jinjie/outsystems-cli?sort=semver)](https://github.com/jinjie/outsystems-cli/releases/latest)
[![Release build](https://github.com/jinjie/outsystems-cli/actions/workflows/release.yml/badge.svg)](https://github.com/jinjie/outsystems-cli/actions/workflows/release.yml)

A command-line interface for the **OutSystems LifeTime REST API v2**, written
in Go with [cobra](https://github.com/spf13/cobra). Output is JSON by default
and written to stdout so it pipes cleanly into tools like `jq`.

📖 **Full documentation: <https://jinjie.github.io/outsystems-cli/>**

> Working on this project? Read [AGENTS.md](./AGENTS.md) for the coding
> conventions and project context.

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

## Quick start

```sh
# 1. Install (macOS/Linux via Homebrew; see the docs for other options).
brew install jinjie/tap/outsystems-cli

# 2. Configure the default profile (prompts for the API key without echo).
outsystems-cli init --hostname lifetime.example.com

# 3. Run a command — output is JSON on stdout, ready to pipe into jq.
outsystems-cli environments list
outsystems-cli applications list --include-env-status | jq '.[].Name'
```

Explore the full command tree with `--help` at any level:

```sh
outsystems-cli --help
outsystems-cli applications --help
```

## Documentation

The full docs are published at <https://jinjie.github.io/outsystems-cli/>:

- [Installation](https://jinjie.github.io/outsystems-cli/installation/) — Homebrew, manual download, `go install`, from source.
- [Configuration](https://jinjie.github.io/outsystems-cli/configuration/) — profiles, selection precedence, global flags.
- [Usage](https://jinjie.github.io/outsystems-cli/usage/) — request bodies, piping, worked examples.
- [Downloading binaries](https://jinjie.github.io/outsystems-cli/downloads/) — fetching packages.
- [Command reference](https://jinjie.github.io/outsystems-cli/commands/) — generated from the CLI.
- [Development](https://jinjie.github.io/outsystems-cli/development/) — building, testing, extending.
- [Releases](https://jinjie.github.io/outsystems-cli/releases/) — releases and the Homebrew tap.

The docs source lives under [`docs/`](./docs) and is built with MkDocs Material.

## License

[MIT](./LICENSE)
