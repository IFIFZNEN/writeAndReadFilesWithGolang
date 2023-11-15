package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Title string `json:"title"`
}

func main() {
	people := Person{Name: "Иван", Age: 30, Title: "Инженер"}

	file, err := os.Create("people.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(people)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл people.json")
}
