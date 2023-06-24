package server

import (
	sum "MicroservisPlus/proto"
	"context"
)

type GRPCSever struct{}

func (s *GRPCSever) Add(ctx context.Context, req *sum.AddRequest) (*sum.AddResponse, error) {
	return &sum.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
