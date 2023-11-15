package main

import (
	"os"
)

const (
	TEXT_FILE_NAME = "documents/pushkin_autumn.txt"
)

func main() {
	txt := "Унылая пора!\nОчей очерование!\nПриятна мне твоя прощальная краса -\nЛюблю я пышное природы увяданье,\nВ багрец и в золото одетые леса,\nВ их сенях ветра шум и свежее дыхание,\nИ мглой волнистою покрыты небеса, И редкий солнца луч, и первые морозы,\nИ отдалённые седой зимы угрозы."

	writeTextToFile(TEXT_FILE_NAME, txt)

	//textFromFile := readTextFromFile(TEXT_FILE_NAME)
	//fmt.Println(textFromFile)

	// os.Remove(TEXT_FILE_NAME) // удаление файла
}

// функция для записи текста в новый файл или в существующий.
func writeTextToFile(filename, str string) {
	byteToWrite := []byte(str)
	if err := os.WriteFile(filename, byteToWrite, 0644); err != nil {
		panic(err)
	}
}
