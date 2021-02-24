package main

import (
	pb "clstream/internal/api"
	"clstream/internal/clstreamer"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	GrpcPort string `yaml:"grpcPort"`
}

var c *Config

func main() {
	viper.SetConfigName("server_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Can't read config file")
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	s := grpc.NewServer()

	srv := &clstreamer.GRPCServer{}
	pb.RegisterCLStreamServiceServer(s, srv)

	l, err := net.Listen("tcp", c.GrpcPort)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Printf("Server started at %s port\n", c.GrpcPort)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case <-sigCh:
		log.Println("Received shutdown signal")
		break
	}

	log.Println("Server's going down")
	s.GracefulStop()
}
