# <h1 align="center">LAPORAN PRAKTIKUM MODUL 3 : STRUKTUR KONTROL PERULANGAN DAN PERCABANGAN</h1>
## <p align="center">RICO ADE PRATAMA - 2311102138</p>

# LATIHAN
## 1. SoalLatihan1.go

```go
package main

import (
	"fmt"
)

func main() {
	// Definisikan warna yang benar
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	// Lakukan 5 percobaan
	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Print("Masukkan warna pertama : ")
		fmt.Scanln(&warna1)
		fmt.Print("Masukkan warna kedua : ")
		fmt.Scanln(&warna2)
		fmt.Print("Masukkan warna ketiga : ")
		fmt.Scanln(&warna3)
		fmt.Print("Masukkan warna keempat : ")
		fmt.Scanln(&warna4)

		// Periksa apakah urutan warna sesuai
		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
		}
	}

	// Tampilkan hasil
	fmt.Println("BERHASIL: ", hasil)
}


```
### OUTPUT SCREENSHOT
![OutputSoalLatihan1.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputSoalLatihan1.go.png)

Kode di atas digunakan untuk memeriksa apakah pengguna dapat memasukkan urutan empat warna ("merah", "kuning", "hijau", "ungu") dengan benar dalam lima percobaan, dan menampilkan hasil "BERHASIL: true" jika semua percobaan benar atau "BERHASIL: false" jika ada kesalahan. Lebih jelasnya seperti pada output diatas.

## 2. SoalLatihan2.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += "-" + input
		}
		bungaCount++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}

```
### OUTPUT SCREENSHOT
![OutputSoalLatihan2.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputSoalLatihan2.go.png)

Kode di atas digunakan untuk menerima input nama bunga secara berulang, menggabungkannya menjadi satu string yang dipisahkan oleh tanda hubung (`-`), dan berhenti ketika pengguna mengetik "selesai". Setelah itu, program menampilkan pita yang berisi daftar bunga dan jumlah bunga yang dimasukkan. Lebih jelasnya seperti pada output diatas.

# TUGAS
## 1. Tugas1_Praktikum3.go

```go
// TUGAS 1 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PAK ANDI MENERIMA INPUT DUA BILANGAN REAL POSITIF
package main

import (
	"fmt"
	"math"
)

