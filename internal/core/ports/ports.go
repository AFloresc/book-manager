package ports

import "book-manager/internal/core/domain"

// Puerto de Salida: Lo que el dominio necesita para persistir datos
type BookRepository interface {
	Save(book domain.Book) error
	GetByID(id string) (domain.Book, error)
}

// Puerto de Entrada: Lo que el mundo exterior puede pedirle al dominio
type BookService interface {
	CreateBook(title, author string) (domain.Book, error)
	FetchBook(id string) (domain.Book, error)
}
