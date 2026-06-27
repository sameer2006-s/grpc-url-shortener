package repository

import (
	"github.com/sameer2006-s/grpc-url-shortner/internal/model"
)

type LinkRepository interface {
	Save(link model.Link)

	Get(
		code string,
	) (model.Link, bool)
}

type MemoryRepository struct {
	links map[string]model.Link
}

func NewMemoryRepository()*MemoryRepository {
	return &MemoryRepository{
		links: make(
			map[string]model.Link,
		),
	}
}

func (r *MemoryRepository)Save(link model.Link,) {
	r.links[
		link.ShortCode,
	] = link
}

func (r *MemoryRepository)Get(code string,) (model.Link, bool) {

	link, ok :=
		r.links[code]

	return link, ok
}