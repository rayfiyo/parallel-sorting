package main

import (
	"log"

	"github.com/rayfiyo/parallel-sorting/cmd/gen"
	"github.com/rayfiyo/parallel-sorting/cmd/ui"
)

const (
	sampleN = 60
)

func main() {
	val := gen.RandomNumber(sampleN)

	if err := ui.Draw(val, "Sorting Visualization"); err != nil {
		log.Panic(err)
	}

	if err := ui.PlaySort(nil, nil, val); err != nil {
		log.Panic(err)
	}
}
