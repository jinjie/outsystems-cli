## outsystems-cli applications version-delete

Discard an application version

### Synopsis

Calls DELETE /applications/{ApplicationKey}/versions/{VersionKey}/.

Running versions, or versions used in deployments, cannot be deleted.

```
outsystems-cli applications version-delete <application-key> <version-key> [flags]
```

### Options

```
  -h, --help   help for version-delete
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications

