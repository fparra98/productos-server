package categoria

import (
	"backend/bd"
	"encoding/json"
	"net/http"
)

type Categoria struct {
	Category_id   int    `json:"category_id" db:"category_id"`
	Category_name string `json:"category_name" db:"category_name"`
	Description   string `json:"description" db:"description"`
}

// Leer
func GetCategorias(w http.ResponseWriter, r *http.Request) {
	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	// creo un array de tipo Categoria
	var categorias []Categoria
	err = con.Select(&categorias, "SELECT  category_id, category_name, description FROM categories")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categorias)
}
