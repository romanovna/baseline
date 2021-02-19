package main

import (
	"context"
	"encoding/json"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	pb "grpc_stream/internal/api"
	"io"
	"log"
)

func main() {
	viper.BindEnv("FILE_ID")
	viper.AutomaticEnv()

	pflag.String("file_id", "", "id of file to get in streaming")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	id := viper.GetString("file_id")

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	c := pb.NewStreamServiceClient(conn)
	in := &pb.FileIdRequest{Id: id}

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
