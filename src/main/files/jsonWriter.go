package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonWriter()
	readingJsonFromFile()
}

type Message struct {
	Name string
	Body string
	Time int64
}

var fileName = "tmpFile.json"

func jsonWriter() {
	os.Create(fileName)
	//m := Message{"Alice", "Hello", 1294706395881547000}
	arr := []string{"dsd", "dsd"}
	b, _ := json.Marshal(arr)
	os.WriteFile(fileName, b, 0666)
}

func readingJsonFromFile() {
	//var m Message
	var arr []string
	b, _ := os.ReadFile(fileName)
	_ = json.Unmarshal(b, &arr)
	fmt.Printf("Type is %T and value is %v \n", arr, arr)
}
