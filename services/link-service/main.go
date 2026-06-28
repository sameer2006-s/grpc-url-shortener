package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
	internalgrpc "github.com/sameer2006-s/grpc-url-shortner/internal/grpc"
	"github.com/sameer2006-s/grpc-url-shortner/internal/repository"
	"github.com/sameer2006-s/grpc-url-shortner/internal/service"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("connect to postgres: %v", err)
	}

	_, err = db.Exec(ctx, `
CREATE TABLE IF NOT EXISTS links (
	id UUID PRIMARY KEY,
	short_code TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL
)
`)
	if err != nil {
		log.Fatalf("create links table: %v", err)
	}
	defer db.Close(ctx)

	repo := repository.NewPostgresRepository(db)
	svc := service.NewLinkService(repo)
	handler := internalgrpc.NewLinkServer(svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen on :50051: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterLinkServiceServer(server, handler)

	log.Println("started :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve grpc: %v", err)
	}
}
