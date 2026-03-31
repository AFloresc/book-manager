package services

import (
	"book-manager/internal/core/domain"
	"book-manager/internal/core/ports"

	"errors"
)

type bookService struct {
	repo ports.BookRepository // Inyectamos la interfaz
}

func NewBookService(r ports.BookRepository) ports.BookService {
	return &bookService{repo: r}
}

func (s *bookService) CreateBook(title, author string) (domain.Book, error) {
	if title == "" {
		return domain.Book{}, errors.New("el título es obligatorio")
	}

	book := domain.Book{ID: "1", Title: title, Author: author} // ID hardcodeado para el ejemplo
	err := s.repo.Save(book)
	return book, err
}

func (s *bookService) FetchBook(id string) (domain.Book, error) {
	return s.repo.GetByID(id)
}
