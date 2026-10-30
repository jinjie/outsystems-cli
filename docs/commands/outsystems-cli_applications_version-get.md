## outsystems-cli applications version-get

Get the details of a specific application version

### Synopsis

Calls GET /applications/{ApplicationKey}/versions/{VersionKey}/.

```
outsystems-cli applications version-get <application-key> <version-key> [flags]
```

### Options

```
  -h, --help              help for version-get
      --include-modules   include module details (sent as IncludeModules)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications

