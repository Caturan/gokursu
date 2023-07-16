package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	Category    int     `json:"category"`
	UnitPrice   float64 `json:"UnitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil { // Defence Programming
		return nil, err
	}

	defer response.Body.Close() //Should close response whatever takes error

	// For read the body

	BodyBytes, _ := ioutil.ReadAll(response.Body)

	// Match the getted data to product struct

	var products []Product
	json.Unmarshal(BodyBytes, &products)
	return products, err

}

// Add a new product

func AddProduct() {
	product := Product{Id: 5, ProductName: "iPhone", Category: 0, UnitPrice: 899.99}
	jsonProduct, err := json.Marshal(product)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Println("Saved")
}

func AddProduct2() {
	product2 := Product{ProductName: "Ring", Category: 1, UnitPrice: 299.99}
	jsonProduct2, err := json.Marshal(product2)

	if err != nil {
		fmt.Println(err)
	}

	response2, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct2))

	if err != nil {
		fmt.Println(err)
	}

	defer response2.Body.Close()

	fmt.Println("Saved")
}

func AddProduct3() {
	product3 := Product{ProductName: "SmartWatch", Category: 0, UnitPrice: 349.99}
	jsonProduct3, err := json.Marshal(product3)

	if err != nil {
		fmt.Println(err)
	}

	response3, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct3))

	if err != nil {
		fmt.Println(err)
	}

	defer response3.Body.Close()

	// Check the addes function especially id point

	BodyBytes, _ := ioutil.ReadAll(response3.Body)

	var productResponse Product

	json.Unmarshal(BodyBytes, &productResponse)
	fmt.Println("Saved : ", productResponse)

}
