## outsystems-cli teams update

Update a team

### Synopsis

Calls PUT /teams/{TeamKey}. Provide a Team body with --data (inline JSON, @file, or - for stdin).

```
outsystems-cli teams update <team-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for update
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli teams](outsystems-cli_teams.md)	 - Manage teams

