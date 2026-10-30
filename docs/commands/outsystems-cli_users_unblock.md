## outsystems-cli users unblock

Unblock a user's login attempts

### Synopsis

Calls DELETE /users/{UserKey}/blockedlogins/.

```
outsystems-cli users unblock <user-key> [flags]
```

### Options

```
      --environment-key string   environment to unblock (all if empty)
  -h, --help                     help for unblock
      --ip-address string        IP address to unblock (all if empty)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli users](outsystems-cli_users.md)	 - Manage users

