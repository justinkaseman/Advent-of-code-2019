package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func calcFuelRecursive(current int) int {
	if current < 9 {
		return 0
	}
	return (int(math.Floor((float64(current / 3)))) - 2) + calcFuelRecursive(int(math.Floor((float64(current/3))))-2)
}

func scanForEach() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total = total + calcFuelRecursive(current)
	}
	return total
}

func main() {
	fmt.Println(scanForEach())
}
