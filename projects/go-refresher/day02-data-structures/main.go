package main
import "fmt"

type Store interface {
	Get(key string) (string, bool)
	Set(Key string, value string)
	Delete(key string) bool
}

type MemoryStore struct {
	data: map[string]string
}

func NewMemoryStore() *MemoryStore{
	return &MemoryStore{
		data: make(map[string]string),
	}
}

