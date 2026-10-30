## outsystems-cli users set-password

Set a user's password

### Synopsis

Calls POST /users/{UserKey}/setpassword/.

Provide a Password body ({"password": "..."}) with --data (inline JSON, @file,
or - for stdin). Prefer @file or - so the secret is not kept in shell history.

```
outsystems-cli users set-password <user-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for set-password
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli users](outsystems-cli_users.md)	 - Manage users

