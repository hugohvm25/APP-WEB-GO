/*

============================================
INTERRUPÇÃO DA APLICAÇÃO NO TERMINAL
CTRL+C
============================================

criar uma rota com a biblioteca http.ListenAndServe para que seja reconhecido pelo navegador

*/

package main

import (
	"net/http"

	"example.com/hello/go/src/APP-WEB-GO/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
