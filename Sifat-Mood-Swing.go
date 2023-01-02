package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("[Masukkan Nama]")
		return
	}

	name := arg[0]

	moods := [...]string{
		"senang", "baik", "optimis", "sedih", "bad mood", "ambisius",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s merasa %s", name, moods[n])
}
