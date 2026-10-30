## outsystems-cli applications versions

List versions of an application

### Synopsis

Calls GET /applications/{ApplicationKey}/versions/ and returns the version list.

```
outsystems-cli applications versions <application-key> [flags]
```

### Options

```
      --change-log-filter string   only versions whose change log contains this string
  -h, --help                       help for versions
      --max-versions int           maximum number of versions to return (API default is 5)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications

