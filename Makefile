clean:
	rm -rf ./pkg/*
	rm -rf ./api/*.lock

build:
	buf build --path api/user_management.proto
	cd api && buf mod update

generate: clean build
	. resource/.gopath && cd api && buf generate
	. resource/.gopath && cd pkg && protoc-go-inject-tag -input="*.pb.go"

run :
	. resource/.gopath && export ENV=${profile} && cd cmd && go run main.go

deploy: generate
	go build -v -o dot-indonesia cmd/main.go
	export ENV=${profile} && ./dot-indonesia

deploy-container:
	profile=${profile} docker-compose down --rmi all
	profile=${profile} docker-compose up -d

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/bufbuild/buf/cmd/buf \
		github.com/favadi/protoc-go-inject-tag \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
