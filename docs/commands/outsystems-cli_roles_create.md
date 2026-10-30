## outsystems-cli roles create

Create a role

### Synopsis

Calls POST /roles/ and returns the key of the new role.

The request body is a PlatformRole. Provide it with --data (inline JSON,
@file, or - for stdin).

```
outsystems-cli roles create [flags]
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

* [outsystems-cli roles](outsystems-cli_roles.md)	 - Manage platform roles

