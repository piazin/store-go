package models

import (
	"fmt"

	"github.com/piazin/store-go/db"
	"github.com/piazin/store-go/utils"
)

type Produto struct {
	Nome, Descricao string
	Preco float64
	Id, Quantidade int
}

func FindAllProducts() []Produto {
	db := db.ConnectToDatabase()

	selectAllProducts, err := db.Query("select * from produtos order by id asc")
	utils.CheckError(err)
	
	p := Produto{}
	produtos := []Produto{}
	
	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		utils.CheckError(err)

		p.Id = id
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
	utils.CheckError(err)

	insertProduct.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeleteProductById(id string) {
	db := db.ConnectToDatabase()
	deleteQuery := fmt.Sprintf("DELETE FROM produtos WHERE id = %s", id)

	db.QueryRow(deleteQuery)

	defer db.Close()
}

func FindProductById(productId string) Produto {
	db := db.ConnectToDatabase()

	var (
		id, quantidade int
		nome, descricao string
		preco float64
	)
	
	rows, err := db.Query("select * from produtos where id = $1", productId)
	utils.CheckError(err)

	for rows.Next() {
		err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		utils.CheckError(err)
	}
	
	product := Produto{Nome: nome, Descricao: descricao, Preco: preco, Quantidade: quantidade, Id: id}
	
	defer db.Close()

	return product
}

func UpdateProductById(id, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectToDatabase()

	updatedProduct, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	utils.CheckError(err)

	updatedProduct.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}