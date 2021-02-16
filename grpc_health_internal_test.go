package grpc_health

import (
	"testing"
)

func TestAddHealthCheck(t *testing.T) {
  const service = "grpc.test.service"

  healthServer := NewServer()

  h := fakeHealthCheck{result: true}
  healthServer.AddHealthCheck(service, h)

  actual, ok := healthServer.checks[service]
  if !ok {
    t.Errorf("expected '%s', got nil", service)
  }

  if len(actual) == 0 {
    t.Error("expected HealthCheck, got 0")
  }

  if !actual[0].IsHealthy() {
    t.Error("expected IsHealthy = true, got false")
  }
}

type fakeHealthCheck struct {
  result bool
}

func (h fakeHealthCheck) IsHealthy() bool {
  return h.result
}
