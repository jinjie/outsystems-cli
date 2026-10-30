## outsystems-cli deployments exec

Execute a deployment command (start, continue, abort)

### Synopsis

Calls POST /deployments/{DeploymentKey}/{Command}/. The command must be one of "start", "continue", or "abort".

```
outsystems-cli deployments exec <deployment-key> <command> [flags]
```

### Options

```
      --continue-with-errors    continue the deployment even if errors occur
  -h, --help                    help for exec
      --include-error-details   include more error detail in returned messages
      --redeploy-outdated       redeploy outdated applications in the target environment
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments

