package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {


	const makan = 35000 * 2   
	const penginapan = 250000 
	const uangSaku = 300000

	return jumlahMhs * (makan + penginapan + uangSaku)
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {


	var hariDitanggung int
	var biayaHarian int

	hariDitanggung = tanggunganHari(lamaPerjalanan, tujuan)

	biayaHarian = biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaHarian)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaHarian) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}