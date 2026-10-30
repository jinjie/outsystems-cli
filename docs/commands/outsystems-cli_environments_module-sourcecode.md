## outsystems-cli environments module-sourcecode

Request the source code package of a module

### Synopsis

Calls POST /environments/{EnvironmentKey}/modules/{ModuleKey}/sourcecodeaccess
and returns the operation key (PackageKey) and status. Poll status with
'module-sourcecode-status' and fetch the link with 'module-sourcecode-download'.

```
outsystems-cli environments module-sourcecode <environment-key> <module-key> [flags]
```

### Options

```
  -h, --help   help for module-sourcecode
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments

