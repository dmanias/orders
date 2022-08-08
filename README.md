

# orders

initialize go project
go mod init https://github.com/pampatzoglou/orders


use grpc
https://grpc.io/docs/languages/go/quickstart/

Go plugins for the protocol compiler:

    Install the protocol compiler plugins for Go using the following commands:

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    Update your PATH so that the protoc compiler can find the plugins:

    $ export PATH="$PATH:$(go env GOPATH)/bin"

https://sahansera.dev/building-grpc-server-go/

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=. --experimental_allow_proto3_optional

#  install grpcurl  https://princepereira.medium.com/install-grpccurl-in-ubuntu-6ad71fd3ed31
go get github.com/fullstorydev/grpcurl/...
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
export PATH="$PATH:$(go env GOPATH)/bin"

# kafka
https://yusufs.medium.com/getting-started-with-kafka-in-golang-14ccab5fa26
sudo apt-get install build-essential
go get github.com/confluentinc/confluent-kafka-go/kafka

for librdkafka undefined reference to `__rawmemchr' -> CMD ["go", "run", "-tags", "musl", "."]
https://github.com/confluentinc/confluent-kafka-go/issues/454
https://github.com/confluentinc/confluent-kafka-go-example/blob/master/README.md