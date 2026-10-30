## outsystems-cli users create

Create a user

### Synopsis

Calls POST /users/ and returns the created user.

Provide a User body with --data (inline JSON, @file, or - for stdin). Teams
and applications must be associated using the dedicated endpoints.

```
outsystems-cli users create [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for create
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli users](outsystems-cli_users.md)	 - Manage users

