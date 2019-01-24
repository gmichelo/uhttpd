# uhttpd
Very simple and minimal implementation of HTTP(S) file server

## Usage

```
uhttpd [flags] [path]
```

### Flags

`-l` Listen port (default 8080)

`-a` Listen IP address (default localhost)

`-c` Path to certificate file

`-p` Path to private key file

If no `path` is specified, it picks automatically the one in which the command was invoked (working directory).
When both certificate and private key files are specified `uhttpd` will start HTTPS server instead of (unsecured) HTTP.
