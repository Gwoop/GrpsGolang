package main

import (
	generate "MicroservisPlus/proto/generate"
	"MicroservisPlus/proto/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCsever{}
	generate.RegisterPlusServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	s.Serve(l)
}
