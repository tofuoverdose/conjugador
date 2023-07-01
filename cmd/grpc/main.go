package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vrischmann/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tofuoverdose/conjugador/internal/auth"
	authPort "github.com/tofuoverdose/conjugador/proto/gen/conjugador/auth/v1"
)

const configPrefix = "CONJ_GRPC"

type config struct {
	Debug  bool `envconfig:"default=false"`
	Server struct {
		Addr string
	}
}

func main() {
	config := config{}
	if err := envconfig.InitWithPrefix(&config, configPrefix); err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	lis, err := net.Listen("tcp", config.Server.Addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %s", config.Server.Addr, err)
	}

	s, err := newGrpcServer(config)
	if err != nil {
		log.Fatalf("failed to create grpc server: %s", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %s", err)
	}
}

func newGrpcServer(config config) (*grpc.Server, error) {
	s := grpc.NewServer()

	sAuth, err := auth.NewServer()
	if err != nil {
		return nil, fmt.Errorf("create auth server: %w", err)
	}
	authPort.RegisterServiceServer(s, sAuth)

	if config.Debug {
		reflection.Register(s)
	}

	return s, nil
}
