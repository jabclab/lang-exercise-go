# lang-exercise-go

Simple RESTful web service which stores data in Redis

## Development
In order to install dependencies you must first install [`godep`](https://github.com/tools/godep):

```bash
$ go get github.com/tools/godep
```

You then need to restore dependencies locally into your `$GOPATH`:

```bash
$ godep restore
```

## Testing
Tests are run via `go test`.
