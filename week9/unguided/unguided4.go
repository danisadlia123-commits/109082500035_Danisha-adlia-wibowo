package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	var selesai bool = false

	*n = 0

	for !selesai && *n < NMAX {
		fmt.Scan(&input)

		for _, karakter := range input {
			if karakter == '.' {
				selesai = true
				break
			}

			t[*n] = karakter
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))

		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func balikArray(t *tabel, n int) {
	var kiri, kanan int
	var temp rune

	kiri = 0
	kanan = n - 1

	for kiri < kanan {
		temp = t[kiri]
		t[kiri] = t[kanan]
		t[kanan] = temp

		kiri++
		kanan--
	}
}

func palindrom(t tabel, n int) bool {
	var salinan tabel

	for i := 0; i < n; i++ {
		salinan[i] = t[i]
	}

	balikArray(&salinan, n)

	for i := 0; i < n; i++ {
		if t[i] != salinan[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab tabel
	var n int
	var hasilPalindrom bool

	fmt.Print("Teks : ")
	isiArray(&tab, &n)

	hasilPalindrom = palindrom(tab, n)

	fmt.Print("Teks terbalik : ")
	balikArray(&tab, n)
	cetakArray(tab, n)

	fmt.Println("Palindrom :", hasilPalindrom)
}