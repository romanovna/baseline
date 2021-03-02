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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res := new(pb.Employee)
	txn := db.Txn(false)
	it, err := txn.Get(tableName, "id")
	if err != nil {
		panic(err)
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			p := obj.(*pb.Employee)
			if p.Email == in.GetId() {
				res = p
				break
			}
		}
	}
	return res, nil
}
