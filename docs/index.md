# OutSystems LifeTime v2 CLI

A command-line interface for the **OutSystems LifeTime REST API v2**, written
in Go with [cobra](https://github.com/spf13/cobra). Output is JSON by default
and written to stdout so it pipes cleanly into tools like `jq`.

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
# 1. Install (macOS/Linux via Homebrew; see Installation for other options).
brew install jinjie/tap/outsystems-cli

# 2. Configure the default profile (prompts for the API key without echo).
outsystems-cli init --hostname lifetime.example.com

# 3. Run a command — output is JSON on stdout, ready to pipe into jq.
outsystems-cli environments list
outsystems-cli applications list --include-env-status | jq '.[].Name'
```

The API key is a LifeTime service account token (generated in the LifeTime
UI). Every request is sent as `Authorization: Bearer <token>` over HTTPS.

Explore the full command tree with `--help` at any level:

```sh
outsystems-cli --help
outsystems-cli applications --help
```

## Where to go next

- [Installation](installation.md) — Homebrew, manual download, `go install`, from source.
- [Configuration](configuration.md) — profiles, selection precedence, global flags.
- [Usage](usage.md) — request bodies, piping, and worked examples.
- [Downloading binaries](downloads.md) — fetching application/source-code packages.
- [Command reference](commands/index.md) — generated from the CLI itself.
- [Development](development.md) — building, testing, and extending the CLI.
- [Releases](releases.md) — how releases and the Homebrew tap are produced.
