## outsystems-cli environments maintenance-mode

Set the maintenance mode of an environment

### Synopsis

Calls PUT /environments/{EnvironmentKey}/maintenancemode.

The request body is a record with IsMaintenanceModeEnabled and Reason. Provide
it with --data (inline JSON, @file, or - for stdin).

```
outsystems-cli environments maintenance-mode <environment-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for maintenance-mode
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

