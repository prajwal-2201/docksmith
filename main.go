package main

import (
	"docksmith/build"
	"docksmith/cache"
	"docksmith/parser"
	"fmt"
)

func main() {

	instructions, err := parser.ParseFile("Docksmithfile")
	if err != nil {
		panic(err)
	}

	builder := build.Builder{
		Cache: cache.NewCache(),
	}

	state, err := builder.Build(instructions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Layers:", state.Layers)
	fmt.Println("Workdir:", state.WorkDir)
	fmt.Println("Env:", state.Env)
	fmt.Println("Cmd:", state.Cmd)

	fmt.Println("Build completed")
}
