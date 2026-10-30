## outsystems-cli environments csp-set

Set the Content Security Policies of an environment

### Synopsis

Calls PUT /environments/{EnvironmentKey}/content_security_policies.

Provide a CSPDirectives body with --data (inline JSON, @file, or - for stdin).

```
outsystems-cli environments csp-set <environment-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for csp-set
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

