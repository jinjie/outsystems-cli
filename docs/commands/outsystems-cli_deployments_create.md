## outsystems-cli deployments create

Create a deployment

### Synopsis

Calls POST /deployments/ and returns the key of the new deployment.

The request body is a deployment record (SourceEnvironmentKey,
TargetEnvironmentKey, Notes, ApplicationOperations). Provide it with --data
(inline JSON, @file, or - for stdin).

```
outsystems-cli deployments create [flags]
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

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

