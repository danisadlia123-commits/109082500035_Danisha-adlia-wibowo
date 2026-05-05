package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, n *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}

	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nim string) int {
	terbesar := -1

	for i := 0; i < n; i++ {
		if T[i].NIM == nim {
			if T[i].nilai > terbesar {
				terbesar = T[i].nilai
			}
		}
	}

	return terbesar
}

func tampilkanHasil(nilaiPertama int, nilaiTerbesar int, nim string) {
	if nilaiPertama == -1 {
		fmt.Println("Data dengan NIM", nim, "tidak ditemukan")
	} else {
		fmt.Println("Nilai pertama dari NIM", nim, "adalah", nilaiPertama)
		fmt.Println("Nilai terbesar dari NIM", nim, "adalah", nilaiTerbesar)
	}
}

func main() {
	var data arrayMahasiswa
	var n int
	var nimDicari string
	var nilaiPertama, nilaiTerbesar int

	inputData(&data, &n)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimDicari)

	nilaiPertama = cariNilaiPertama(data, n, nimDicari)
	nilaiTerbesar = cariNilaiTerbesar(data, n, nimDicari)

	tampilkanHasil(nilaiPertama, nilaiTerbesar, nimDicari)
}