package service

import (
	"github.com/google/uuid"

	"github.com/sameer2006-s/grpc-url-shortner/internal/model"
	"github.com/sameer2006-s/grpc-url-shortner/internal/repository"
)

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) CreateLink(url string) (string, error) {
	code := "link-" + uuid.NewString()[:8]

	err := s.repo.Save(model.Link{ShortCode: code, URL: url})
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *LinkService) GetLink(code string) (string, bool) {
	link, ok := s.repo.Get(code)
	return link.URL, ok
}
