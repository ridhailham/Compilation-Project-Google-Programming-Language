package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	arg := os.Args[1:]

// 	var (
// 		hasil float64
// 		angka [5]float64
// 		total float64
// 	)

// 	if l := len(arg); l == 0 || l > 5 {
// 		fmt.Println("Wajib masukkan angka minimal 1 dan maksimal 5")
// 		return
// 	}

// 	for i, v := range arg {
// 		n, err := strconv.ParseFloat(v, 64)
// 		if err != nil {
// 			continue
// 		}

// 		angka[i] = n

// 		hasil += n
// 		total++
// 	}

// 	fmt.Printf("Angka Anda: %.0f\n", angka)
// 	fmt.Printf("Rata - rata: %.1f\n", hasil/total)
// }
