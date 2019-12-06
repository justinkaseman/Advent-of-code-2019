package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	strings :=
		strings.Split(string(b), "-")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func checkTwoAdjacent(current int) bool {
	strg := strconv.Itoa(current)
	adjMap := make(map[int64]int64)
	for i := 1; i < len(strg); i++ {
		first, err := strconv.ParseInt(string(strg[i]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.ParseInt(string(strg[i-1]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if first == second {
			adjCount, ok := adjMap[first]
			if ok {
				adjMap[first] = adjCount + 1
			} else {
				adjMap[first] = 1
			}
		}
	}

	for _, value := range adjMap {
		if value == 1 {
			return true
		}
	}

	return false
}

func checkAscending(current int) bool {
	strg := strconv.Itoa(current)
	for i := 1; i < len(strg); i++ {
		first, err := strconv.ParseInt(string(strg[i]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.ParseInt(string(strg[i-1]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if first < second {
			return false
		}
	}
	return true
}

func main() {
	rng := readFile()
	count := 0
	current := rng[0]
	for current <= rng[1] {
		if checkTwoAdjacent(current) && checkAscending(current) {
			fmt.Println("counted ", current)
			count = count + 1
		}
		current = current + 1
	}
	fmt.Println("Final count: ", count)
}
