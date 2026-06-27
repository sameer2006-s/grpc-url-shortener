package grpc

import (
	"context"
	"fmt"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
	"github.com/sameer2006-s/grpc-url-shortner/internal/store"
)

type LinkServer struct {
	pb.UnimplementedLinkServiceServer

	store *store.MemoryStore
}

func NewLinkServer(
	store *store.MemoryStore,
) *LinkServer {

	return &LinkServer{
		store: store,
	}
}

func (s *LinkServer) CreateLink(
	ctx context.Context,
	req *pb.CreateLinkRequest,
) (*pb.CreateLinkResponse, error) {

	code := fmt.Sprintf(
		"link%d",
		len(req.Url),
	)

	s.store.Save(
		code,
		req.Url,
	)

	return &pb.CreateLinkResponse{
		ShortUrl: code,
	}, nil
}

func (s *LinkServer) GetLink(
	ctx context.Context,
	req *pb.GetLinkRequest,
) (*pb.GetLinkResponse, error) {

	url, _ :=
		s.store.Get(
			req.ShortUrl,
		)

	return &pb.GetLinkResponse{
		Url: url,
	}, nil
}
