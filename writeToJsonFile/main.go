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

	// Открываем файл для записи
	file, err := os.Create("people.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Создаем JSON-кодировщик
	encoder := json.NewEncoder(file)

	// Кодируем данные и записываем в файл
	err = encoder.Encode(people)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл people.json")
}
