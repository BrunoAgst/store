package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Value       float64
	Amount      int
}

func SearchProducts() []Product {
	db := db.ConectDatabase()
	allProducts, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, amount int
		var name, description string
		var value float64

		err := allProducts.Scan(&id, &name, &description, &value, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Value = value
		p.Amount = amount

		products = append(products, p)

	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, value float64, amount int) {
	db := db.ConectDatabase()

	newProduct, err := db.Prepare("insert into produtos(nome, descricao, valor, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	newProduct.Exec(name, description, value, amount)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectDatabase()

	delete, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConectDatabase()

	product, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	dataProduct := Product{}

	for product.Next() {
		var id, amount int
		var name, description string
		var value float64

		err := product.Scan(&id, &name, &description, &value, &amount)

		if err != nil {
			panic(err.Error())
		}
		dataProduct.Id = id
		dataProduct.Name = name
		dataProduct.Description = description
		dataProduct.Value = value
		dataProduct.Amount = amount
	}

	defer db.Close()

	return dataProduct
}

func UpdateProduct(name, description string, value float64, id, amount int) {
	db := db.ConectDatabase()

	update, err := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, valor=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, amount, value, id)

	defer db.Close()
}
