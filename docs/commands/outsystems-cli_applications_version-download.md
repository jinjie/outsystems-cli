## outsystems-cli applications version-download

Get a download link for an application version binary

### Synopsis

Calls GET /applications/{ApplicationKey}/versions/{VersionKey}/content/
and returns a link to download the binary. The link expires in 60 minutes.

```
outsystems-cli applications version-download <application-key> <version-key> [flags]
```

### Options

```
  -h, --help   help for version-download
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications

