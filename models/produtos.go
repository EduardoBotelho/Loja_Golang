package models

import (
	"github.com/EduardoBotelho/db"

)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodos() []Produto {
	db := db.ConectaDb()

	selectDeTodosOsProdutos, err := db.Query("Select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}                       //cria uma variável do tipo Produto
	produtos := []Produto{}              //
	for selectDeTodosOsProdutos.Next() { //enquanto houver linhas
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id,
			&nome, &descricao,
			&preco, &quantidade) //lê os dados do banco
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) //adiciona o produto na lista
	}
	defer db.Close()
	return produtos
}
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDb()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao,preco,quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func DeletaProduto(id string) {
	db := db.ConectaDb()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}
func EditaProduto(id string) Produto {
	conn := db.ConectaDb()

	produtoDoBanco, err := conn.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer conn.Close()
	return produtoParaAtualizar
}
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	conn := db.ConectaDb()

	AtualizaProduto, err := conn.Prepare("update produtos set nome=$1, descricao=$2,preco=$3,quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer conn.Close()
}
