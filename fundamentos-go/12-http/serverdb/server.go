package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHanlder)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
