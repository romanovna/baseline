package main

import (
	"google.golang.org/grpc"
	stream "grpc_stream/pkg/api"
	"grpc_stream/pkg/streamer"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()

	srv := &streamer.GRPCServer{}
	stream.RegisterStreamerServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
