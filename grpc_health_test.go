package grpc_health_test

import (
	"testing"

	"google.golang.org/grpc"
  "github.com/danielpieper/go-grpc-health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func TestRegister(t *testing.T) {
	s := grpc.NewServer()
	healthgrpc.RegisterHealthServer(s, grpc_health.NewServer())
	s.Stop()
}