func main() {
	// Deklarasi variabel untuk menyimpan berat kantong
	var KantongKiri_2138, KantongKanan_2138 float64

	// Loop untuk meminta input berat dari pengguna
	for {
		// Meminta pengguna memasukkan berat belanjaan di kedua kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&KantongKiri_2138, &KantongKanan_2138)

		// Cek apakah salah satu kantong negatif atau total berat melebihi 150 kg
		if KantongKiri_2138 < 0 || KantongKanan_2138 < 0 || KantongKiri_2138+KantongKanan_2138 > 150 {
			// Jika ya, tampilkan pesan dan keluar dari loop
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		Selisih_2138 := math.Abs(KantongKiri_2138 - KantongKanan_2138)

		// Tentukan apakah sepeda motor Pak Andi akan oleng
		Oleng_2138 := Selisih_2138 >= 9

		// Tampilkan hasil apakah sepeda motor oleng
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", Oleng_2138)
	}
}


```
### OUTPUT SCREENSHOT
![OutputTugas1_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas1_Praktikum3.go.png)

Kode di atas digunakan untuk meminta input berat belanjaan di dua kantong (kiri dan kanan) secara berulang, kemudian menghitung selisih berat antara kedua kantong. Jika selisihnya 9 kg atau lebih, sepeda motor dianggap oleng. Program berhenti jika salah satu berat negatif atau total berat melebihi 150 kg, lalu menampilkan pesan "Proses selesai". Lebih jelasnya seperti pada output diatas.

## 2. Tugas2_Praktikum3.go

```go
// TUGAS 2 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENERIMA INPUT BILANGAN SEBUAH K
package main

import (
	"fmt"
	"math"
)

// Menghitung nilai fungsi f(k_2138) berdasarkan rumus yang diberikan
func f(k_2138 float64) float64 {
	// Menghitung pembilang: (4*k_2138 + 2)^2
	Numerator_2138 := math.Pow((4*k_2138 + 2), 2)
	// Menghitung penyebut: (4*k_2138 + 1)*(4*k_2138 + 3)
	Denominator_2138 := (4*k_2138 + 1) * (4*k_2138 + 3)
	// Mengembalikan nilai f(k_2138)
	return Numerator_2138 / Denominator_2138
}

// Menghitung nilai √2 dengan menjumlahkan f(k_2138) dari k_2138=0 hingga k_2138 yang diberikan
func sqrt2(k_2138 int) float64 {
	// Inisialisasi hasil dengan 1
	Hasil_2138 := 1.0
	// Loop untuk menghitung hasil
	for i := 0; i <= k_2138; i++ {
		// Mengalikan hasil dengan f(i)
		Hasil_2138 *= f(float64(i))
	}
	// Mengembalikan nilai hasil akhir
	return Hasil_2138
}

func main() {
	// Deklarasi variabel untuk menyimpan nilai K
	var K_2138 int

	// Loop untuk meminta input 3 kali dari pengguna
	for i := 1; i <= 3; i++ {
		// Meminta pengguna memasukkan nilai K
		fmt.Print("Nilai K = ")
		fmt.Scan(&K_2138)

		// Menghitung nilai akar 2 berdasarkan K yang diberikan
		ApproxSqrt2_1238 := sqrt2(K_2138)
		// Menampilkan hasil dengan format 10 angka di belakang koma
		fmt.Printf("Nilai akar 2 = %.10f\n\n", ApproxSqrt2_1238)
	}
}


```
### OUTPUT SCREENSHOT
![OutputTugas2_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas2_Praktikum3.go.png)

Kode di atas digunakan untuk menerima input bilangan bulat `K` dari pengguna sebanyak tiga kali, kemudian menghitung nilai pendekatan akar 2 menggunakan fungsi `f(k)` yang dijumlahkan dari `k = 0` hingga `k = K`. Hasil pendekatan akar 2 ditampilkan dengan 10 angka di belakang koma. Lebih jelasnya seperti pada output diatas.

## 3. Tugas3_Praktikum3.go

```go
// TUGAS 3 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM BIAYA POS UNTUK MENGHITUNG BIAYA PENGIRIMAN
package main

import (
	"fmt"
)

func main() {
	for i_2138 := 1; i_2138 <= 3; i_2138++ {
		var BeratParsel_2318, kg_2138, gram_2138 int
		var BiayaPerkg_2138, BiayaSisa_2138, TotalBiaya_2138 int

		// Input berat parsel dalam gram
		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&BeratParsel_2318)

		// Menghitung berat dalam kg dan sisa gram
		kg_2138 = BeratParsel_2318 / 1000   // Berat dalam kilogram
		gram_2138 = BeratParsel_2318 % 1000 // Sisa berat dalam gram

		// Menghitung biaya per kg
		BiayaPerkg_2138 = kg_2138 * 10000 // Biaya per kg adalah Rp. 10.000 per kg

		// Menghitung biaya sisa berdasarkan gram
		if kg_2138 >= 10 {
			// Jika berat lebih dari atau sama dengan 10 kg, biaya sisa gram diabaikan (gratis)
			BiayaSisa_2138 = 0
			gram_2138 = 0 // Mengatur gram menjadi 0 untuk mencerminkan gratisnya biaya
		} else {
			// Jika berat kurang dari 10 kg, hitung biaya berdasarkan sisa gram
			if gram_2138 <= 500 {
				BiayaSisa_2138 = gram_2138 * 5 // Jika gram <= 500, biaya Rp. 5 per gram
			} else {
				BiayaSisa_2138 = gram_2138 * 15 // Jika gram > 500, biaya Rp. 15 per gram
			}
		}

		// Menghitung total biaya
		TotalBiaya_2138 = BiayaPerkg_2138 + BiayaSisa_2138 // Total biaya adalah biaya per kg + biaya sisa gram

		// Output hasil perhitungan
		fmt.Printf("Detail berat: %d kg + %d gram\n", kg_2138, gram_2138)              // Menampilkan berat dalam kg dan gram
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", BiayaPerkg_2138, BiayaSisa_2138) // Menampilkan rincian biaya
		fmt.Printf("Total biaya: Rp. %d\n\n", TotalBiaya_2138)                         // Menampilkan total biaya
	}
}

