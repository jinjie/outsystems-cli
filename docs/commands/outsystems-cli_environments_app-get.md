## outsystems-cli environments app-get

Get the running version of an application in an environment

### Synopsis

Calls GET /environments/{EnvironmentKey}/applications/{ApplicationKey}/.

```
outsystems-cli environments app-get <environment-key> <application-key> [flags]
```

### Options

```
  -h, --help                 help for app-get
      --include-env-status   include per-environment status
      --include-modules      include module details
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

