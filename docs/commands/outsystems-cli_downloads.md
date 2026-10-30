## outsystems-cli downloads

Download binary files produced by LifeTime

### Synopsis

Commands for fetching binary files (application packages, source code
packages, solutions) that the *-download endpoints expose.

The *-download endpoints (e.g. 'applications version-download') return a
DownloadLink JSON object with a "url" and an "expires" timestamp. The "url"
points at a binary file. Use 'downloads get' to fetch that binary, passing
either a download key or the full --url from the DownloadLink.

### Options

```
  -h, --help   help for downloads
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli](outsystems-cli.md)	 - CLI for the OutSystems LifeTime API v2
* [outsystems-cli downloads get](outsystems-cli_downloads_get.md)	 - Download a binary by download key or full URL

