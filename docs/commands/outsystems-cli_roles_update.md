## outsystems-cli roles update

Update a role

### Synopsis

Calls PUT /roles/{RoleKey} and returns the key of the updated role.

Provide the PlatformRole body with --data (inline JSON, @file, or - for stdin).

```
outsystems-cli roles update <role-key> [flags]
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

* [outsystems-cli roles](outsystems-cli_roles.md)	 - Manage platform roles

