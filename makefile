genUserRPC:
	goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.