# http-echo
http-echo is an http server for examining request headers sent by http clients.

## Usage

Running `http-echo` with no arguments will start an http server listening on port 8000 on all available interfaces.

```text
$ http-echo
2020/12/05 23:05:46 http-echo
2020/12/05 23:05:46 Listening on :8000...
```

A specific network address to listen on can be specified using the `-http` command line argument.
```text
$ http-echo -http 127.0.0.1:8080
2020/12/05 23:05:46 http-echo
2020/12/05 23:05:46 Listening on 127.0.0.1:8080...
```
