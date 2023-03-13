.DEFAULT_GOAL = run

build-grpc:
	@echo "Building protobuf"
	@protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/minion.proto

build:
	go build .

test:
	echo "Testing /minion.MinionService/RegisterMinion"
	echo '{"name":"minion1"}' | grpc-client-cli --insecure --address localhost:4505 --service minion.MinionService --method RegisterNewMinion
	echo '{"name":"minion2"}' | grpc-client-cli --insecure --address localhost:4505 --service minion.MinionService --method RegisterNewMinion
	echo '{"readytoreceive":false,"result":"","success":false}' | grpc-client-cli --insecure --address localhost:4505 --service minion.MinionService --method GetCommands

run: build
	go run . server