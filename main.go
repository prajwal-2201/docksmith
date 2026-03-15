package main

import (
	"docksmith/build"
	"docksmith/utils"
	"fmt"
)

func main() {

	files, _ := build.CollectFiles("*.go")

	layer, err := build.CreateLayer(files)
	if err != nil {
		panic(err)
	}

	digest := utils.ComputeDigest(layer)

	fmt.Println("layer size:", len(layer))
	fmt.Println("digest:", digest)
}
