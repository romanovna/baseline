package clstreamer

import (
	pb "clstream/gen/pb-go/api/proto"
	"context"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

type GRPCServer struct{}

var (
	tableName = "employee"
	db        *memdb.MemDB
)

//Init db & schema
func init() {
	var err error
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: {
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
		},
	}

	// Create a new database
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
}

func (s *GRPCServer) Stream(stream pb.CLStreamService_StreamServer) error {
	var count uint

	for {
		f, err := stream.Recv()
		if err == io.EOF {
			//JUST FOR DEBUG
			//Read all records from db
			//txn := db.Txn(false)
			//it, err := txn.Get(tableName, "id")
			//if err != nil {
			//	txn.Abort()
			//	log.Println(err)
			//}
			//for obj := it.Next(); obj != nil; obj = it.Next() {
			//	e := obj.(*pb.Employee)
			//	log.Println(e.GetEmail())
			//}
			return stream.SendAndClose(&pb.StatusResponse{Result: fmt.Sprintf("%v entities written", count)})
		}
		// Create a write transaction
		txn := db.Txn(true)
		if err := txn.Insert(tableName, f.GetReq()); err != nil {
			log.Print(err)
			return status.Error(codes.Aborted, err.Error())
		}
		txn.Commit()
		count++
	}
}

func (s *GRPCServer) GetEmployee(ctx context.Context, in *pb.EmployeeReq) (*pb.Employee, error) {
	res := new(pb.Employee)
	chMsg := make(chan struct{})
	txn := db.Txn(false)
	it, err := txn.Get(tableName, "id")
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	go func() {
		// Simulate some highload here
		time.Sleep(100 * time.Millisecond)

		for obj := it.Next(); obj != nil; obj = it.Next() {
			p := obj.(*pb.Employee)
			if p.Email == in.GetId() {
				res = p
				break
			}
		}
		chMsg <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil, err
	case <-chMsg:
		return res, nil
	}
}
