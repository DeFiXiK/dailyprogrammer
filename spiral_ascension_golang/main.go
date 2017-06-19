package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readValues("input.txt")
	for _, size := range input {
		spiral := make([][]int, size)
		for i := range spiral {
			spiral[i] = make([]int, size)
		}
		for i := range spiral {
			for j := range spiral[i] {
				spiral[i][j] = calcValue(size, j, i)
			}
		}
		printSpiral(spiral)
	}
}

func printSpiral(spiral [][]int) {
	for i := range spiral {
		for j := range spiral[i] {
			fmt.Printf("%3d ", spiral[i][j])
		}
		fmt.Println("")
	}
	fmt.Println()
}

func calcValue(size, indexX, indexY int) (value int) {
	if indexY == 0 {
		return indexX + 1
	}
	if indexX == size-1 {
		return size + indexY
	}
	if indexY == size-1 {
		return (size + size - 1 + size - 1) - indexX
	}
	if indexX == 0 {
		return (size + size - 1 + size - 1 + size - 1) - indexY
	}
	return calcValue(size-2, indexX-1, indexY-1) + (size + size - 1 + size - 1 + size - 2)
}

func readValues(fname string) []int {
	f, err := os.Open(fname)
	if err != nil {
		fmt.Printf("File was not opened")
		return nil
	}
	var out = []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			inp, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Input value not converted")
				continue
			}
			out = append(out, inp)
		}
	}
	return out
}
