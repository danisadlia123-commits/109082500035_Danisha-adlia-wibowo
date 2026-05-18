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

	ketua := -1
	wakil := -1

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			if ketua == -1 || hitungSuara[i] > hitungSuara[ketua] {
				wakil = ketua
				ketua = i
			} else if wakil == -1 || hitungSuara[i] > hitungSuara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	if ketua != -1 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != -1 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}