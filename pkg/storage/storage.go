package storage

type Storage struct {
	hashmap map[string]int
}

func (s *Storage) Set(key string, value int) {
	s.hashmap[key] = value
}

func (s *Storage) Get(key string) int {
	return s.hashmap[key]
}

func NewStorage() *Storage {
	return &Storage{hashmap: make(map[string]int)}
}
