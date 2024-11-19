package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"shopy/ui"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category_id string  `json:"category"`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getTopProducts(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	qry := "SELECT id, title, price, image,description, category_id FROM products"
	if limit != "" {
		qry = qry + " LIMIT " + limit
	}
	rows, err := db.Query(qry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Price, &p.Image, &p.Description, &p.Category_id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// getCategories devuelve las categorías únicas
func getCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,name  FROM category")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// getProductsByCategory devuelve productos según una categoría
func getProductsByCategory(w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]
	rows, err := db.Query("SELECT id, title, price, image,description, category_id FROM products WHERE category_id = ? ", category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Price, &p.Image, &p.Description, &p.Category_id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	// Conexión a la base de datos
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/shopy" // Reemplaza con tus credenciales

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Comprobación de conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error verificando la conexión:", err)
	}

	// Router
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/api/products/list-products", getTopProducts).Methods("GET")
	router.HandleFunc("/api/products/categories", getCategories).Methods("GET")
	router.HandleFunc("/api/products/category/{category}", getProductsByCategory).Methods("GET")

	// Archivos estáticos
	embetStatic, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}
	filesStatic := http.FileServer(http.FS(embetStatic))
	router.PathPrefix("/").Handler(filesStatic)
	// Archivos estáticos (carpeta `public`)
	publicDir := "./public"
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		fmt.Printf("Carpeta %s no encontrada, creándola...\n", publicDir)
		os.MkdirAll(publicDir, os.ModePerm)
	}
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	// http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	// Servidor
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", addCORSHeaders(router)))
}
func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
