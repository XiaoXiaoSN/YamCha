package storage

// Storage define interface
type Storage interface {
}

// Memory is a golang process memory value stague
type Memory struct {
	Data map[string]interface{}
}

// NewMemoryStorage return a memory storage
func NewMemoryStorage() Storage {
	return &Memory{}
}
