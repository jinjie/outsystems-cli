## outsystems-cli init

Configure a profile with a LifeTime hostname and API key

### Synopsis

init creates or updates a profile in ~/.outsystems-cli/config.yaml.

Each profile stores the LifeTime hostname (without scheme or path; HTTPS and
the API base path are assumed) and the API access token used for Bearer
authentication.

Provide values with --hostname and --api-key, or omit them to be prompted
interactively. The API key is read without echo when prompted. Use --profile
to target a profile other than "default".

```
outsystems-cli init [flags]
```

### Options

```
      --api-key string    API access token (prompted if omitted)
  -h, --help              help for init
      --hostname string   LifeTime hostname (without scheme or path)
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli](outsystems-cli.md)	 - CLI for the OutSystems LifeTime API v2

