generate:
	protoc \
		-I. \
		--go_out=. \
		--go-grpc_out=. \
		wakadoo.proto proto/*.proto

last-tag:
	git tag | sort -V | tail -1