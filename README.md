# Galt - a SaltStack alternative written in Go

# Rquirements

* Golang 1.19+
* Golang Cobra 
* Golang gRPC


# Install development requirements

```

git clone github.com/bbruun/galt
cd galt
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/protobuf/protoc-gen
``` 



# Galt usage

## Run the server

`galt server [options]`

## Run the client 

`galt client [options]`

## Run a galt command 

`galt <command...>`


