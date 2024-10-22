// TUGAS 1 : PRAKTIKUM 4
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM 3 BUAH FUNGSI MATEMATIKA (Fogoh, Gohof, Hofog)
package main

import (
	"fmt"
)

// Fungsi f mengkuadratkan nilai input
func f(x_2138 int) int { // Mengkuadratkan nilai x_2138
	return x_2138 * x_2138
}

// Fungsi g mengurangi 2 dari nilai input
func g(x_2138 int) int { // Mengurangi nilai x_2138 dengan 2
	return x_2138 - 2
}

// Fungsi h menambah 1 ke nilai input
func h(x_2138 int) int { // Menambah nilai x_2138 dengan 1
	return x_2138 + 1
}

// Fungsi fogoh menerapkan f ke hasil dari g yang diterapkan pada h dari x_2138
func fogoh(x_2138 int) int { // Menerapkan f(g(h(x_2138)))
	return f(g(h(x_2138)))
}

// Fungsi gohof menerapkan g ke hasil dari h yang diterapkan pada f dari x_2138
func gohof(x_2138 int) int { // Menerapkan g(h(f(x_2138)))
	return g(h(f(x_2138)))
}

// Fungsi hofog menerapkan h ke hasil dari f yang diterapkan pada g dari x_2138
func hofog(x_2138 int) int { // Menerapkan h(f(g(x_2138)))
	return h(f(g(x_2138)))
}

func main() {
	var a_2138, b_2138, c_2138 int

	// Meminta pengguna untuk memasukkan nilai a_2138, b_2138, dan c_2138 dalam satu baris dipisahkan oleh spasi
	fmt.Println("Masukkan nilai a, b, dan c (dipisahkan dengan spasi): ")
	fmt.Scanf("%d %d %d", &a_2138, &b_2138, &c_2138) // Membaca input pengguna

	// Mencetak hasil dari fogoh, gohof, dan hofog berdasarkan input pengguna
	fmt.Printf("fogoh(%d) = %d\n", a_2138, fogoh(a_2138)) // Mencetak hasil dari fogoh(a_2138)
	fmt.Printf("gohof(%d) = %d\n", b_2138, gohof(b_2138)) // Mencetak hasil dari gohof(b_2138)
	fmt.Printf("hofog(%d) = %d\n", c_2138, hofog(c_2138)) // Mencetak hasil dari hofog(c_2138)
}