```
### OUTPUT SCREENSHOT
![OutputTugas3_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas3_Praktikum3.go.png)

Kode di atas digunakan untuk menghitung biaya pengiriman berdasarkan berat parsel dalam gram. Berat diubah menjadi kilogram dan gram, kemudian biaya dihitung dengan tarif Rp. 10.000 per kilogram. Jika berat kurang dari 10 kg, biaya tambahan dikenakan untuk sisa gram, dan jika lebih dari 10 kg, sisa gram digratiskan. Program ini mengulangi proses tiga kali dan menampilkan total biaya pengiriman untuk setiap input. Lebih jelasnya seperti pada output diatas.

## 4. Tugas4_Praktikum3.go

```go
// TUGAS 4 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENERIMA INPUT SEBUAH BILANGAN REAL YANG MENYATAKAN NAM, MENGHITUNG NMK DAN MENAMPILKANNYA
package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nam = "A"
	}
	if nam > 72.5 {
		nam = "AB"
	}
	if nam > 65 {
		nam = "B"
	}
	if nam > 57.5 {
		nam = "BC"
	}
	if nam > 50 {
		nam = "C"
	}
	if nam > 40 {
		nam = "D"
	} else if nam <= 40 {
		nam = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}


```
### OUTPUT SCREENSHOT
![OutputTugas4_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas4_Praktikum3.go.png)

### Tugas 4 A 
#### A. Jika NAM diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
Jika NAM diberikan adalah 80.1, maka keluaran dari program adalah "A", karena 80.1 lebih besar dari 80. Namun, program ini tidak sesuai dengan spesifikasi soal. Pada kondisi yang diberikan, setiap pernyataan if dievaluasi secara terpisah, sehingga nilai 80.1 juga akan dievaluasi untuk kondisi-kondisi setelahnya, yaitu untuk AB, B, BC, C, D, dan E, yang menyebabkan hasil bisa menjadi tidak benar.

## Tugas 4 B
### B. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
Kesalahan:
- Program menggunakan kondisi if terpisah tanpa menggunakan else if. Hal ini membuat semua kondisi diperiksa, meskipun kondisi sebelumnya sudah terpenuhi.
- Alur program yang benar seharusnya menggunakan struktur if-else if untuk menghentikan pengecekan setelah kondisi pertama yang sesuai terpenuhi.
Alur yang benar:
- Program harus memeriksa nilai NAM menggunakan struktur if-else if yang berurutan, dimulai dari kondisi tertinggi (NAM ≥ 80) hingga kondisi terendah (NAM < 40).
- Dengan ini, program akan langsung berhenti setelah menemukan kondisi yang sesuai dan tidak memeriksa kondisi lainnya.

## Tugas 4 C_Praktikum3.go
### C. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.

```go
// TUGAS 4 C : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENERIMA INPUT SEBUAH BILANGAN REAL YANG MENYATAKAN NAM, MENGHITUNG NMK DAN MENAMPILKANNYA
package main

import (
	"fmt"
)

func main() {
	var nam_2138 float64 // Variabel untuk menyimpan nilai akhir mata kuliah (NAM)
	var nmk_2138 string  // Variabel untuk menyimpan nilai mutu kuliah (NMK)

	// Input nilai akhir mata kuliah dari pengguna
	fmt.Print("Nilai akhir mata kuliah (NAM): ")
	fmt.Scanln(&nam_2138) // Membaca input nilai NAM

	// Menentukan NMK berdasarkan nilai NAM menggunakan percabangan
	if nam_2138 >= 80 {
		// Jika nilai lebih besar atau sama dengan 80, NMK adalah A
		nmk_2138 = "A"
	} else if nam_2138 >= 72.5 {
		// Jika nilai antara 72.5 hingga 79.9, NMK adalah AB
		nmk_2138 = "AB"
	} else if nam_2138 >= 65 {
		// Jika nilai antara 65 hingga 72.4, NMK adalah B
		nmk_2138 = "B"
	} else if nam_2138 >= 57.5 {
		// Jika nilai antara 57.5 hingga 64.9, NMK adalah BC
		nmk_2138 = "BC"
	} else if nam_2138 >= 50 {
		// Jika nilai antara 50 hingga 57.4, NMK adalah C
		nmk_2138 = "C"
	} else if nam_2138 >= 40 {
		// Jika nilai antara 40 hingga 49.9, NMK adalah D
		nmk_2138 = "D"
	} else {
		// Jika nilai di bawah 40, NMK adalah E
		nmk_2138 = "E"
	}

	// Output NMK yang telah ditentukan berdasarkan nilai NAM
	fmt.Println("Nilai mutu kuliah (NMK):", nmk_2138) // Menampilkan hasil NMK
}

