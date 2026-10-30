## outsystems-cli

CLI for the OutSystems LifeTime API v2

### Synopsis

outsystems-cli interacts with the OutSystems LifeTime REST API v2.

Configure a profile with 'outsystems-cli init', then run resource commands
such as 'applications list' or 'environments list'. Output is JSON by default
and written to stdout so it can be piped into tools like jq. Use --output to
write the result to a file instead.

Note: the 'version' subcommand reports the remote LifeTime server version;
use '--version' to see this CLI's own build version.

### Options

```
      --format string    output format (only 'json' is supported) (default "json")
  -h, --help             help for outsystems-cli
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli applications](outsystems-cli_applications.md)	 - Manage and inspect applications
* [outsystems-cli auth](outsystems-cli_auth.md)	 - Authentication operations
* [outsystems-cli deployments](outsystems-cli_deployments.md)	 - Manage deployments
* [outsystems-cli downloads](outsystems-cli_downloads.md)	 - Download binary files produced by LifeTime
* [outsystems-cli environments](outsystems-cli_environments.md)	 - Manage and inspect environments
* [outsystems-cli init](outsystems-cli_init.md)	 - Configure a profile with a LifeTime hostname and API key
* [outsystems-cli modules](outsystems-cli_modules.md)	 - Inspect modules
* [outsystems-cli roles](outsystems-cli_roles.md)	 - Manage platform roles
* [outsystems-cli teams](outsystems-cli_teams.md)	 - Manage teams
* [outsystems-cli users](outsystems-cli_users.md)	 - Manage users
* [outsystems-cli version](outsystems-cli_version.md)	 - Show the LifeTime version and supported Platform Server version

