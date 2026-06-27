package grpc

import (
	"context"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
	"github.com/sameer2006-s/grpc-url-shortner/internal/service"
)

type LinkServer struct {
	pb.UnimplementedLinkServiceServer

	service *service.LinkService
}

func NewLinkServer(service *service.LinkService,) *LinkServer {
	return &LinkServer{
		service: service,
	}
}

func (s *LinkServer)CreateLink(ctx context.Context,req *pb.CreateLinkRequest,) (*pb.CreateLinkResponse,error,) {
	code :=
		s.service.
			CreateLink(
				req.Url,
			)

	return &pb.CreateLinkResponse{
		ShortUrl: code,
	}, nil
}

func (s *LinkServer)GetLink(ctx context.Context,req *pb.GetLinkRequest,) (*pb.GetLinkResponse,error,) {
	url, _ :=
		s.service.
			GetLink(
				req.ShortUrl,
			)

	return &pb.GetLinkResponse{
		Url: url,
	}, nil
}