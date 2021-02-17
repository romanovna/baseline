package main

import (
	"context"
	"encoding/json"
	"flag"
	"google.golang.org/grpc"
	stream "grpc_stream/pkg/api"
	"io"
	"log"
)

func main() {
	var id string
	flag.StringVar(&id, "id", "1", "Id of json file to get streamed")
	flag.Parse()

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := stream.NewStreamerClient(conn)
	in := &stream.Request{Id: id}

	getStream, err := c.Stream(context.Background(), in)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := getStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			respJson, err := json.MarshalIndent(resp.Result, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			log.Printf("%s\n", string(respJson))

		}
	}()
	<-done
	log.Println("Done")
}
