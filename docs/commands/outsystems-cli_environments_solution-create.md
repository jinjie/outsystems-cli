## outsystems-cli environments solution-create

Create a solution in an environment

### Synopsis

Calls POST /environments/{EnvironmentKey}/solution and returns the solution
key.

Provide a SolutionCreate body with --data (inline JSON, @file, or - for stdin).
Required field: ApplicationKeys.

```
outsystems-cli environments solution-create <environment-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for solution-create
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

