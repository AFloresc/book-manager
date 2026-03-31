package repository

import (
	"book-manager/internal/core/domain"
	"errors"
)

type MemoryRepo struct {
	books map[string]domain.Book
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{books: make(map[string]domain.Book)}
}

func (m *MemoryRepo) Save(b domain.Book) error {
	m.books[b.ID] = b
	return nil
}

func (m *MemoryRepo) GetByID(id string) (domain.Book, error) {
	b, ok := m.books[id]
	if !ok {
		return domain.Book{}, errors.New("libro no encontrado")
	}
	return b, nil
}
