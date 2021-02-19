package streamer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc_stream/internal/api"
	"log"
	"os"
)

var (
	decError = "error during decoding JSON file"
)

type GRPCServer struct{}

func (s *GRPCServer) Stream(req *pb.FileIdRequest, srv pb.StreamService_StreamServer) error {
	id := req.GetId()
	log.Println("Fetching ID...", id)

	fileName := fmt.Sprintf("ex%s.json", id)
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Cannot open %v\n", fileName)
		return status.Errorf(codes.InvalidArgument, "cannot open ex%v.json document", req.GetId())
	}
	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	_, err = dec.Token()
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "%s - %v\n", decError, fileName)
	}
	for dec.More() {
		m := new(pb.Employee)

		err := dec.Decode(&m)
		if err != nil {
			return status.Errorf(codes.Internal, "%s - %v\n", decError, fileName)
		}
		if err := srv.Send(&pb.FileResponse{
			Result: m,
		}); err != nil {
			log.Println(err)
			return status.Errorf(codes.Internal, "error at sending operation %+s", err.Error())
		}
		//Just for testing of graceful shutdown func
		//time.Sleep(2 * time.Second)
	}

	_, err = dec.Token()
	if err != nil {
		log.Println(err)
		return status.Errorf(codes.Internal, "%s - %v\n", decError, fileName)
	}
	defer f.Close()
	return nil
}
