## outsystems-cli teams add-user

Add a user to a team with a role

### Synopsis

Calls POST /teams/{TeamKey}/users/.

Provide a UserRoleShortInfo body (userKey, roleKey) with --data (inline JSON,
@file, or - for stdin).

```
outsystems-cli teams add-user <team-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for add-user
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli teams](outsystems-cli_teams.md)	 - Manage teams

