.DEFAULT_GOAL = run

build-grpc:
	@echo "Building protobuf"
	@protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/minion.proto

build:
	go build .

test:
	echo "Testing if server returns output"
	./galt -m "*s*" cmd.run 'ls -l '

run: build
	go run . server