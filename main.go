package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"shopy/ui"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	embetStatic, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}

	filesStatic := http.FileServer(http.FS(embetStatic))

	// Sirviendo archivos estáticos
	router.PathPrefix("/").Handler(filesStatic)
	// Arrancando el servidor
	fmt.Println("server on port 8080")
	http.ListenAndServe(":8080", router)
}
