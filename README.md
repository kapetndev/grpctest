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

## Usage

To use this module start by implementing a gRPC server as defined within
a protocol buffer definition. This will form the system under test. In this
example a type implementing the `EchoServer` interface has been created to echo
back the request to the caller.

```go
package server

import (
	"context"

	echopb "github.com/kapetndev/kapetn-api-go/echo/v1"
)

type EchoServer struct {
	echopb.UnimplementedEchoServiceServer
}

func (s *EchoServer) Echo(ctx context.Context, in *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	return &echopb.EchoResponse{Message: in.Message}, nil
}
```

To then test the server implementation create a gRPC server using the
`grpctest.NewServer` function. This instantiates a wrapper around a regular
gRPC server using an in-memory buffer to send and receive data. The wrapper
also exposes methods to create gRPC client connections across the buffer. These
may be used when instantiating new gRPC clients.

The example below demonstrates this within the context of a test.

```go
package server_test

import (
	"context"
	"testing"

	"github.com/kapetndev/echo/internal/server"
	"github.com/kapetndev/grpctest"
  echopb "github.com/kapetndev/kapetn-api-go/echo/v1"
)

func setupServer(t *testing.T) (grpctest.Closer, echopb.EchoClient) {
	s := grpctest.NewServer()

	conn, err := s.ClientConn()
	if err != nil {
		t.Fatal(err)
	}

	echopb.RegisterEchoServer(s, &server.EchoServer{})
	s.Serve()

	return s.Close, echopb.NewEchoClient(conn)
}

func TestEcho(t *testing.T) {
	t.Run("returns the same message sent to the server", func(t *testing.T) {
		closer, client := setupServer(t)
		defer closer()

		message := "Hello, world"
		resp, err := client.Echo(context.Background(), &echopb.EchoRequest{Message: message})
		if err != nil {
			t.Errorf("expected: nil, got: %v", err)
		}

		if resp.Message != message {
			t.Errorf("expected: %s, got: %s", message, resp.Message)
		}
	})
}
```

## License

This project is licensed under the [MIT License](LICENSE.md).
