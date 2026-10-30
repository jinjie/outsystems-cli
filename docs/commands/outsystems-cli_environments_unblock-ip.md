## outsystems-cli environments unblock-ip

Unblock an IP address

### Synopsis

Calls DELETE /environments/blockedipaddresses/{ipaddress}/.

Use --environment-key to scope the unblock to a single environment; if omitted,
the IP is unblocked in all environments.

```
outsystems-cli environments unblock-ip <ip-address> [flags]
```

### Options

```
      --environment-key string   scope to a single environment (all if empty)
  -h, --help                     help for unblock-ip
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

