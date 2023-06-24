package server

import (
	gen "MicroservisPlus/proto/generate"
	"golang.org/x/net/context"
)

type GRPCsever struct{}

func (s *GRPCsever) Add(ctx context.Context, req *gen.AddRequest) (*gen.AddResponse, error) {
	return &gen.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
