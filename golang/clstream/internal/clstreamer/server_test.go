package clstreamer

import (
	pb "clstream/internal/api"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var (
	lis *bufconn.Listener
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCLStreamServiceServer(s, &GRPCServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGRPCServer(t *testing.T) {
	tests := []struct {
		name    string
		res     *pb.StatusResponse
		payload string
		errMsg  string
		errCode codes.Code
	}{
		{
			"valid request",
			&pb.StatusResponse{Result: "1 entities written"},
			`{
					"Company": "Apple",
					"Email": "TimCook@apple.com",
					"About": "Timothy Donald Cook"
					 }`,
			"",
			codes.OK,
		},
		{
			"empty request",
			&pb.StatusResponse{Result: ""},
			`{}`,
			"object missing primary index",
			codes.Aborted,
		},
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewCLStreamServiceClient(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := new(pb.Employee)
			err := json.Unmarshal([]byte(tt.payload), e)
			if err != nil {
				log.Fatalln(err)
			}
			req := &pb.FileStream{Req: e}
			resp, err := client.Stream(ctx)
			if err != nil {
				log.Fatalln(err)
			}
			err = resp.Send(req)
			if err != nil {
				log.Fatalln(err)
			}
			reply, err := resp.CloseAndRecv()
			if err != nil {
				if er, ok := status.FromError(err); ok {
					if er.Code() != tt.errCode {
						t.Error("error code: expected", tt.errCode, "received", er.Code())
					}
					if er.Message() != tt.errMsg {
						t.Error("error message: expected", tt.errMsg, "received", er.Message())
					}
				}
			}
			if reply.GetResult() != tt.res.GetResult() {
				t.Error("error message: expected", tt.res.GetResult(), "received", reply.GetResult())
			}
		})
	}
}
