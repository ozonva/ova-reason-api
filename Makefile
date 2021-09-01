.PHONY: build
build:
	go build -o ./bin/main.exe ./cmd/ova-reason-api/main.go
.PHONY: run
run:	
	go run ./cmd/ova-reason-api/main.go

BUILD_FILENAME = ova-reason-api

.PHONY: test
test:
	ginkgo -r

.PHONY: generate_proto
generate_proto:
	rmdir pkg /s /q
	mkdir pkg\ova-reason-api
	C:\proto\bin\protoc \
			--go_out=pkg\ova-reason-api --go_opt=paths=import \
			--go-grpc_out=pkg\ova-reason-api --go-grpc_opt=paths=import \
			api\ova-reason-api.proto
	move pkg\ova-reason-api\github.com\ozonva\ova-reason-api\pkg\ova-reason-api\* pkg\ova-reason-api\

