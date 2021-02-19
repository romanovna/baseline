package streamer

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	pb "grpc_stream/internal/api"
	"io"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &GRPCServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGRPCServer_Stream(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		res     *pb.FileResponse
		errCode codes.Code
		errMsg  string
	}{
		{
			"valid request with existing id",
			"99",
			&pb.FileResponse{Result: &pb.Employee{
				Company: "EPAM",
				Email:   "Dobkin@epam.com",
				About:   "I am the Dobkin",
			}},
			codes.OK,
			"",
		},
		{
			"invalid request with non existing doc ID",
			"50",
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("cannot open ex%s.json document", "50"),
		},
		{
			"invalid req w/out ID",
			"",
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("cannot open ex%s.json document", ""),
		},
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.FileIdRequest{Id: tt.id}
			response, err := client.Stream(ctx, request)
			if response != nil {
				for {
					rcv, err := response.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						if er, ok := status.FromError(err); ok {
							if er.Code() != tt.errCode {
								t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
							}
						}
						log.Println("Failed to receive:", err)
						break
					}
					if rcv.GetResult().String() != tt.res.GetResult().String() {
						t.Error("error message: expected", tt.res.GetResult(), "received", rcv.GetResult())
					}
				}
			}
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
		})
	}
}
