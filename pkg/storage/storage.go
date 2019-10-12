package storage

// Storage define interface
type Storage interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	GetInt(key string) int
}

// Memory is a golang process memory value stague
type Memory struct {
	Data map[string]interface{}
}

// NewMemoryStorage return a memory storage
func NewMemoryStorage() Storage {
	return &Memory{}
}

// Set implement Storage Set
func (s *Memory) Set(key string, value interface{}) error {
	return nil
}

// Get implement Storage Get
func (s *Memory) Get(key string) (interface{}, error) {
	return nil, nil
}

// GetInt implement Storage Get
func (s *Memory) GetInt(key string) int {
	return 0
}
