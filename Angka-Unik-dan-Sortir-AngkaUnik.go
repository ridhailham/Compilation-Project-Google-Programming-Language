package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5

	var unik []int

loop:
	for ada := 0; ada < max; {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range unik {
			if u == n {
				continue loop
			}
		}
		unik = append(unik, n)
		ada++
	}
	fmt.Println()

	fmt.Println("angka unik:", unik)

	sort.Ints(unik)

	fmt.Println("sortir angka unik:", unik)
}
