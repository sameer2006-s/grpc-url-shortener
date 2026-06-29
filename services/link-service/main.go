package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
	internalgrpc "github.com/sameer2006-s/grpc-url-shortner/internal/grpc"
	"github.com/sameer2006-s/grpc-url-shortner/internal/repository"
	"github.com/sameer2006-s/grpc-url-shortner/internal/service"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	db := connectDB(ctx)
	defer db.Close(ctx)

	repo := repository.NewPostgresRepository(db)
	svc := service.NewLinkService(repo)
	handler := internalgrpc.NewLinkServer(svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterLinkServiceServer(server, handler)

	log.Println("link-service listening on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve grpc: %v", err)
	}
}

func connectDB(ctx context.Context) *pgx.Conn {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		requiredEnv("DB_USER"),
		requiredEnv("DB_PASSWORD"),
		requiredEnv("DB_HOST"),
		requiredEnv("DB_PORT"),
		requiredEnv("DB_NAME"),
	)

	db, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}

	return db
}

func requiredEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("%s is required", key)
	}

	return value
}
