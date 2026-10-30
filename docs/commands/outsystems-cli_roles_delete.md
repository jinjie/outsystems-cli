## outsystems-cli roles delete

Delete a role

### Synopsis

Calls DELETE /roles/{RoleKey}.

Use --users-new-role-key to grant an alternative role to users who currently
have the role being deleted.

```
outsystems-cli roles delete <role-key> [flags]
```

### Options

```
  -h, --help                        help for delete
      --users-new-role-key string   role to grant to users who had the deleted role
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli roles](outsystems-cli_roles.md)	 - Manage platform roles

