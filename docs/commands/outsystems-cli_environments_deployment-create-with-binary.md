## outsystems-cli environments deployment-create-with-binary

Create a deployment from a staging binary

### Synopsis

Calls POST /environments/{EnvironmentKey}/deployment and returns the key of
the new deployment.

The body is the staging binary payload. Provide it with --data (inline string,
@file, or - for stdin). Use @file to point at the binary content.

```
outsystems-cli environments deployment-create-with-binary <environment-key> [flags]
```

### Options

```
      --data string   request body: inline string, @file, or - for stdin
  -h, --help          help for deployment-create-with-binary
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

