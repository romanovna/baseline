package clstreamer

import (
	pb "clstream/gen/pb-go/api/proto"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var (
	lis           *bufconn.Listener
	commonPayload = `{
					"Company": "Example",
					"Email": "one@example.com",
					"About": "Hello world"
					 }`
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

	//Upload some data
	Setup()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Setup() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	e := new(pb.Employee)
	err = json.Unmarshal([]byte(commonPayload), e)
	if err != nil {
		log.Fatalln(err)
	}
	req := &pb.FileStream{Req: e}

	client := pb.NewCLStreamServiceClient(conn)
	resp, err := client.Stream(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	err = resp.Send(req)
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = resp.CloseAndRecv()
}

func Test_Stream(t *testing.T) {
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

func Test_GetEmployee_Basic(t *testing.T) {
	tests := []struct {
		name    string
		req     string
		res     *pb.Employee
		errMsg  string
		errCode codes.Code
	}{
		{
			"Request: Non existed ID ",
			"worker@example.com",
			&pb.Employee{},
			"",
			codes.Internal,
		},
		{
			"Valid request",
			"one@example.com",
			&pb.Employee{
				Company: "Example",
				Email:   "one@example.com",
				About:   "Hello world",
			},
			"",
			codes.OK,
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
			req := pb.EmployeeReq{Id: tt.req}
			resp, err := client.GetEmployee(ctx, &req)
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
			if resp.String() != tt.res.String() {
				t.Error("error message: expected", tt.res.String(), "received", resp.String())
			}
		})
	}
}

func Test_GetEmployee_Context(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewCLStreamServiceClient(conn)
	cancellingCtx, cancel := context.WithCancel(ctx)
	time.AfterFunc(5*time.Millisecond, cancel)
	resp, err := client.GetEmployee(cancellingCtx,
		&pb.EmployeeReq{Id: "asd@qwerty.com"})
	assert.Nil(t, resp)
	assert.Error(t, context.Canceled, err)
}
