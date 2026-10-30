## outsystems-cli users get

Get the details of a user

### Synopsis

Calls GET /users/{UserKey} and returns the user, teams, and apps.

```
outsystems-cli users get <user-key> [flags]
```

### Options

```
  -h, --help                        help for get
      --include-application-roles   include application roles
      --include-teams               include teams the user belongs to
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli users](outsystems-cli_users.md)	 - Manage users

