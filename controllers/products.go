package controllers

import (
	"html/template"
	"net/http"
	"store/models"
	"strconv"
)

const REDIRECT = 301

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchProducts()
	temp.ExecuteTemplate(w, "index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		value, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		amount, _ := strconv.Atoi(r.FormValue("quantidade"))

		models.CreateNewProduct(name, description, value, amount)
	}

	http.Redirect(w, r, "/", REDIRECT)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", REDIRECT)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.EditProduct(id)
	temp.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		value, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		amount, _ := strconv.Atoi(r.FormValue("quantidade"))

		models.UpdateProduct(name, description, value, id, amount)
	}

	http.Redirect(w, r, "/", REDIRECT)
}
