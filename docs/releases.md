# Releases

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

## Homebrew tap

On each stable release, GoReleaser generates a Homebrew cask and commits it to
the [`jinjie/homebrew-tap`](https://github.com/jinjie/homebrew-tap) repository,
so users can `brew install jinjie/tap/outsystems-cli`.

Because this is a cross-repository push, the release workflow needs a token
with write access to the tap repo (the default `GITHUB_TOKEN` cannot push to
another repository). Store it as a repository secret named
`HOMEBREW_TAP_GITHUB_TOKEN`.

The cask strips the macOS quarantine attribute on install (via declarative
`postflight_steps`) so the unsigned binary runs without a Gatekeeper prompt.
