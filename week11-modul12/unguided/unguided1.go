package main

import "fmt"

func main() {
	var suara, totalMasuk, suaraSah int
	var hitungSuara [21]int

	totalMasuk = 0
	suaraSah = 0

	fmt.Scan(&suara)
	for suara != 0 {
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			hitungSuara[suara]++
			suaraSah++
		}
		fmt.Scan(&suara)
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuara[i])
		}
	}
}