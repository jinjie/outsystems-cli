## outsystems-cli environments app-version-create

Create a new application version from the running application

### Synopsis

Calls POST /environments/{EnvironmentKey}/applications/{ApplicationKey}/versions/
and returns the key of the new version.

Provide an ApplicationVersionCreate body with --data (inline JSON, @file, or -
for stdin).

```
outsystems-cli environments app-version-create <environment-key> <application-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for app-version-create
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

