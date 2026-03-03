package main
import "fmt"

func main() {
	var berat, tinggi, bmi float64

	fmt.Print("masukan berat badan (KG): ")
	fmt.Scanln(&berat)
	fmt.Print("masukan tinggi badan (m): ")
	fmt.Scanln(&tinggi)

	bmi = berat / (tinggi * tinggi)
	fmt.Printf("Nilai BMI anda adalah: %.2f\n", bmi)
}