## outsystems-cli environments app-download

Get a download link for a running application's binary

### Synopsis

Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/content/
and returns a link that expires in 60 minutes. Use --type to pick the binary
type when applicable (one of "oap", "apk", or "ipa").

```
outsystems-cli environments app-download <environment-key> <application-key> [flags]
```

### Options

```
  -h, --help          help for app-download
      --type string   binary type: "oap", "apk", or "ipa"
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

