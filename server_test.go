package health_test

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func TestRegister(t *testing.T) {
	s := grpc.NewServer()
	healthgrpc.RegisterHealthServer(s, health.NewServer())
	s.Stop()
}
