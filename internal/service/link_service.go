package service

import (
	"fmt"

	"github.com/sameer2006-s/grpc-url-shortner/internal/model"
	"github.com/sameer2006-s/grpc-url-shortner/internal/repository"
)

type LinkService struct {
	repo repository.LinkRepository
}

func NewLinkService(
	repo repository.LinkRepository,
) *LinkService {

	return &LinkService{
		repo: repo,
	}
}

func (s *LinkService)CreateLink(url string,) string {
	code :=
		fmt.Sprintf(
			"link%d",
			len(url),
		)

	s.repo.Save(
		model.Link{
			ShortCode: code,
			URL: url,
		},
	)

	return code
}

func (s *LinkService)GetLink(code string,) (string, bool) {

	link, ok :=
		s.repo.Get(code)

	return link.URL, ok
}