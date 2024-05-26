package controlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"example.com/hello/go/src/APP-WEB-GO/models"
)

// armazenando todos os templates pela variavel TEMP por encapsulamento onde na pasta templates o * irá trazer todos os arquivos .html
var temp = template.Must(template.ParseGlob("templates/*.html"))

// sempre que tiver uma requisição no site (index) vou poder escrever e mostrar
func Index(w http.ResponseWriter, r *http.Request) {
	//retornar as informações da função para buscar todos os produtos da outra pasta models e exibir no site
	todosOsProdutos := models.BuscaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	//executar o template na nova página NEW
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	//se a requisição usar o metodo POST
	if r.Method == "POST" {
		//cria a variavel e busca o input do usuário no formulário preenchido
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//como os dados do preço está em float e o preenchido no formulario apesar de passar o tipo numerico, ele retorna em string e é necessário converter
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro de conversão do preço", err)
		}
		//atoi é a função para converter int para string
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro de conversão do quantidade", err)
		}

		//essa parte de inclusão efetivamente não é função do Controler e sim do Models, então vamos direcionar os dados já prontos para o arquivo models
		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	//ao terminar de preencher o formulário, a página deve atualizar e retornar para a página principal
	http.Redirect(w, r, "/", 301)
}
