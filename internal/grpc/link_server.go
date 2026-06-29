package grpc

import (
	"context"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"
	"github.com/sameer2006-s/grpc-url-shortner/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LinkServer struct {
	pb.UnimplementedLinkServiceServer

	service *service.LinkService
}

func NewLinkServer(service *service.LinkService) *LinkServer {
	return &LinkServer{
		service: service,
	}
}

func (s *LinkServer) CreateLink(ctx context.Context, req *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	code, err := s.service.CreateLink(req.Url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create link: %v", err)
	}

	return &pb.CreateLinkResponse{ShortUrl: code}, nil
}

func (s *LinkServer) GetLink(ctx context.Context, req *pb.GetLinkRequest) (*pb.GetLinkResponse, error) {
	url, _ := s.service.GetLink(req.ShortUrl)
	return &pb.GetLinkResponse{Url: url}, nil
}

func (s *LinkServer) VisitLink(ctx context.Context, req *pb.VisitLinkRequest) (*pb.VisitLinkResponse, error) {
	redirectURL, err := s.service.VisitLink(req.ShortUrl)
	if err != nil {
		if err == service.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "link not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "visit link: %v", err)
	}
	return &pb.VisitLinkResponse{RedirectUrl: redirectURL}, nil
}

func (s *LinkServer) GetStats(ctx context.Context, req *pb.GetStatsRequest) (*pb.GetStatsResponse, error) {
	clicks, createdAt, err := s.service.GetStats(req.ShortUrl)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get stats: %v", err)
	}
	return &pb.GetStatsResponse{Clicks: int32(clicks), CreatedAt: createdAt}, nil
}
