package main

import (
	"bytes"
	"fmt"
	"log"
	"test_api/model"
)

func main() {
	var buf bytes.Buffer
	var logger = log.New(&buf, "logger: ", log.Lshortfile)
	logger.Println("Hello world.")
	result := 1 + 1
	logger.Printf("result: %d\n", result)
	p_slice := []*model.Person{model.New("hiroki", "kayanuma", 25), model.New("Shiho", "Ito", 24)}
	p_slice = append(p_slice, model.New("Chieko", "Kayanuma", 60))
	for _, p := range p_slice {
		p.Greetings(*logger)
	}
	logger.Printf("type: %T, value: %v\n", p_slice, p_slice)
	fmt.Println(&buf)
}
