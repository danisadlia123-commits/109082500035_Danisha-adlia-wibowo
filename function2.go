package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int


	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&c)

	
	result1 := f(g(h(a)))

	
	result2 := g(h(f(b)))


	result3 := h(f(g(c)))


	fmt.Println("F(G(H(", a, "))) :", result1)
	fmt.Println("G(H(F(", b, "))) :", result2)
	fmt.Println("H(F(G(", c, "))) :", result3)
}