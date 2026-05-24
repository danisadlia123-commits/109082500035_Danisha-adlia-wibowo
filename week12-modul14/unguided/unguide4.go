package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Println("=== BUKU TERFAVORIT ===")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := n
	if batas > 5 {
		batas = 5
	}
	fmt.Println("=== 5 BUKU RATING TERTINGGI ===")
	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s | %s | %s | %d | Rating: %d\n",
			i+1,
			pustaka[i].judul,
			pustaka[i].penulis,
			pustaka[i].penerbit,
			pustaka[i].tahun,
			pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = tengah
			break
		} else if pustaka[tengah].rating > r {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("=== HASIL PENCARIAN RATING", r, "===")
	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	} else {
		fmt.Println("Judul     :", pustaka[ketemu].judul)
		fmt.Println("Penulis   :", pustaka[ketemu].penulis)
		fmt.Println("Penerbit  :", pustaka[ketemu].penerbit)
		fmt.Println("Tahun     :", pustaka[ketemu].tahun)
		fmt.Println("Eksemplar :", pustaka[ketemu].eksemplar)
		fmt.Println("Rating    :", pustaka[ketemu].rating)
	}
}

func main() {
	var r int

	DaftarkanBuku(&Pustaka, &nPustaka)

	fmt.Scan(&r)

	CetakTerfavorit(Pustaka, nPustaka)
	fmt.Println()

	UrutBuku(&Pustaka, nPustaka)

	Cetak5Terbaru(Pustaka, nPustaka)
	fmt.Println()

	CariBuku(Pustaka, nPustaka, r)
}