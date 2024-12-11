package producto

import (
	"backend/bd"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Producto struct {
	Producto_id     int     `json:"product_id" db:"product_id"`
	Producto_Nombre string  `json:"product_name" db:"product_name"`
	Categoria       int     `json:"category_id" db:"category_id"`
	Unidad          string  `json:"unit" db:"unit"`
	Precio          float64 `json:"price" db:"price"`
}

//Crear

func CrearProducto(w http.ResponseWriter, r *http.Request) {
	//creo la viriable producto de tipo producto
	var producto Producto

	// decodifico la solicitud de vue en un json y lo almaceno en producto
	err := json.NewDecoder(r.Body).Decode(&producto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//me conecto a la base de datos
	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	//preparo la consulta que realizaré
	query := `INSERT INTO products (product_name, category_id, unit, price ) VALUES (:product_name, :category_id, :unit, :price)`
	//ejecuto la consulta
	_, err = con.NamedExec(query, producto)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(producto)
}

// Leer
func GetProductos(w http.ResponseWriter, r *http.Request) {
	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()
	// creo un array de tipo Prodcutos
	var productos []Producto
	err = con.Select(&productos, "SELECT  product_id, product_name, category_id, unit, price FROM products")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

// leer Producto
func GetProducto(w http.ResponseWriter, r *http.Request) {
	var id string = chi.URLParam(r, "id")
	var producto Producto

	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	err = con.Get(&producto, "SELECT product_id, product_name, category_id, unit, price FROM products WHERE product_id=$1", id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(producto)
}

// Editar
func ActualizarProducto(w http.ResponseWriter, r *http.Request) {
	var producto Producto
	var id string = chi.URLParam(r, "id")

	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	err = json.NewDecoder(r.Body).Decode(&producto)
	if err != nil {
		panic(err)
	}
	query := `UPDATE products SET product_name= :product_name, category_id = :category_id, unit = :unit, price =:price WHERE product_id= :product_id`
	_, err = con.NamedExec(query, map[string]interface{}{
		"product_id":   id,
		"product_name": producto.Producto_Nombre,
		"category_id":  producto.Categoria,
		"unit":         producto.Unidad,
		"price":        producto.Precio,
	})
	if err != nil {
		panic(err)

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(producto)
}

// Borrar
func BorrarProducto(w http.ResponseWriter, r *http.Request) {
	con, err := bd.ConDB()
	if err != nil {
		panic(err)
	}
	defer con.Close()

	id := chi.URLParam(r, "id")
	query := `DELETE FROM products WHERE product_id= :id`
	_, err = con.NamedExec(query, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
