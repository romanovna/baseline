package main

import (
	"bufio"
	pb "clstream/gen/pb-go/api/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"os"
)

var (
	filename string
	address  string
	port     string
)

func streamFile(ctx context.Context, filename string, client pb.CLStreamServiceClient) (*pb.StatusResponse, error) {
	stream, err := client.Stream(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open %v\n", filename)
	}
	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	_, err = dec.Token()
	if err != nil {
		log.Fatalln(err)
	}

	for dec.More() {
		m := new(pb.Employee)

		err := dec.Decode(&m)
		if err != nil {
			log.Fatalln(err)
		}
		if err := stream.Send(&pb.FileStream{Req: m}); err != nil {
			log.Fatalln(err)
		}
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	return reply, nil
}

func main() {
	viper.SetEnvPrefix("APP")
	viper.BindEnv("FILE")
	viper.AutomaticEnv()

	pflag.StringVarP(&filename, "file", "f", "ex1.json", "file to stream, default is ex1.json")
	pflag.StringVarP(&address, "address", "a", "localhost", "address to connect, default is localhost")
	pflag.StringVarP(&port, "port", "p", "8080", "port to connect, default is 8080")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	addrStr := fmt.Sprintf("%s:%s", address, port)
	conn, err := grpc.Dial(addrStr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	client := pb.NewCLStreamServiceClient(conn)
	resp, err := streamFile(ctx, filename, client)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.GetResult())
}
