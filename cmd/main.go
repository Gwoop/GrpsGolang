package main

import (
	sum "MicroservisPlus/proto"
	"MicroservisPlus/proto/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCSever{}
	sum.RegisterPlusServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("ошибка сервера (main.go 18)")
		return
	}

	if err := s.Serve(l); err != nil {
		log.Println("ошибка сервера (main.go 23)")
	}

}
