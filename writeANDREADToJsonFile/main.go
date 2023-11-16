package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Title string `json:"title"`
}

const JSON_FILE_NAME = "people.json"

func main() {
	people := []Person{
		{Name: "Иван", Age: 30, Title: "Инженер"},
		{Name: "IFIFZNEN", Age: 25, Title: "Программист Go (Junior)"},
		{Name: "Зеня", Age: 18, Title: "Киберспортсмен"},
	}

	file, err := os.OpenFile(JSON_FILE_NAME, os.O_WRONLY|os.O_CREATE, 0644)
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

	fmt.Print(readJsonFile(JSON_FILE_NAME))
}

func readJsonFile(fileName string) *[]Person {
	bytesFromFile, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	strFromFile := string(bytesFromFile)
	reader := strings.NewReader(strFromFile)
	decoder := json.NewDecoder(reader)

	var PersonsFromFile []Person
	err = decoder.Decode(&PersonsFromFile)
	if err != nil {
		panic(err)
	}
	return &PersonsFromFile
}
