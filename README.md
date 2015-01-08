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

### Running web server
In development it is recommended to use [`gin`](https://github.com/codegangsta/gin) for achieving live reloading of code. In order to do this install `gin`:

```bash
$ go get github.com/codegangsta/gin
```

And then from within `$GOPATH/src/github.com/jabclab/lang-exercise-go` you can simply run:

```bash
$ gin --appPort 8888
```

## Testing
Tests are run via `go test`.
