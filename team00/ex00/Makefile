install_requirements:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin/protoc-gen-go"
	export PATH="$PATH:$(go env GOPATH)/bin/protoc-gen-go-grpc"

gen_proto:
	protoc -I proto proto/transmitter.proto --go_out=./pkg/transmitter/ --go_opt=paths=source_relative --go-grpc_out=./pkg/transmitter/ --go-grpc_opt=paths=source_relative
