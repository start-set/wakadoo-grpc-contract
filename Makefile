generate:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		wakadoo.proto
