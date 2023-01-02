package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 && len(arg) != 2 {
		fmt.Println("[Masukkan Nama]")
		return
	}

	name, mood := arg[0], arg[1]

	moods := [3][3]string{
		{"senang", "baik", "optimis"},
		{"sedih", "bad mood", "ambisius"},
	}

	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(len(moods))

	var n int

	if mood == "positif" {
		n = 0
	} else if mood == "negatif" {
		n = 1
	}

	fmt.Printf("%s merasa %s\n", name, moods[n][m])
}
