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
	"text/template"
)

// criando uma estrutura para que sejam armazenados novos produtos de forma automatizada sem ter que ficar criando tabelas dentro do arquivo HTML individualmente
type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// armazenando todos os templates pela variavel TEMP por encapsulamento onde na pasta templates o * irá trazer todos os arquivos .html
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	//sempre que tiver um / ela vai respoder para um segundo parametro - NÃO ESQUECER DO : PARA DETERMINAR A ROTA HTML
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// sempre que tiver uma requisição no site (index) vou poder escrever e mostrar
func index(w http.ResponseWriter, r *http.Request) {
	//criando um Slice (array em go)
	produtos := []Produto{
		//2 formas diferentes de armazenar os dados no slice pela struct: passando os tipos e somente os dados diretamente
		{Nome: "Camisa", Descricao: "Azul", Preco: 25, Quantidade: 4},
		{"Bermuda", "Verde", 80, 2},
		{"Tenis", "Preto", 100, 1},
		{"Teste Produto", "Diferente", 1000, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
