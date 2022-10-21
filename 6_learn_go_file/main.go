package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "I want to write this line."

	file, err := os.Create("./myfile.text")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The length is %v\n", length)

	defer file.Close()
	readFile("./myfile.text")

}
func readFile(fileName string) {
	bytedata, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is: \n", string(bytedata))
}
