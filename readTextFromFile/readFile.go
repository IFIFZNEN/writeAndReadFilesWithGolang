package main

import (
	"fmt"
	"os"
)

func readTextFromFile(fileName string) string {
	str, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(str)
}

const (
	TEXT_FILE_NAME = "C:\\Users\\Artykov\\Documents\\GitHub\\writeAndReadFilesWithGolang\\writeToFileStrings\\documents\\pushkin_autumn.txt"
)

func main() {
	stih := readTextFromFile(TEXT_FILE_NAME)
	fmt.Print(stih)

	// os.Remove(TEXT_FILE_NAME) // для удаления файла
}