```
### OUTPUT SCREENSHOT
![OutputTugas4C_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas4C_Praktikum3.go.png)

Dalam soal, program perlu diperbaiki untuk menangani tiga input:
- 93.5 harus menghasilkan nilai A.
- 70.6 harus menghasilkan nilai B.
- 49.5 harus menghasilkan nilai D.

Perbaikan Masalah Pada Program Asli:
1. Program yang asli kemungkinan memiliki logika yang salah atau tidak lengkap, sehingga beberapa nilai mungkin tidak menghasilkan NMK yang benar.
2. Kondisi if-else sudah diperbaiki sehingga setiap rentang nilai sekarang diatur sesuai dengan spesifikasi soal.
3. Setiap rentang dijaga agar tidak terjadi tumpang tindih, misalnya rentang nilai B (65 ≤ NAM < 72.5), C (50 ≤ NAM < 57.5), dan D (40 ≤ NAM < 50), sudah diterapkan dengan benar.
4. Dari logika yang diberikan, kita tahu bahwa:
- 93.5: Hasilnya A (karena lebih besar dari 80).
- 70.6: Hasilnya B (karena berada dalam rentang 65 sampai kurang dari 72.5 (65 ≤ NAM < 72.5)).
- 49.5: Hasilnya D (karena berada dalam rentang 40 sampai kurang dari 50 (40 ≤ NAM < 50)).

## 5. Tugas5_Praktikum3.go

```go
// TUGAS 5 : PRAKTIKUM 3
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENCARI DAN MENAMPILKAN SEMUA FAKTOR BILANGAN, KEMUDIAN MENENTUKAN APAKAH B MERUPAKAN BILANGAN PRIMA
package main

import (
	"fmt"
)

func main() {
	var b_2138 int // Variabel untuk menyimpan bilangan yang dimasukkan pengguna

	// Input bilangan bulat dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&b_2138)

	// Mencari dan menampilkan faktor-faktor dari bilangan b_2138
	fmt.Print("Faktor: ")
	HitungFaktor_2138 := 0 // Variabel untuk menghitung jumlah faktor (digunakan untuk menentukan bilangan prima)
	for i_2138 := 1; i_2138 <= b_2138; i_2138++ {
		if b_2138%i_2138 == 0 {
			fmt.Print(i_2138, " ") // Menampilkan faktor
			HitungFaktor_2138++    // Menambah jumlah faktor
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan b_2138 adalah bilangan prima
	if HitungFaktor_2138 == 2 {
		fmt.Println("Prima: true") // Bilangan prima hanya memiliki 2 faktor: 1 dan dirinya sendiri
	} else {
		fmt.Println("Prima: false") // Jika faktor lebih dari 2, berarti bukan bilangan prima
	}
}


```
### OUTPUT SCREENSHOT
![OutputTugas5_Praktikum3.go.png](/PRAKTIKUM%203/Assets%20dan%20Laprak/OutputTugas5_Praktikum3.go.png)

Kode di atas digunakan untuk menerima input bilangan dari pengguna, menampilkan semua faktor dari bilangan tersebut, lalu menentukan apakah bilangan tersebut merupakan bilangan prima. Bilangan prima hanya memiliki dua faktor (1 dan dirinya sendiri), sehingga jika jumlah faktornya 2, bilangan tersebut adalah prima, jika tidak, bukan bilangan prima. Lebih jelasnya seperti pada output diatas.