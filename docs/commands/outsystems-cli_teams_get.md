## outsystems-cli teams get

Get the details of a team

### Synopsis

Calls GET /teams/{TeamKey}, with its users and applications.

```
outsystems-cli teams get <team-key> [flags]
```

### Options

```
  -h, --help                   help for get
      --include-applications   include applications in the team
      --include-users          include users in the team
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli teams](outsystems-cli_teams.md)	 - Manage teams

