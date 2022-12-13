generate:
	protoc \
		--proto_path=. \
		--go_opt=paths=source_relative \
		--go_out=. \
		--go-grpc_out=. \
		wakadoo.proto
