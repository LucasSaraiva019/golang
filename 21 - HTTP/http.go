package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE(requisição) - SERVIDOR(resposta)

	//Rotas identificar que ação deve ser feita
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuarios", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Carregar Página de Usuairos"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
