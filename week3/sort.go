package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
)

func chunkSlice(slice []int, chunkSize int, rest int) [][]int {
	var chunks [][]int
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		if rest+chunkSize == len(slice) {
			chunkSize += rest
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	return chunks
}

func readSlice() []int {
	slice := make([]int, 0)
	var input string
	for {
		fmt.Println("Insert an integer number:                          X to exit")
		_, err := fmt.Scanf("%s", &input)
		if err == nil {
			if input == exitValue {
				break
			}
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Incorrect input")
				continue
			}
			slice = append(slice, i)
			fmt.Printf("Slice is composed of: %v \n", slice)
			continue
		}
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return slice
}

func sub_sort(wg *sync.WaitGroup, is []int) {
	sort.Ints(is)
	wg.Done()
}

func merge(is1 []int, is2 []int) []int {
	result := make([]int, 0)
	j := 0
	i := 0
	for {
		if i >= len(is1) && j >= len(is2) {
			break
		}
		if j >= len(is2) {
			result = append(result, is1[i])
			i++
			continue

		}
		if i >= len(is1) {
			result = append(result, is2[j])
			j++
			continue

		}
		if is1[i] < is2[j] {
			result = append(result, is1[i])
			i++
			continue

		}
		result = append(result, is2[j])
		j++
	}
	return result
}

const exitValue = "X"
const numberOfChunks = 4

func main() {
	fmt.Println("Assignement week3: sort.go")
	input := readSlice()
	chunk_size := len(input) / numberOfChunks
	rest := len(input) % numberOfChunks
	splitted := chunkSlice(input, chunk_size, rest)
	var wg sync.WaitGroup
	for c := 0; c < numberOfChunks; c++ {
		wg.Add(1)
		go sub_sort(&wg, splitted[c])
	}
	wg.Wait()

	mg1 := merge(splitted[0], splitted[1])
	mg2 := merge(splitted[2], splitted[3])
	sorted := merge( mg1, mg2)
	fmt.Println(sorted)
}
