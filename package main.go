package main

import (
	"fmt"
)

func main() {
	// Mendeklarasikan variabel nilai
	nilai := 75

	// Penggunaan if-else untuk mengevaluasi nilai
	if nilai >= 80 {
		// Jika nilai lebih besar atau sama dengan 80
		fmt.Println("Selamat! Anda lulus dengan nilai sangat baik.")
	} else if nilai >= 60 {
		// Jika nilai antara 60 dan 79
		fmt.Println("Anda lulus dengan nilai cukup baik.")
	} else {
		// Jika nilai kurang dari 60
		fmt.Println("Maaf, Anda belum lulus. Perlu belajar lebih giat lagi.")
	}
}
