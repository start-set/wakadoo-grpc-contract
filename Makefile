generate:
	protoc \
		-I. \
		--go_out=gen \
		--go-grpc_out=gen \
		wakadoo.proto proto/*.proto

last-tag:
	git tag | sort -V | tail -1