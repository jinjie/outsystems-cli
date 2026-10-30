## outsystems-cli auth renew

Renew the token of the calling user

### Synopsis

Calls POST /auth/tokens/ and returns the renewed token and its expiry.

The request body is a Duration ({"timePeriod": <int>, "unit": "M"}). Provide it
with --data (inline JSON, @file, or - for stdin). The only documented unit is
"M" (months).

Note: this renews the token server-side and returns a new token; it does not
automatically update the stored profile. Save the returned token with
'outsystems-cli init' if you want to persist it.

```
outsystems-cli auth renew [flags]
```

### Options

```
      --data string   request body: inline JSON, @file, or - for stdin
  -h, --help          help for renew
```

### Options inherited from parent commands

```
      --format string    output format (only 'json' is supported) (default "json")
  -o, --output string    write output to a file instead of stdout
      --profile string   profile to use (overrides $OUTSYSTEMS_CLI_PROFILE and the config default)
```

### SEE ALSO

* [outsystems-cli auth](outsystems-cli_auth.md)	 - Authentication operations

