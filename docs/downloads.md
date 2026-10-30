# Downloading binaries

The `*-download` endpoints (application/module source code, app versions,
solutions) return a `DownloadLink` JSON object:

```json
{ "url": "...", "expires": "..." }
```

That `url` points at the actual binary (served as `application/octet-stream`
with a `Content-Disposition` filename). Use `downloads get` to stream it:

```sh
# From the DownloadLink url (recommended — it may be pre-signed):
outsystems-cli downloads get --url "<DownloadLink.url>" -o app.oap

# From a bare download key, hitting /downloads/{DownloadKey}/ on the profile host:
outsystems-cli downloads get <download-key> -o app.oap

# Pipe to stdout (e.g. into another tool):
outsystems-cli downloads get --url "<url>" -o - > app.oap
```

## Destination precedence

`-o <file>` → `-o -` (stdout) → server-suggested filename in the current
directory → stdout.

Existing files are not overwritten unless `--force` is given.

## Security note

The Bearer token is sent **only** when the URL host matches the profile host,
so pre-signed links to other hosts (e.g. object storage) never receive your
LifeTime credential. Downloads are streamed, never buffered fully in memory.
