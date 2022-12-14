generate:
	protoc \
		-I. \
		--go_out=. \
		--go-grpc_out=. \
		wakadoo.proto proto/*.proto
	npx grpc_tools_node_protoc \
        --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
        --ts_out=. \
        -I . \
        wakadoo.proto proto/*.proto

last-tag:
	git tag | sort -V | tail -1