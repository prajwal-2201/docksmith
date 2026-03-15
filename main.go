package main

import (
	"docksmith/build"
	"docksmith/parser"
	"fmt"
)

func main() {

	instructions, err := parser.ParseFile("Docksmithfile")
	if err != nil {
		panic(err)
	}

	builder := build.Builder{}

	err = builder.Build(instructions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Build completed")
}
