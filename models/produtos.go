package models

import "EduardoBotelho/db"


type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}
func BuscaTodos()[]Produto{
	db := db.conectaDb()

	selectDeTodosOsProdutos, err := db.Query("Select * from produtos")
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
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) //adiciona o produto na lista
	}
	defer db.Close()
	return produtos
}