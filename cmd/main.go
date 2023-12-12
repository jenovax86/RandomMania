package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Fact struct {
	Facts []string
}

func randomFactGenerator(filePath string) string {
	var payload Fact
	var randomNumber int
	file, err := os.Open(filePath)
	defer file.Close()

	jsonFile, err := io.ReadAll(file)
	if err != nil {
		panic("Error reading file")
	}
	if err := json.Unmarshal(jsonFile, &payload); err != nil {
		panic(err)
	}
	randomNumber = rand.Intn(len(payload.Facts))
	return payload.Facts[randomNumber]
}

func main() {
	path := "../facts.json"
	facts := randomFactGenerator(path)
	fmt.Println(facts)
}
