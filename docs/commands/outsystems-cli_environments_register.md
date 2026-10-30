## outsystems-cli environments register

Register an environment in LifeTime

### Synopsis

Calls POST /environments/ to register an environment.

Provide an EnvironmentRegistration body with --data (inline JSON, @file, or -
for stdin). Required fields include EnvironmentName, EnvironmentAddress,
ServiceCenterUsername, ServiceCenterPassword, LifeTimeAddress, and
EnvironmentPosition.

```
outsystems-cli environments register [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for register
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

