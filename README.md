# ctxstdio

`ctxstdio` allows accessing stdio streams in Golang via a common context interface.
By default the `os` streams are returned, but they can be overridden on a context
by calling the `ctxstdio.Set` function.

The intended use is for providing overrideable stdio streams to application wide
contexts - e.g. a context which is cancelled by the signal handler is a good
target.

## Installation

```
go get -u github.com/wrouesnel/ctxstdio
```
