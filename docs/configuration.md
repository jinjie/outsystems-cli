# Configuration

## Configure a profile

```sh
# Interactive (prompts for hostname and API key; key is read without echo):
outsystems-cli init

# Non-interactive:
outsystems-cli init --hostname lifetime.example.com --api-key <token>

# A named profile:
outsystems-cli init --profile staging --hostname lt-staging.example.com --api-key <token>
```

Config is stored at `~/.outsystems-cli/config.yaml` (`%USERPROFILE%\.outsystems-cli\`
on Windows) with `0600` permissions, because it holds API tokens.

The file looks like this:

```yaml
default_profile: default
profiles:
  default:
    hostname: lifetime.example.com
    api_key: <token>
  staging:
    hostname: lt-staging.example.com
    api_key: <token>
```

The hostname is stored without scheme or path; the client always uses HTTPS
and prepends the API base path (`/lifetimeapi/rest/v2`).

## Selecting a profile

Precedence: `--profile` flag → `OUTSYSTEMS_CLI_PROFILE` env var →
`default_profile` in config → `default`.

```sh
outsystems-cli applications list --profile staging
OUTSYSTEMS_CLI_PROFILE=staging outsystems-cli applications list
```

## Global flags

| Flag | Description |
| --- | --- |
| `--profile` | Profile to use. |
| `--format` | Output format. Only `json` is supported for now. |
| `--output`, `-o` | Write the payload to a file instead of stdout. |
