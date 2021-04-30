package jsonData

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Data struct {
// 	Name      string `json:"name"`
// 	Batch     int    `json:"batch"`
// 	BatchName string `json:"batch_name"`
// }

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Job    string `json:"job"`
	Social Social `json:"social"`
}

type Social struct {
	Instagram string `json:"instagram"`
	Whatsapp  int    `json:"whatsapp"`
}

func jsonData() {
	var usersData Users
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("success get file data.json")
	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteData, &usersData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, data := range usersData.Users {
		fmt.Println("name :", data.Name)
		fmt.Println("age :", data.Age)
		fmt.Println("job :", data.Job)
		fmt.Println("social :", data.Social)
	}

	// JSON data
	// var data Data

	// jsonString := `{
	// 	"name" : "impact Byte",
	// 	"batch" : 4,
	// 	"batch_name": "excellent echo"
	// }`

	// jsonData := []byte(jsonString)

	// err := json.Unmarshal(jsonData, &data)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(data)

	// kita bisa tampung json data ke dalam struct / map
	// kalau membuat web service, kita sering pakai struct
}
