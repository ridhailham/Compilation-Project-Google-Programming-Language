package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)

	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println(
			`Selamat datang di game tebak angka keberuntungan

		Selamat bermain dengan bahagia :)`)
		return
	}

	tebakan, err := strconv.Atoi(arg[0])

	if err != nil {
		fmt.Println("Silakan tulis angka bukan kalimat atau kata")
		return
	}

	if tebakan < 0 {
		fmt.Println("Silakan masukkan angka positif, bukan negatif yaa :v")
		return
	}

	for n := 0; n != tebakan; {
		n = rand.Intn(tebakan + 1)

		if n == tebakan {
			fmt.Printf("Anda mendapatkan angka tebakan: %d\n", n)
			fmt.Println("Anda berhasil yaa :)")
			break
		} else {
			fmt.Printf("%d bukan angka tebakan \n", n)
		}
	}
	fmt.Println("")
}
