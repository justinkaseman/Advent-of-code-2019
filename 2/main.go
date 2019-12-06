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
		strings.Split(string(b), ",")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func runTape(noun, verb int) []int {
	output := readFile()
	ip := 0

	output[1] = noun
	output[2] = verb

	for output[ip] != 99 {
		if output[ip] == 1 {
			output = opCode1(output, ip)
		} else if output[ip] == 2 {
			output = opCode2(output, ip)
		}
		ip = ip + 4
	}
	return output
}

func opCode1(tape []int, cursor int) []int {
	arg1Location := tape[cursor+1]
	arg2Location := tape[cursor+2]
	outputLocation := tape[cursor+3]
	// fmt.Println(tape[arg1Location], "+", tape[arg2Location], "to", tape[outputLocation])
	tape[outputLocation] = tape[arg1Location] + tape[arg2Location]
	return tape
}

func opCode2(tape []int, cursor int) []int {
	arg1Location := tape[cursor+1]
	arg2Location := tape[cursor+2]
	outputLocation := tape[cursor+3]
	// fmt.Println(tape[arg1Location], "*", tape[arg2Location], "to", tape[outputLocation])
	tape[outputLocation] = tape[arg1Location] * tape[arg2Location]
	return tape
}

func main() {
	input := 19690720

	zeroRun := runTape(0, 0)[0]
	remainder := input - zeroRun
	noun := runTape(1, 0)[0] - zeroRun
	verb := runTape(0, 1)[0] - zeroRun
	if input/noun < input/verb {
		outputNoun := remainder / noun
		outputVerb := (input - runTape(outputNoun, 0)[0]) / verb
		fmt.Println("noun:", outputNoun, " verb:", outputVerb, " state:", runTape(outputNoun, outputVerb)[0])
	} else {
		outputVerb := remainder / verb
		outputNoun := (runTape(0, outputVerb)[0]) / noun
		fmt.Println("noun:", outputNoun, " verb:", outputVerb, " state:", runTape(outputNoun, outputVerb)[0])
	}
}
