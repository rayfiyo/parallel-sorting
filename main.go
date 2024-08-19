package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/rayfiyo/parallel-sorting/cmd/gen"
	"github.com/rayfiyo/parallel-sorting/cmd/sort"
	"github.com/rayfiyo/parallel-sorting/cmd/ui"
)

const (
	sampleN = 60
)

func main() {
	val := gen.RandomNumber(sampleN)

	fmt.Println("Select sorting algorithm:")
	fmt.Println("1. QuickSort")
	fmt.Println("2. Parallel QuickSort")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var sorter sort.Sorter
	switch input {
	case "1":
		sorter = sort.QuickSort{}
	case "2":
		sorter = sort.ParallelQuickSort{}
	default:
		log.Fatal("Invalid selection")
	}

	go func() {
		sortedVal := sorter.Sort(val)
		for i, v := range sortedVal {
			val[i] = v
			time.Sleep(100 * time.Millisecond)
		}
	}()

	if err := ui.Draw(val, "Sorting: "+sorter.Name()); err != nil {
		log.Panic(err)
	}
}
