proto-gen:
	protoc -I ./grpc/proto --go-grpc_out=./grpc/proto --go_out=./grpc/proto ./grpc/proto/*.proto