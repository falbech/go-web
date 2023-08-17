package models

import (
	"go-web/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func SearchAllProducts() []Product {
	db := db.ConnectToDb()
	selectAll, err := db.Query("select * from product order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAll.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectAll.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConnectToDb()

	preparedStatement, err := db.Prepare("insert into product(name,description,price,amount) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	preparedStatement.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectToDb()
	preparedStatement, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		panic(err.Error())
	}
	preparedStatement.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectToDb()
	productResult, err := db.Query("select * from product where id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	productModel := Product{}

	for productResult.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productResult.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		productModel.Id = id
		productModel.Name = name
		productModel.Price = price
		productModel.Amount = amount
		productModel.Description = description
	}
	defer db.Close()
	return productModel
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectToDb()

	statement, err := db.Prepare("update product set name=$1, description=$2, price=$3, amount=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	statement.Exec(name, description, price, amount, id)
	defer db.Close()
}
