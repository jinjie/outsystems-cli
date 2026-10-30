## outsystems-cli environments blocked-ips

List IP addresses with blocked login attempts

### Synopsis

Calls GET /environments/blockedipaddresses/.

Use --environment-key to scope the query to a single environment; if omitted,
all environments are returned.

```
outsystems-cli environments blocked-ips [flags]
```

### Options

```
      --environment-key string   scope to a single environment (all if empty)
  -h, --help                     help for blocked-ips
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

