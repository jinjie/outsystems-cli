# Usage

Output is JSON by default and written to stdout, so commands pipe cleanly into
tools like `jq`. All diagnostics, prompts, and progress go to stderr, keeping
stdout clean for piping. Use `--output <file>` / `-o` to write the payload to a
file instead (stdout stays quiet on success).

## Request bodies

Write commands that take a JSON body accept `--data`:

```sh
--data '{"Name":"My App"}'      # inline JSON
--data @body.json               # read from a file
--data -                        # read from stdin (pipe-friendly)
```

Complex bodies are passed as JSON rather than per-field flags, so the CLI
mirrors the API and stays scriptable.

## Examples

```sh
outsystems-cli version
outsystems-cli applications list --include-env-status | jq '.[].Name'
outsystems-cli environments list
outsystems-cli deployments create --data @deploy.json
outsystems-cli deployments exec <deployment-key> start
outsystems-cli users get <user-key> --include-teams

# Download a binary: fetch the link, then stream the file.
LINK=$(outsystems-cli applications version-download <app-key> <ver-key> | jq -r .url)
outsystems-cli downloads get --url "$LINK" -o app.oap
```

## Discovering commands

Every LifeTime v2 API endpoint is implemented. Run `--help` at any level to
explore the command tree, or browse the [command reference](commands/index.md):

```sh
outsystems-cli --help
outsystems-cli environments --help
outsystems-cli environments dbconnections --help
```

Resource groups: `applications`, `environments` (including
`environments dbconnections`), `deployments`, `modules`, `roles`, `teams`,
`users`, `downloads`, `auth`, plus `init` and `version`.
