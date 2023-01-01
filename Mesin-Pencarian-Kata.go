package main

import (
	"fmt"
	"os"
	"strings"
)

const tampunganKata = "" + "lazy cat jump again and again and again"

func main() {
	kata := strings.Fields(tampunganKata)
	cariKata := os.Args[1:]

looping:
	for _, q := range cariKata {

		for i, w := range kata {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i, q)
				continue looping
			}

		}

	}

}
