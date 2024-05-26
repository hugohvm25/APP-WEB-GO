package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// para que a função ou estrutura seja publica precisamso colocar a 1 letra sempre maiúscula
func ConectaComBancoDeDados() *sql.DB {
	// criando a conexao com o banco de dados
	conexao := "user=postgres dbname=loja_hugo password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
