generate:
	protoc \
		-I. \
		--go_out=. \
		--go-grpc_out=. \
		wakadoo.proto proto/*.proto
	protoc -I=. \
	--js_out=import_style=commonjs,binary:. \
	--grpc-web_out=import_style=typescript,mode=grpcwebtext:. \
	wakadoo.proto proto/*.proto

last-tag:
	git tag | sort -V | tail -1