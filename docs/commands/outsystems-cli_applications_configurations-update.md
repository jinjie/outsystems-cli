## outsystems-cli applications configurations-update

Update the application configurations per environment

### Synopsis

Calls PUT /applications/{ApplicationKey}/configurations/.

The request body is an ApplicationConfigurationListRecord. Provide it with
--data, which accepts an inline JSON string, @file to read a file, or - to
read from stdin.

```
outsystems-cli applications configurations-update <application-key> [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for configurations-update
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications

