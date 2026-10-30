# Installation

Pre-built binaries for Linux, macOS, and Windows are attached to each
[GitHub release](https://github.com/jinjie/outsystems-cli/releases).

## Homebrew (macOS / Linux)

```sh
brew install jinjie/tap/outsystems-cli
```

This installs from the [`jinjie/homebrew-tap`](https://github.com/jinjie/homebrew-tap)
tap, which GoReleaser updates automatically on each release.

## Linux / macOS (manual download)

Download the archive for your OS/arch, extract the binary, and put it on your
`PATH`. Replace `VERSION` with the release tag (e.g. `v0.1.0`), and pick the
asset matching your platform:

| OS | arch | asset |
| --- | --- | --- |
| Linux | x86_64 (amd64) | `outsystems-cli_Linux_x86_64.tar.gz` |
| Linux | arm64 | `outsystems-cli_Linux_arm64.tar.gz` |
| macOS | Intel (amd64) | `outsystems-cli_Darwin_x86_64.tar.gz` |
| macOS | Apple Silicon (arm64) | `outsystems-cli_Darwin_arm64.tar.gz` |

```sh
VERSION=v0.1.0
ASSET=outsystems-cli_Linux_x86_64.tar.gz   # change to match your platform
BASE=https://github.com/jinjie/outsystems-cli/releases/download/$VERSION

curl -fsSLO "$BASE/$ASSET"
tar -xzf "$ASSET" outsystems-cli
sudo mv outsystems-cli /usr/local/bin/
outsystems-cli --version
```

On macOS, if Gatekeeper blocks the manually downloaded binary, clear the
quarantine flag: `xattr -dr com.apple.quarantine /usr/local/bin/outsystems-cli`.
(The Homebrew install does this for you.)

## Windows

Download `outsystems-cli_Windows_x86_64.zip` from the release, extract
`outsystems-cli.exe`, and place it in a folder on your `PATH` (or add its
folder to `PATH`). For example, in PowerShell:

```powershell
$Version = "v0.1.0"
$Base = "https://github.com/jinjie/outsystems-cli/releases/download/$Version"
Invoke-WebRequest "$Base/outsystems-cli_Windows_x86_64.zip" -OutFile outsystems-cli.zip
Expand-Archive outsystems-cli.zip -DestinationPath .
.\outsystems-cli.exe --version
```

## Verify the download (optional)

Each release includes `checksums.txt`. To verify an archive:

```sh
curl -fsSLO "https://github.com/jinjie/outsystems-cli/releases/download/$VERSION/checksums.txt"
sha256sum --ignore-missing -c checksums.txt   # macOS: shasum -a 256 --ignore-missing -c checksums.txt
```

## Go install

If you have a Go toolchain:

```sh
go install github.com/jinjie/outsystems-cli@latest
```

## From source

```sh
go build -o outsystems-cli .
```

Check the CLI's own build version (distinct from the `version` subcommand,
which reports the remote LifeTime server version):

```sh
./outsystems-cli --version
```
