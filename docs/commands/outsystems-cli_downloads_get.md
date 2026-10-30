## outsystems-cli downloads get

Download a binary by download key or full URL

### Synopsis

Fetches a binary and writes it to a file or stdout.

Provide the source one of two ways:
  - a download key argument: fetched from the LifeTime
    /downloads/{DownloadKey}/ endpoint on the active profile host, with the
    Bearer token attached; or
  - --url <url>: fetched from an absolute URL, typically the "url" field of a
    DownloadLink returned by a *-download command. The Bearer token is sent
    only when the URL host matches the profile host, so pre-signed links to
    other hosts are not sent your LifeTime credential.

Destination (in order of precedence):
  - --output/-o <file> : write to that path
  - --output/-o -      : write to stdout (pipe-friendly, e.g. "-o - > app.oap")
  - otherwise          : use the server-suggested filename in the current
                         directory (from Content-Disposition)
  - if no name is suggested and no --output is given, write to stdout

Existing files are not overwritten unless --force is given.

Examples:
  # Two-step: get the link, then download it.
  LINK=$(outsystems-cli applications version-download <app> <ver> | jq -r .url)
  outsystems-cli downloads get --url "$LINK" -o app.oap

  # One-liner piping the key through jq (when the link is a plain key).
  outsystems-cli downloads get "$KEY" -o app.oap

```
outsystems-cli downloads get [download-key] [flags]
```

### Options

```
      --force        overwrite the destination file if it already exists
  -h, --help         help for get
      --url string   absolute URL to download (e.g. the DownloadLink 'url' field)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli downloads](outsystems-cli_downloads.md)	 - Download binary files produced by LifeTime

