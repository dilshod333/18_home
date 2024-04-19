package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		createTxt()
	}()

	go func() {
		defer wg.Done()
		createJson()
	}()

	wg.Wait()

	fmt.Println("All files created successfully.")
	fmt.Println()

	PrintFiles()
}

func createTxt() {
	n, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()

	_, err = n.WriteString("hello it is txt file")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Text file created successfully.")
}

func createJson() {
	s := "Hello, it is json file"
	n, err := os.Create("file2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()

	err = json.NewEncoder(n).Encode(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON file created successfully.")
}

func PrintFiles() {

	txtData, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" file.txt:", string(txtData))

	jsonData, err := os.ReadFile("file2.json")
	if err != nil {
		log.Fatal(err)
	}

	var s string
	err = json.Unmarshal(jsonData, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" file2.json:", s)
}
