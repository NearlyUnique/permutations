package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type (
	//File definition
	File struct {
		name  string
		lines []string
	}
)

var (
	source []File
)

func main() {
	flag.Parse()
	source = make([]File, len(flag.Args()))
	total := 1
	for i, f := range flag.Args() {
		source[i] = File{name: f}
		file, err := os.Open(f)
		if err != nil {
			log.Printf("Unable to read: %s\n\t%v\n", f, err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			source[i].lines = append(source[i].lines, scanner.Text())
		}
		total *= len(source[i].lines)
	}

	generate()
}

var index []int

func generate() {
	index = make([]int, len(source))
	for {
		printCurrentLine()

		carry := true
		for x := len(index) - 1; x >= 0; x-- {
			if carry {
				index[x]++
			}
			if index[x] == len(source[x].lines) {
				if x == 0 {
					return
				}
				index[x] = 0
				carry = true
			} else {
				carry = false
			}
		}
	}
}

func printCurrentLine() {
	var out []string
	for i, src := range source {
		out = append(out, src.lines[index[i]])
	}
	fmt.Println(strings.Join(out, "|"))
}
