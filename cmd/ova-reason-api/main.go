package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozonva/ova-reason-api/internal/server"
	api "github.com/ozonva/ova-reason-api/pkg/ova-reason-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	grpcPort           = ":8080"
	grpcServerEndpoint = "localhost:8080"
)

func main() {

	go runJSON()
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterReasonRpcServer(s, server.NewReasonRpcServer(&log))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := api.RegisterReasonRpcHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
