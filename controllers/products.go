package controllers

import (
	"go-web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}
func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Fail to convert price:", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Fail to convert amount:", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	productModel := models.GetProduct(productId)
	templates.ExecuteTemplate(w, "Edit", productModel)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		description := r.FormValue("description")
		price := r.FormValue("price")
		name := r.FormValue("name")
		amount := r.FormValue("amount")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during id conversion to int:", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error during price conversion to float64:", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error during amount conversion to int:", err)
		}
		models.UpdateProduct(convertedId, name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(w, r, "/", 301)
}
