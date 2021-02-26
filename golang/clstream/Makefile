PROTOC_ALL_VER=1.34_4

build:
	docker build -t clstream/grpc-server .

tidy:
	go mod tidy

proto:
	docker pull namely/protoc-all:${PROTOC_ALL_VER}
	docker run -v `pwd`:/defs namely/protoc-all:${PROTOC_ALL_VER} -f api/proto/clstream.proto -l go

clean:
	docker rmi clstream/grpc-server:latest