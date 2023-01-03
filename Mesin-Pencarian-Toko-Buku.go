package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [5]string{
		"Filosofi Teras",
		"Atomic Habit",
		"Blue Ocean",
		"Art of Leadership",
		"Laskar Pelangi",
	}

	var ada bool

	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("Silakan masukkan judul buku yang anda cari")
		return
	}

	hurufKecil := strings.ToLower(arg[0])

	fmt.Println("Hasil Pencarian: ")

	for _, v := range books {
		if strings.Contains(strings.ToLower(v), hurufKecil) {
			fmt.Println("+", v)
			ada = true
		}

	}

	if !ada {
		fmt.Printf("Kami tidak mempunyai buku yang anda cari: %q\n", hurufKecil)
		return
	}
}
