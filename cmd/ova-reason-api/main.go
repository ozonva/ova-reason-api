package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-reason-api/internal/repo"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/ozonva/ova-reason-api/internal/server"
	api "github.com/ozonva/ova-reason-api/pkg/ova-reason-api"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

const (
	grpcPort           = ":8080"
	grpcServerEndpoint = "localhost:8080"
	gatewayPort        = ":8081"
)

func main() {

	go runJSON()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sqlx.Open("pgx", dsn) // *sql.DB
	if err != nil {
		log.Fatalf("failed to load driver: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	// работаем с db

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	repo := repo.NewReasonRepository(db)
	api.RegisterReasonRpcServer(s, server.NewReasonRpcServer(&repo))
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
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

	err = http.ListenAndServe(gatewayPort, mux)
	if err != nil {
		panic(err)
	}
}
