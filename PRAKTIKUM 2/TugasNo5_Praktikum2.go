// TUGAS 5 : PRAKTIKUM 2
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM ASCII
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Variabel untuk menyimpan input
	var int1_2138, int2_2138, int3_2138, int4_2138, int5_2138 int
	var char1_2138, char2_2138, char3_2138 rune
	reader := bufio.NewReader(os.Stdin)

	// Input baris pertama: 5 buah integer
	fmt.Print("<<<<<<< PROGRAM ASCII BY RICO ADE PRATAMA (2311102138) >>>>>>>\n")
	fmt.Println(">> Masukkan 5 buah data integer (pisahkan dengan spasi):")
	fmt.Scanf("%d %d %d %d %d", &int1_2138, &int2_2138, &int3_2138, &int4_2138, &int5_2138)

	// Membersihkan buffer input (newline setelah integer)
	reader.ReadString('\n')

	// Input baris kedua: 3 karakter
	fmt.Println(">> Masukkan 3 buah karakter (berdampingan tanpa spasi):")
	fmt.Scanf("%c%c%c", &char1_2138, &char2_2138, &char3_2138)

	// Konversi integer menjadi karakter ASCII dan cetak (tanpa spasi)
	fmt.Printf("\n<<<<<<< KELUARAN >>>>>>>\n")
	fmt.Printf("== %c%c%c%c%c\n", int1_2138, int2_2138, int3_2138, int4_2138, int5_2138)

	// Geser karakter input sebesar 1 nilai ASCII
	fmt.Printf("== %c%c%c\n", char1_2138+1, char2_2138+1, char3_2138+1)
}
