package clstreamer

import (
	pb "clstream/internal/api"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"io"
	"log"
)

type GRPCServer struct{}

var (
	tableName = "employee"
)

func (s *GRPCServer) Stream(stream pb.CLStreamService_StreamServer) error {
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
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	var count uint

	for {
		f, err := stream.Recv()
		if err == io.EOF {
			//Read from existing db if any
			txn := db.Txn(false)
			it, err := txn.Get(tableName, "id")
			if err != nil {
				txn.Abort()
				log.Println(err)
			}
			for obj := it.Next(); obj != nil; obj = it.Next() {
				e := obj.(*pb.Employee)
				log.Println(e.GetEmail())
			}
			//log.Println("===")
			return stream.SendAndClose(&pb.StatusResponse{Result: fmt.Sprintf("%v entities written", count)})
		}
		// Create a write transaction
		txn := db.Txn(true)
		if err := txn.Insert(tableName, f.GetReq()); err != nil {
			log.Println(err)
		}
		txn.Commit()
		count++
	}
}
