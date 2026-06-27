package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	internalgrpc "github.com/sameer2006-s/grpc-url-shortner/internal/grpc"
	"github.com/sameer2006-s/grpc-url-shortner/internal/store"
	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
)

func main() {

	lis, err :=
		net.Listen(
			"tcp",
			":50051",
		)

	if err != nil {
		log.Fatal(err)
	}

	server :=
		grpc.NewServer()

	storage :=
		store.NewMemoryStore()

	pb.RegisterLinkServiceServer(
		server,
		internalgrpc.NewLinkServer(
			storage,
		),
	)

	log.Println(
		"link-service :50051",
	)

	server.Serve(lis)
}