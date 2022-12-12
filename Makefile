generate:
	protoc \
		--go_out=./gen/server \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/server \
		--go-grpc_opt=paths=source_relative \
		.proto
