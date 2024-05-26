package models

import "example.com/hello/go/src/APP-WEB-GO/db"

// criando uma estrutura para que sejam armazenados novos produtos de forma automatizada sem ter que ficar criando tabelas dentro do arquivo HTML individualmente
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		//adicionar a listas de produtos um a um
		produtos = append(produtos, p)
	}

	//obrigatoriamente fecha o banco de dados
	defer db.Close()
	return produtos

	//criando um Slice (array em go)
	// produtos := []Produto{
	// 	//2 formas diferentes de armazenar os dados no slice pela struct: passando os tipos e somente os dados diretamente
	// 	{Nome: "Camisa", Descricao: "Azul", Preco: 25, Quantidade: 4},
	// 	{"Bermuda", "Verde", 80, 2},
	// 	{"Tenis", "Preto", 100, 1},
	// 	{"Teste Produto", "Diferente", 1000, 1},
	// 	{"Teste Produto2", "Diferente", 1000, 1},
}

//a função recebe alguns parametros que foram mencionados no controler para adicionar novos produtos
func CriarNovoProduto(nome, descricao string, preco float64, quantidde int) {
	db := db.ConectaComBancoDeDados()

	//comando de inclusão no banco passando o comando de insert e fazendo a verificação dos valores de acordo com o tipo de dados aceito no banco
	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	//se não tiver algum erro, execute a inserção no banco
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidde)
	defer db.Close()
}

// recebe o parametro porém não faz diferença apesar de ser int pois será deletado
func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	//prepara o banco de dados para executar a função desejada de deletar o produto e a condição pelo ID do produto
	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	//executa a ação caso não tenha erro
	deletarOProduto.Exec(id)
	defer db.Close()
}
