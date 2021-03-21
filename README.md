# grpctest ![test](https://github.com/kapetndev/grpctest/workflows/test/badge.svg?event=push)

`grpctest` is a module providing utilities for testing gRPC servers.
Specifically it formalises a pattern of writing integration style tests by
exercising the full gRPC stack.

## Prerequisites

You will need the following things properly installed on your computer.

- [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://github.com/golang/go/wiki/Modules) support (Go 1.11+),
simply add the following import

```go
import "github.com/kapetndev/grpctest"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `grpctest` module, run the following command:

```bash
$ go get -u github.com/kapetndev/grpctest
```

## License

This project is licensed under the [MIT License](LICENSE.md).
