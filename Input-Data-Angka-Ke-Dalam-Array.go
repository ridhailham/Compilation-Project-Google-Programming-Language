package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		baris  [5]int
		jumlah float64
		total  float64
	)

	arg := os.Args[1:]

	if l := len(arg); l < 1 || l > 5 {
		fmt.Println("masukkan angka")
		return
	}

	for i, v := range arg {
		n, err := strconv.Atoi(v)

		if err != nil {
			baris[i] = 0
			jumlah += float64(0)
			total++
		} else {
			baris[i] = n
			jumlah += float64(n)
			total++
		}

	}

	fmt.Println("hasil array menjadi", baris)
	fmt.Println("jumlahnya menjadi", jumlah)
	fmt.Println("rata - rata menjadi", jumlah/total)

}
