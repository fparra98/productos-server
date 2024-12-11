package main

import (
	"backend/categoria"
	"backend/producto"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Permitir todos los orígenes, cambiar a url del frontend si es necesario
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Máximo tiempo de cacheo de la respuesta preflight
	}))

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Bienvenido!"))
	})

	r.Post("/productos/crear-producto", producto.CrearProducto)
	r.Get("/productos", producto.GetProductos)
	r.Get("/productos/{id}", producto.GetProducto)
	r.Put("/productos/{id}", producto.ActualizarProducto)
	r.Delete("/productos/{id}", producto.BorrarProducto)
	//categorias
	r.Get("/categorias", categoria.GetCategorias)

	http.ListenAndServe(":8080", r)
}
