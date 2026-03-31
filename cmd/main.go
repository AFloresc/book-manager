package main

import (
	handler "book-manager/internal/adapters/handlers"
	"book-manager/internal/adapters/repository"
	"book-manager/internal/core/services"
	"net/http"
)

func main() {
	// 1. Elegimos nuestra persistencia (Salida)
	repo := repository.NewMemoryRepo()

	// 2. Creamos el servicio inyectando el repo (Núcleo)
	service := services.NewBookService(repo)

	// 3. Creamos el adaptador HTTP con el servicio (Entrada)
	api := handler.NewHTTPHandler(service)

	// 4. Servimos
	http.HandleFunc("/books", api.CreateBook)
	http.ListenAndServe(":8080", nil)
}
