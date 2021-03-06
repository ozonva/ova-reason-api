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
generate_proto_wsl:
	mkdir -p pkg/ova-reason-api
	protoc   -I vendor.protogen --proto_path=api \
			--go_out=pkg/ova-reason-api --go_opt=paths=import \
			--go-grpc_out=pkg/ova-reason-api --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg/ova-reason-api \
            --grpc-gateway_opt=logtostderr=true \
            --grpc-gateway_opt=paths=import \
            --swagger_out=allow_merge=true,merge_file_name=api:swagger \
			ova-reason-api.proto
	mv pkg/ova-reason-api/github.com/ozonva/ova-reason-api/pkg/ova-reason-api/* pkg/ova-reason-api/
	rm -rf pkg/ova-reason-api/github.com

.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init github.com/ozonva/ova-reason-api
    GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
    GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
    GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
    GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
    GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger


DBSTRING:="postgres://mish:mish@localhost:5432/reasonstorage?sslmode=disable"
.PHONY: goose-migrate
goose-migrate:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migrations status
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir migrations up