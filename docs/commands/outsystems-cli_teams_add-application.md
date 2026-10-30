## outsystems-cli teams add-application

Assign an application to a team

### Synopsis

Calls POST /teams/{TeamKey}/applications/.

Provide an ApplicationShortInfo body (Name, Key) with --data (inline JSON,
@file, or - for stdin).

```
outsystems-cli teams add-application <team-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for add-application
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli teams](outsystems-cli_teams.md)	 - Manage teams

