package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/piazin/store-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, _ := strconv.ParseFloat(preco, 64)
		quantidadeConvertida, _ := strconv.ParseInt(quantidade, 10, 64)

		models.CreateNewProduct(nome, descricao, precoConvertido, int(quantidadeConvertida))
		
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProductById(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindProductById(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, _ := strconv.ParseFloat(preco, 64)
		quantidadeConvertida, _ := strconv.ParseInt(quantidade, 10, 64)

		models.UpdateProductById(id, nome, descricao, precoConvertido, int(quantidadeConvertida))
		
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}