package main

import (
	"fmt"
	"test_api/model"
)

func main() {
	fmt.Println("Hello world.")
	result := 1 + 1
	fmt.Printf("result: %d\n", result)
	p_slice := [...]*model.Person{model.New("hiroki", "kayanuma", 25), model.New("Shiho", "Ito", 24)}
	for _, p := range p_slice {
		p.Greetings()
	}
	fmt.Printf("type: %T, value: %v\n", p_slice, p_slice)
}
