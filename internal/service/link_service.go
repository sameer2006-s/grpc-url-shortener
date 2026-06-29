package service

import (
	"github.com/google/uuid"
	"time"

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

	err := s.repo.Save(model.Link{ShortCode: code, URL: url, CreatedAt: time.Now(), Clicks: 0})

	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *LinkService) GetLink(code string) (string, bool) {
	link, ok := s.repo.Get(code)
	return link.URL, ok
}

func (s *LinkService) VisitLink(code string) (string, error) {
	link, ok := s.repo.Get(code)

	if !ok {
		return "", ErrNotFound
	}

	err := s.repo.IncrementClicks(code)

	if err != nil {
		return "", err
	}

	return link.URL, nil
}

func (s *LinkService) GetStats(code string) (int, string, error) {
	link, ok := s.repo.Get(code)

	if !ok {
		return 0, "", ErrNotFound
	}

	return link.Clicks, link.CreatedAt.Format(time.RFC3339), nil
}
