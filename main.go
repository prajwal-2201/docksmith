package main

import (
	"docksmith/parser"
	"fmt"
)

func main() {

	instructions, err := parser.ParseFile("Docksmithfile")
	if err != nil {
		panic(err)
	}

	for _, inst := range instructions {
		fmt.Println(inst)
	}
}
