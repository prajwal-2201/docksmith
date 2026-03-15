package main

import (
	"docksmith/build"
	"fmt"
)

func main() {

	files, err := build.CollectFiles("*.go")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f)
	}
}
