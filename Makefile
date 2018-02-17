generate_go:
	protoc \
		--go_out=plugins=grpc:$$GOPATH/src \
		--grpc-gateway_out=logtostderr=true:$$GOPATH/src \
		--swagger_out=logtostderr=true:./go \
		-I/usr/local/include \
        -I./go \
        -I./vendor/github.com/grpc-ecosystem/grpc-gateway/ \
        -I./vendor/ \
		./go/go.proto

generate_gofast:
	# Need to generate custom timestamp type
	protoc \
		--gofast_out=plugins=grpc:$$GOPATH/src \
		-I./gofast \
		./gofast/google/protobuf/timestamp.proto
	protoc \
		--gofast_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/johanbrandhorst/gogoproto-experiments/gofast/timestamp:$$GOPATH/src \
		--grpc-gateway_out=logtostderr=true,Mgoogle/protobuf/timestamp.proto=github.com/johanbrandhorst/gogoproto-experiments/gofast/timestamp:$$GOPATH/src \
		--swagger_out=logtostderr=true:./gofast \
		-I/usr/local/include \
        -I./gofast \
        -I./vendor/github.com/grpc-ecosystem/grpc-gateway/ \
        -I./vendor/ \
		./gofast/gofast.proto

generate_gogo:
	protoc \
		--gogo_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:$$GOPATH/src \
		--grpc-gateway_out=logtostderr=true,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api:$$GOPATH/src \
		--swagger_out=logtostderr=true:./gogo \
		-I/usr/local/include \
        -I./gogo \
        -I./vendor/github.com/gogo/googleapis/ \
        -I./vendor/github.com/grpc-ecosystem/grpc-gateway/ \
        -I./vendor/ \
		./gogo/gogo.proto
	# Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229
	 sed -i 's/empty.Empty/types.Empty/g' ./gogo/server/gogo.pb.gw.go

generate: generate_go generate_gofast generate_gogo

install:
	go install \
		./vendor/github.com/golang/protobuf/protoc-gen-go \
		./vendor/github.com/gogo/protobuf/protoc-gen-gofast \
		./vendor/github.com/gogo/protobuf/protoc-gen-gogo \
		./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
