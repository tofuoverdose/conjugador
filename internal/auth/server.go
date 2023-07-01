package auth

import (
	"github.com/tofuoverdose/conjugador/proto/gen/conjugador/auth/v1"
)

type server struct {
	auth.UnimplementedServiceServer
}

func NewServer() (server, error) {
	s := server{}
	return s, nil
}
