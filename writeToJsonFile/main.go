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
	person := Person{
		Name:  "Иван",
		Age:   30,
		Title: "Инженер",
	}
	// Открываем файл для записи
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}
	fmt.Println("Данные успешно записаны в файл person.json")
}
