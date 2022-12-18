package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func allArticulos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EndPoint Solicitado: HomePage")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	//LLamada sin Parámetros
	myRouter.HandleFunc("/", homePage)
	//Mostrar todos los Artículos
	myRouter.HandleFunc("/articulos", allArticulos)
	//Crear un nuevo Artículo
	myRouter.HandleFunc("/articulo", createNewArticle).Methods("POST")
	//Borrar un Artículo
	myRouter.HandleFunc("/articulo/{id}", deleteArticle).Methods("DELETE")
	//Actualizar un Artículo
	myRouter.HandleFunc("/articulo/{id}", updateArticle).Methods("PUT")
	//Buscar un Artículo
	myRouter.HandleFunc("/articulo/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {

	//Manejando las Peticiones
	handleRequest()

}
