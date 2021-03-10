PROTOC_ALL_VER=1.34_4

build: build-gw build-serv

build-serv:
	docker build -t clstream/grpc-server .

build-gw:
	docker build -t clstream-grpc-gw gen/grpc-gateway/

tidy:
	go mod tidy

loadtest:
	ghz --call=clstream.CLStreamService/Stream --insecure -c 1000 -n 10000 --connections=100 -D req.json localhost:8080

proto:
	docker pull namely/protoc-all:${PROTOC_ALL_VER}
	docker run -v `pwd`:/defs namely/protoc-all:${PROTOC_ALL_VER} -f api/proto/clstream.proto -l go

gw:
	docker run -v `pwd`:/defs namely/gen-grpc-gateway -f api/proto/clstream.proto -s CLStreamService

clean:
	docker rmi clstream/grpc-server:latest
	docker rmi clstream/grpc-gw:latest