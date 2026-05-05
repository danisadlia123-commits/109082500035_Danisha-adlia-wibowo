package main

import "fmt"

const nProv = 10

type namaProv [nProv]string
type popProv [nProv]int
type tumbuhProv [nProv]float64

func inputData(prov *namaProv, pop *popProv, tumbuh *tumbuhProv, n int) {
	for i := 0; i < n; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&(*prov)[i], &(*pop)[i], &(*tumbuh)[i])
	}
}

func provinsiTercepat(tumbuh tumbuhProv, n int) int {
	indeks := 0

	for i := 1; i < n; i++ {
		if tumbuh[i] > tumbuh[indeks] {
			indeks = i
		}
	}

	return indeks
}

func indeksProvinsi(prov namaProv, n int, nama string) int {
	for i := 0; i < n; i++ {
		if prov[i] == nama {
			return i
		}
	}

	return -1
}

func prediksi(prov namaProv, pop popProv, tumbuh tumbuhProv, n int) {
	fmt.Println("----- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% -----")

	for i := 0; i < n; i++ {
		if tumbuh[i] >= 0.02 {
			hasil := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s %.0f\n", prov[i], hasil)
		}
	}
}

func main() {
	var prov namaProv
	var pop popProv
	var tumbuh tumbuhProv
	var n int
	var namaDicari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	inputData(&prov, &pop, &tumbuh, n)

	fmt.Print("Masukkan nama provinsi yang dicari : ")
	fmt.Scan(&namaDicari)

	indeksCepat := provinsiTercepat(tumbuh, n)
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[indeksCepat])

	indeksCari := indeksProvinsi(prov, n, namaDicari)

	if indeksCari == -1 {
		fmt.Println("Data provinsi yang dicari tidak ditemukan")
	} else {
		fmt.Println("Data provinsi yang dicari :", prov[indeksCari])
	}

	prediksi(prov, pop, tumbuh, n)
}