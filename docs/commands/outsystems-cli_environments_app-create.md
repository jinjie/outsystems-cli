## outsystems-cli environments app-create

Create a new application in an environment

### Synopsis

Calls POST /environments/{EnvironmentKey}/applications/ and returns the key
of the new application.

Provide an ApplicationCreate body with --data (inline JSON, @file, or - for
stdin). Required fields: TemplateKey, TeamKey, Name, Color.

```
outsystems-cli environments app-create <environment-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for app-create
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

