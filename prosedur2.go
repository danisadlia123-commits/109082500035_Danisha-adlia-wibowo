package main

import "fmt"

func main() {
	var pilih int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	switch pilih {

	case 1:
		var sisi int
		fmt.Print("\nMasukkan sisi : ")
		fmt.Scan(&sisi)

		fmt.Println("Luas persegi :", sisi*sisi)
		fmt.Println("Keliling persegi :", 4*sisi)

	case 2:
		var p, l int
		fmt.Print("\nMasukkan panjang : ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&l)

		fmt.Println("Luas persegi panjang :", p*l)
		fmt.Println("Keliling persegi panjang :", 2*(p+l))

	case 3:
		var r float64
		fmt.Print("\nMasukkan jari-jari : ")
		fmt.Scan(&r)

		fmt.Println("Luas lingkaran :", 3.14*r*r)
		fmt.Println("Keliling lingkaran :", 2*3.14*r)

	}
}