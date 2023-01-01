package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "lazy cat jump again and again and again"

func main() {
	word := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {

		for i, w := range word {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i, q)
				continue queries
			}

		}

	}

}
