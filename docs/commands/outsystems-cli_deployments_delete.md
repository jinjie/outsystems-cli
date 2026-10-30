## outsystems-cli deployments delete

Discard a saved deployment

### Synopsis

Calls DELETE /deployments/{DeploymentKey}/. Only deployments in the "saved" state can be deleted.

```
outsystems-cli deployments delete <deployment-key> [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

