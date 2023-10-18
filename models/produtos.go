package models

import (
	"github.com/piazin/store-go/db"
)

type Produto struct {
	Nome, Descricao string
	Preco float64
	Id, Quantidade int
}

func FindAllProducts() []Produto {
	db := db.ConnectToDatabase()

	selectAllProducts, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	
	p := Produto{}
	produtos := []Produto{}
	
	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectToDatabase()

	insertProduct, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}