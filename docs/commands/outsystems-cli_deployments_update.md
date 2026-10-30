## outsystems-cli deployments update

Update a deployment

### Synopsis

Calls PUT /deployments/{DeploymentKey}/ and returns the updated record.

Provide the deployment information to update with --data (inline JSON, @file,
or - for stdin).

```
outsystems-cli deployments update <deployment-key> [flags]
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

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

