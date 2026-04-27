package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func diDalamLingkaran(t Titik, l Lingkaran) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y

	jarakKuadrat := dx*dx + dy*dy
	radiusKuadrat := l.r * l.r

	return jarakKuadrat <= radiusKuadrat
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	dalamL1 := diDalamLingkaran(t, l1)
	dalamL2 := diDalamLingkaran(t, l2)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}