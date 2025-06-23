package main

import (
	"fmt"
	"log"
	"os"
)

func writeFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	filename := "write.txt"
	content := "hello sumon ovverride"
	err := writeFile(filename, content)

	if err != nil {
		log.Fatalf("error : %v", err)
	}

	fmt.Printf("File '%s' written successfully.\n", filename)
}
