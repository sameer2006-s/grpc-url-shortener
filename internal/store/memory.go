package store

type MemoryStore struct {
	links map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		links: make(map[string]string),
	}
}

func (s *MemoryStore) Save(
	code string,
	url string,
) {
	s.links[code] = url
}

func (s *MemoryStore) Get(
	code string,
) (string, bool) {

	url, ok := s.links[code]

	return url, ok
}