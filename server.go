package health

import (
	"context"

	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type HealthCheck interface {
	isHealthy() bool
}

// Server implements grpc health checks
type Server struct {
	healthpb.UnimplementedHealthServer
	checks map[string][]HealthCheck
}

// NewServer returns a new Server.
func NewServer(checks map[string][]HealthCheck) *Server {
	return &Server{
		checks: checks,
	}
}

func (s *Server) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	if checks, ok := s.checks[in.Service]; ok {
		for _, c := range checks {
			if !c.isHealthy() {
				return &healthpb.HealthCheckResponse{
					Status: healthpb.HealthCheckResponse_NOT_SERVING,
				}, nil
			}
		}

		return &healthpb.HealthCheckResponse{
			Status: healthpb.HealthCheckResponse_SERVING,
		}, nil
	}

	return nil, status.Error(codes.NotFound, "unknown service")
}

func (s *Server) Watch(in *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "Unimplemented")
}

func (s *Server) AddHealthCheck(service string, check HealthCheck) {
  if _, ok := s.checks[service]; !ok {
    s.checks[service] = []HealthCheck{}
  }
  s.checks[service] = append(s.checks[service], check)
}
