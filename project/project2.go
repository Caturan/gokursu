package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllCategories() ([]Category, error) {
	response, err := http.Get("http://localhost:3000/categories")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	BodyBytes, _ := ioutil.ReadAll(response.Body)

	var categories []Category
	json.Unmarshal(BodyBytes, &categories)
	return categories, err

}

func AddCategory() {
	category := Category{Id: 2, CategoryName: "Office&Home"}
	jsonCategory, err := json.Marshal(category)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("http://localhost:3000/categories", "application/json;charset=utf-8", bytes.NewBuffer(jsonCategory))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Println("Saved")
}

func AddCategory2() {
	category2 := Category{CategoryName: "Textile"}
	jsoncategory2, err := json.Marshal(category2)

	if err != nil {
		fmt.Println(err)
	}

	response2, err := http.Post("http://localhost:3000/categories", "application/json;charset=utf-8", bytes.NewBuffer(jsoncategory2))

	if err != nil {
		fmt.Println(err)
	}

	defer response2.Body.Close()

	BodyBytes, _ := ioutil.ReadAll(response2.Body)

	var category2Response Category

	json.Unmarshal(BodyBytes, &category2Response)
	fmt.Println("Saved : ", category2Response)

}
