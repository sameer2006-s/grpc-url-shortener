package main

import (
	"log"
	"net"

	internalgrpc "github.com/sameer2006-s/grpc-url-shortner/internal/grpc"
	"github.com/sameer2006-s/grpc-url-shortner/internal/repository"
	"github.com/sameer2006-s/grpc-url-shortner/internal/service"
	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"

	"google.golang.org/grpc"
)

func main() {

	// Transport
	lis, err := net.Listen(
		"tcp",
		":50051",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	repo :=
		repository.
			NewMemoryRepository()

	// Business logic
	linkService :=
		service.
			NewLinkService(
				repo,
			)

	// gRPC handler
	linkServer :=
		internalgrpc.
			NewLinkServer(
				linkService,
			)

	// Server
	grpcServer :=
		grpc.NewServer()

	pb.RegisterLinkServiceServer(
		grpcServer,
		linkServer,
	)

	log.Println(
		"link-service listening on :50051",
	)

	err =
		grpcServer.
			Serve(
				lis,
			)

	if err != nil {
		log.Fatal(err)
	}
}