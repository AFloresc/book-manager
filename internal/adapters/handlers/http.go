package handler

import (
	"book-manager/internal/core/ports"
	"encoding/json"
	"net/http"
)

type HTTPHandler struct {
	service ports.BookService
}

func NewHTTPHandler(s ports.BookService) *HTTPHandler {
	return &HTTPHandler{service: s}
}

func (h *HTTPHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la decodificación del JSON...
	book, _ := h.service.CreateBook("El Quijote", "Cervantes")
	json.NewEncoder(w).Encode(book)
}
