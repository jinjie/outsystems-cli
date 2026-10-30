## outsystems-cli modules versions

List versions of a module

### Synopsis

Calls GET /modules/{ModuleKey}/versions/ and returns the version list.

```
outsystems-cli modules versions <module-key> [flags]
```

### Options

```
  -h, --help               help for versions
      --max-versions int   maximum number of versions to return (API default is 5)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli modules](outsystems-cli_modules.md)	 - Inspect modules

