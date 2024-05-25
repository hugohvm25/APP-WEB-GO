/*

============================================
INTERRUPÇÃO DA APLICAÇÃO NO TERMINAL
CTRL+C
============================================

criar uma rota com a biblioteca http.ListenAndServe para que seja reconhecido pelo navegador


*/

package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	// criando a conexao com o banco de dados
	conexao := "user=postgres dbname=loja_hugo password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// criando uma estrutura para que sejam armazenados novos produtos de forma automatizada sem ter que ficar criando tabelas dentro do arquivo HTML individualmente
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// armazenando todos os templates pela variavel TEMP por encapsulamento onde na pasta templates o * irá trazer todos os arquivos .html
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	//teste de conexão
	// db := conectaComBancoDeDados()
	// defer db.Close()
	//sempre que tiver um / ela vai respoder para um segundo parametro - NÃO ESQUECER DO : PARA DETERMINAR A ROTA HTML
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// sempre que tiver uma requisição no site (index) vou poder escrever e mostrar
func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()
	selecProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	//instanciando 1 produto depois armazenar no slice para exibir na página com um laço de repetição FOR
	p := Produto{}
	produtos := []Produto{}

	for selecProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//vai scanear cada linha e armazenar as informações de cada produto no endereço &
		err = selecProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		//após o scaneamento vai inserir na variavel p cada dado para ser incluido no slice
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		//adicionar a listas de produtos um a um
		produtos = append(produtos, p)
	}
	temp.ExecuteTemplate(w, "Index", produtos)

	//obrigatoriamente fecha o banco de dados
	defer db.Close()

	//criando um Slice (array em go)
	// produtos := []Produto{
	// 	//2 formas diferentes de armazenar os dados no slice pela struct: passando os tipos e somente os dados diretamente
	// 	{Nome: "Camisa", Descricao: "Azul", Preco: 25, Quantidade: 4},
	// 	{"Bermuda", "Verde", 80, 2},
	// 	{"Tenis", "Preto", 100, 1},
	// 	{"Teste Produto", "Diferente", 1000, 1},
	// 	{"Teste Produto2", "Diferente", 1000, 1},
}
