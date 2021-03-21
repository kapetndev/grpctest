package grpctest_test

import (
	"context"
	"testing"

	"github.com/kapetndev/grpctest"
	echopb "github.com/kapetndev/grpctest/testdata/echo/v1"
)

func TestNewServer(t *testing.T) {
	t.Parallel()

	t.Run("returns a server with a buffered listener", func(t *testing.T) {
		if s := grpctest.NewServer(); s == nil {
			t.Error("server was <nil>")
		}
	})
}

func TestNewTLSServer(t *testing.T) {
	t.Parallel()

	t.Run("returns a TLS server with a buffered listener", func(t *testing.T) {
		if s := grpctest.NewTLSServer(); s == nil {
			t.Error("server was <nil>")
		}
	})
}

func TestCertificate(t *testing.T) {
	t.Parallel()

	t.Run("returns nil when the server does not use TLS", func(t *testing.T) {
		if s := grpctest.NewServer(); s.Certificate() != nil {
			t.Errorf("server certificate was not <nil>: %v", s.Certificate())
		}
	})

	t.Run("returns a TLS certificate when the server uses TLS", func(t *testing.T) {
		if s := grpctest.NewTLSServer(); s.Certificate() == nil {
			t.Error("certificate was <nil>")
		}
	})
}

func TestClientConn(t *testing.T) {
	t.Parallel()

	t.Run("returns a client connection to the server", func(t *testing.T) {
		s := grpctest.NewServer()
		t.Cleanup(s.Close)

		conn, err := s.ClientConn()
		if err != nil {
			t.Errorf("error was not <nil>: %s", err)
		}

		if conn == nil {
			t.Error("client connection was <nil>")
		}
	})
}

type server struct {
	echopb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	return &echopb.EchoResponse{Message: in.Message}, nil
}

func setupRecoveryServer(t *testing.T) (grpctest.Closer, echopb.EchoServiceClient) {
	s := grpctest.NewServer()

	conn, err := s.ClientConn()
	if err != nil {
		t.Fatal(err)
	}

	echopb.RegisterEchoServiceServer(s, &server{})
	s.Serve()

	return s.Close, echopb.NewEchoServiceClient(conn)
}

func TestEcho(t *testing.T) {
	t.Run("returns the same message sent to the server", func(t *testing.T) {
		closer, client := setupRecoveryServer(t)
		t.Cleanup(closer)

		expectedMessage := "Hello, world"
		resp, err := client.Echo(context.Background(), &echopb.EchoRequest{Message: expectedMessage})
		if err != nil {
			t.Errorf("error was not <nil>: %s", err)
		}

		if resp.Message != expectedMessage {
			t.Errorf("messages are not equal: %s != %s", resp.Message, expectedMessage)
		}
	})
}
