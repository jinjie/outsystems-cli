## outsystems-cli deployments configuration

Set the extension database configuration for a deployment

### Synopsis

Calls POST /deployments/{DeploymentKey}/configuration.

The request body is an array of ExtensionDBFullMapping. Provide it with --data
(inline JSON, @file, or - for stdin).

```
outsystems-cli deployments configuration <deployment-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for configuration
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

