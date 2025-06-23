package main

import (
	"fmt"
	"log"
	"os"
)

func readFileAndReturnString(filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return "error happen", err
	}

	content := string(data)
	return content, nil
}

func main() {

	filename := "test.txt"

	content, err := readFileAndReturnString(filename)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println("Contents of the file that read ")

	fmt.Println(content)

}
