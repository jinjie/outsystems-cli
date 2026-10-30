## outsystems-cli environments app-sourcecode-download

Get a download link for an application source code package

### Synopsis

Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/sourcecodeaccess/{PackageKey}/download
and returns a link that expires in 60 minutes.

```
outsystems-cli environments app-sourcecode-download <environment-key> <application-key> <package-key> [flags]
```

### Options

```
  -h, --help   help for app-sourcecode-download
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

