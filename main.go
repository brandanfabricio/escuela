package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"shopy/ui"
	"strconv"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite"
)

var db *sql.DB

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category_id int     `json:"category_id"`
	Name        string  `json:"category"`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getTopProducts(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	// qry := "SELECT id, title, price, image,description, category_id FROM products"
	qry := "select    products.id, title, price, image,description, category_id,name from category  left join products  on category_id = products.category_id "
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
		if err := rows.Scan(&p.ID, &p.Title, &p.Price, &p.Image, &p.Description, &p.Category_id, &p.Name); err != nil {
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
func handleFileUpload(w http.ResponseWriter, r *http.Request) {

	// Limitar el tamaño máximo de la solicitud (por ejemplo, 20 MB)
	r.Body = http.MaxBytesReader(w, r.Body, 20<<20) // 20 MB

	// Parsear la solicitud multipart/form-data
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(w, "Error al procesar la solicitud: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener los campos del formulario
	category := r.FormValue("category")
	title := r.FormValue("title")
	description := r.FormValue("description")
	price := r.FormValue("price")

	// Validar campos requeridos
	if category == "" || title == "" || description == "" || price == "" {
		http.Error(w, "Todos los campos son requeridos", http.StatusBadRequest)
		return
	}
	num, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Número como int:", num)
	}
	newPrice := num
	// Obtener el archivo subido
	file, header, err := r.FormFile("files")

	if err != nil {
		http.Error(w, "Error al obtener el archivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// // Crear el directorio de subida si no existe
	uploadDir := "./public/uploads/"
	err = ensureDir(uploadDir)
	if err != nil {
		http.Error(w, "Error al crear el directorio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// // Guardar el archivo en el servidor
	filePath := uploadDir + header.Filename
	destFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error al guardar el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, "Error al copiar el archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// // Aquí puedes guardar los datos en la base de datos
	fmt.Printf("Nuevo producto:\nCategory: %s\nTitle: %s\nDescription: %d\nPrice: %s\nImage Path: %s\n", category, title, description, newPrice, filePath)

	newProduct := Product{
		Title:       title,
		Price:       newPrice,
		Description: description,
		Image:       filePath,
		Category_id: 1,
	}
	// Insertar producto en la base de datos
	insertProductQuery := `
		INSERT INTO products ( title, price,  description, image,category_id) 
		VALUES (?, ?, ?, ?,?);`

	_, err = db.Exec(insertProductQuery, newProduct.Title, newProduct.Price, newProduct.Description, newProduct.Image, newProduct.Category_id)
	if err != nil {
		log.Fatalf("Error insertando producto: %v", err)
	}

	fmt.Println("Producto insertado exitosamente")

	// Responder al cliente
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Producto subido exitosamente",
		"path":    filePath,
	}

	json.NewEncoder(w).Encode(response)
}

func delecteProdcuto(w http.ResponseWriter, r *http.Request) {
	// Insertar producto en la base de datos
	qryDelete := `DELETE FROM products WHERE id = ?;`

	id := mux.Vars(r)["id"]

	fmt.Println(id)

	_, err := db.Exec(qryDelete, id)
	if err != nil {
		log.Fatalf("Error insertando producto: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Producto Eliminado exitosamente",
		"id":      id,
	}

	json.NewEncoder(w).Encode(response)
}
func cleanProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aki")
	id := mux.Vars(r)["id"]


	
	response := map[string]string{
		"message": "Producto Eliminado exitosamente",
		"id":      id,
	}

	json.NewEncoder(w).Encode(response)
}

// ensureDir verifica si un directorio existe, y si no, lo crea
func ensureDir(dir string) error {

	// Verifica si el directorio ya existe
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Crea el directorio con permisos 0755
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
func main() {
	// Conexión a la base de datos
	var err error
	// dsn := "root:root@tcp(127.0.0.1:3306)/shopy" // Reemplaza con tus credenciales
	dbpath := "shopy.db"
	//	db, err = sql.Open("mysql", dsn)

	db, err = sql.Open("sqlite", dbpath)

	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Comprobación de conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error verificando la conexión:", err)
	}
	// Crear tablas si no existen
	createTables()

	// Router
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/api/products/list-products", getTopProducts).Methods("GET")
	router.HandleFunc("/api/products/categories", getCategories).Methods("GET")
	router.HandleFunc("/api/products/category/{category}", getProductsByCategory).Methods("GET")
	router.HandleFunc("/api/products/{id}", delecteProdcuto).Methods("DELETE")
	router.HandleFunc("/api/products/{id}", cleanProduct).Methods("PUT")
	router.HandleFunc("/api/new-products", handleFileUpload).Methods("POST")

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
func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS category (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			price REAL NOT NULL,
			description TEXT DEFAULT NULL,
			image TEXT DEFAULT NULL,
			category_id INTEGER NOT NULL,
			FOREIGN KEY(category_id) REFERENCES category(id)
		)`,
	}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Fatalf("Error creando tablas: %v", err)
		}
	}
}
