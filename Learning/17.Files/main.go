package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling in GoLang")
	content := "this code need to go into the file -LearnCodeOnline"

	file, err := os.Create("./mylcogofile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the data \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
