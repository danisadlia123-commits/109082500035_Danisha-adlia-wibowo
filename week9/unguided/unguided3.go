package main

import "fmt"

const NMAX int = 100

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil [NMAX]string
	var jumlahHasil int
	var pertandingan int = 1

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	for {
		fmt.Print("Pertandingan ", pertandingan, " : ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlahHasil] = klubA
		} else if skorB > skorA {
			hasil[jumlahHasil] = klubB
		} else {
			hasil[jumlahHasil] = "Draw"
		}

		jumlahHasil++
		pertandingan++
	}

	for i := 0; i < jumlahHasil; i++ {
		fmt.Println("Hasil", i+1, ":", hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}
