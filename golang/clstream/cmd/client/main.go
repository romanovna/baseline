package main

import (
	"bufio"
	pb "clstream/internal/api"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"os"
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
	filename := "./ex1.json"
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
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
