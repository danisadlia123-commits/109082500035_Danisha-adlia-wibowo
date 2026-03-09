package main

import "fmt"

func hitung(berat int) {
	kg := berat / 1000
	sisa := berat % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg > 10 {
		biayaSisa = 0
	} else if sisa > 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 5
	}

	total := biayaKg + biayaSisa

	fmt.Println("Masukkan total berat (gram):", berat)
	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %2d kg + %3d  gram\n", kg, sisa)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", total)
	fmt.Println()
}

func main() {

	hitung(8500)
	hitung(9250)
	hitung(11750)

}