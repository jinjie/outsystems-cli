## outsystems-cli deployments list

List deployments (newest first)

### Synopsis

Calls GET /deployments/ and returns deployments ordered by creation date.

```
outsystems-cli deployments list [flags]
```

### Options

```
  -h, --help                            help for list
      --max-date string                 maximum creation date (YYYY-MM-DD)
      --min-date string                 minimum creation date (YYYY-MM-DD)
      --target-environment-key string   filter by target environment key
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

