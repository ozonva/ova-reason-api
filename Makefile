.PHONY: build
build:
	go build -o ./bin/main.exe ./cmd/ova-reason-api/main.go
.PHONY: run
run:	
	go run ./cmd/ova-reason-api/main.go

.PHONY: test
test:
	ginkgo -r

.PHONY: generate_proto_win
generate_proto:
	rmdir pkg /s /q
	mkdir pkg\ova-reason-api
	C:\proto\bin\protoc \
			--go_out=pkgova-reason-api --go_opt=paths=import \
			--go-grpc_out=pkg\ova-reason-api --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg\ova-reason-api \
            --grpc-gateway_opt=logtostderr=true \
            --grpc-gateway_opt=paths=import \
            --swagger_out=allow_merge=true,merge_file_name=api:swagger \
			api\ova-reason-api.proto
	move pkg\ova-reason-api\github.com\ozonva\ova-reason-api\pkg\ova-reason-api\* pkg\ova-reason-api\

.PHONY: generate_proto_wsl
generate_proto2:
	mkdir -p pkg/ova-reason-api
	protoc  \
			--go_out=pkg/ova-reason-api --go_opt=paths=import \
			--go-grpc_out=pkg/ova-reason-api --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg/ova-reason-api \
            --grpc-gateway_opt=logtostderr=true \
            --grpc-gateway_opt=paths=import \
            --swagger_out=allow_merge=true,merge_file_name=api:swagger \
			api/ova-reason-api.proto
	mv pkg/ova-reason-api/github.com/ozonva/ova-reason-api/pkg/ova-reason-api/* pkg/ova-reason-api/
	rm -rf pkg/ova-reason-api/github.com
