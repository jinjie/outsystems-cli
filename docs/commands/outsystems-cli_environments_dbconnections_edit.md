## outsystems-cli environments dbconnections edit

Update a DB connection

### Synopsis

Calls PUT /environments/{environmentKey}/dbconnections/{dbConnectionName}/.

Provide a DBConnection_Request_POST body with --data (inline JSON, @file, or -
for stdin).

```
outsystems-cli environments dbconnections edit <environment-key> <connection-name> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for edit
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments dbconnections](outsystems-cli_environments_dbconnections.md)	 - Manage environment database connections

