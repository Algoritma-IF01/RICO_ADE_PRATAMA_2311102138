# <h1 align="center">LAPORAN PRAKTIKUM MODUL 5 : STRUCK DAN ARRAY</h1>
## <p align="center">RICO ADE PRATAMA - 2311102138</p>

# LATIHAN
## 1. Alias(Type).go

```go
package main

import "fmt"

type bilangan int
type pecahan float64

func main() {
	var a, b bilangan
	var hasil pecahan
	a = 9
	b = 5
	hasil = pecahan(a) / pecahan(b)
	fmt.Println(hasil)

}

```
### OUTPUT SCREENSHOT
![Output_Alias(Type).go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_Alias(Type).go.png)

Kode di atas digunakan untuk mendeklarasikan tipe data `bilangan` sebagai `int` dan `pecahan` sebagai `float64`. Program kemudian menghitung hasil pembagian dari dua variabel `a` dan `b` yang dideklarasikan sebagai tipe `bilangan`, yaitu `9` dan `5`. Variabel-variabel ini dikonversi ke tipe `pecahan` untuk pembagian, sehingga hasilnya berupa `float64`. Program mencetak hasil pembagian tersebut, yaitu `1.8`. Lebih jelasnya seperti pada output diatas.

## 2. Struct.go

```go
package main

import "fmt"

type waktu struct {
	jam, menit, detik int
}

func main() {
	var wParkir, wPulang, durasi waktu
	var dParkir, dPulang, lParkir int
	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)
	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
	lParkir = dPulang - dParkir
	durasi.jam = lParkir % 3600 / 60
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 / 60
	fmt.Printf("Lama Parkir: %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}

```
### OUTPUT SCREENSHOT
![Output_Struck.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_Struct.go.png)

Kode di atas digunakan untuk Program ini menghitung durasi parkir dengan tipe data `waktu`, yang berisi `jam`, `menit`, dan `detik`. Waktu parkir dan waktu pulang dikonversi menjadi total detik untuk menghitung durasi parkir (`lParkir`). Durasi tersebut kemudian dikonversi kembali menjadi `jam`, `menit`, dan `detik` sebelum ditampilkan. Lebih jelasnya seperti pada output diatas.

## 3. Array.go

```go
package main

import "fmt"

// Definisi tipe CircType
type CircType struct {
	Radius float64
}

// Definisi tipe NewType
type NewType struct {
	Name string
}

func main() {
	var (
		// Array arr mempunyai 73 elemen, masing-masing bertipe CircType
		arr [73]CircType

		// Array buf dengan 5 elemen, dengan nilai awal 7,3,5,2,11
		buf = [5]byte{7, 3, 5, 2, 11}

		// Array mhs berisi 2000 elemen NewType
		mhs [2000]NewType

		// Array dua dimensi rec berisi nilai float64
		rec [20][40]float64
	)

	// Mengisi beberapa elemen contoh
	arr[0] = CircType{Radius: 5.5}
	mhs[0] = NewType{Name: "Alice"}
	rec[0][0] = 3.14

	// Contoh penggunaan variabel
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("buf:", buf)
	fmt.Println("mhs[0]:", mhs[0])
	fmt.Println("rec[0][0]:", rec[0][0])
}

```
### OUTPUT SCREENSHOT
![Output_Array.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_Array.go.png)

Kode di atas digunakan untuk mendefinisikan tipe data `CircType` dan `NewType`, kemudian mendeklarasikan beberapa array: `arr` dengan 73 elemen `CircType`, `buf` dengan 5 elemen byte yang diinisialisasi, `mhs` dengan 2000 elemen `NewType`, dan `rec` dua dimensi [20][40] dengan nilai `float64`. Program ini mengisi beberapa elemen dan mencetaknya. Lebih jelasnya seperti pada output diatas.

## 4. Map.go

```go
package main

import "fmt"

func main() {
	// Membuat map dengan tipe string sebagai kunci dan integer sebagai nilai
	scores := map[string]int{
		"John": 90,
		"Anne": 85,
	}

	// Mengambil nilai dari kunci
	johnScore := scores["John"]
	fmt.Println("Nilai John:", johnScore)

	// Mengganti nilai dari kunci yang ada
	scores["John"] = 95
	fmt.Println("Nilai John setelah diganti:", scores["John"])

	// Menambah kunci baru dengan nilai
	scores["David"] = 88
	fmt.Println("Nilai David ditambahkan:", scores["David"])

	// Menghapus kunci dari map
	delete(scores, "Anne")
	fmt.Println("Map setelah kunci 'Anne' dihapus:", scores)

	// Mengecek apakah kunci ada dalam map
	if score, ada := scores["David"]; ada {
		fmt.Println("Nilai David ditemukan:", score)
	} else {
		fmt.Println("Nilai David tidak ditemukan")
	}
}


```
### OUTPUT SCREENSHOT
![Output_Map.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Map.go.png)

Kode di atas digunakan untuk Program ini menggunakan map untuk menyimpan skor siswa. Program memuat skor "John" dan "Anne", kemudian menampilkan, mengubah, menambah, dan menghapus elemen dalam map. Terakhir, program mengecek apakah kunci tertentu ada di dalam map. Lebih jelasnya seperti pada output diatas.

## 5. Latihan1.go

```go
package main

import (
	"fmt"
	"math"
)

// Definisi tipe bentukan untuk titik
type Titik struct {
	x int
	y int
}

// Definisi tipe bentukan untuk lingkaran
type Lingkaran struct {
	center Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p Titik, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.center) < float64(c.radius)
}

func main() {
	var (
		// Mengambil input untuk lingkaran 1
		lingkaran1 Lingkaran
		// Mengambil input untuk lingkaran 2
		lingkaran2 Lingkaran
		// Mengambil input untuk titik sembarang
		point Titik
	)

	// Input untuk lingkaran 1 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)

	// Input untuk lingkaran 2 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)

	// Input untuk titik sembarang (x, y)
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&point.x, &point.y)

	// Mengecek posisi titik terhadap kedua lingkaran
	inLingkaran1 := didalam(lingkaran1, point)
	inLingkaran2 := didalam(lingkaran2, point)

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}


```
### OUTPUT SCREENSHOT
![Output_Latihan1.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Latihan1.go.png)

Kode di atas digunakan untuk mendefinisikan tipe `Titik` dan `Lingkaran`, lalu menerima input untuk dua lingkaran dan satu titik sembarang. Program mengecek apakah titik tersebut berada di dalam salah satu atau kedua lingkaran dengan menggunakan fungsi `jarak` untuk menghitung jarak antara dua titik dan fungsi `didalam` untuk memeriksa apakah jarak tersebut lebih kecil dari radius lingkaran. Berdasarkan hasilnya, program mencetak apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar. Lebih jelasnya seperti pada output diatas.

# TUGAS
## 1. TugasNo2.go

```go
// TUGAS 2 : PRAKTIKUM 5
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM OPERASI ARRAY
package main

import (
	"fmt"  // Mengimpor paket fmt untuk I/O
	"math" // Mengimpor paket math untuk operasi matematika
)

// Fungsi untuk menampilkan seluruh elemen array
func tampilkanSEluruhArray(array_2138 []int) { // array_2138 - array dari tipe int yang ingin ditampilkan
	fmt.Println("Seluruh isi array:", array_2138)
}

// Fungsi untuk menampilkan elemen array dengan indeks ganjil
func tampilkanIndeksGanjil(array_2138 []int) { // array_2138 - array dari tipe int
	fmt.Print("Elemen dengan indeks ganjil: ")
	for _, value_2138 := range array_2138 { // Loop untuk setiap elemen dalam array
		if value_2138%2 != 0 { // Memeriksa apakah elemen adalah ganjil
			fmt.Print(value_2138, " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array dengan indeks genap
func tampilkanIndeksGenap(array_2138 []int) { // array_2138 - array dari tipe int
	fmt.Print("Elemen dengan indeks genap: ")
	for _, value_2138 := range array_2138 { // Loop untuk setiap elemen dalam array
		if value_2138%2 == 0 && value_2138 != 0 { // Memeriksa apakah elemen adalah genap dan bukan nol
			fmt.Print(value_2138, " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array pada indeks kelipatan x_2138
func tampilkanKelipatanX(array_2138 []int, x_2138 int) { // array_2138 - array dari tipe int, x_2138 - nilai kelipatan indeks
	fmt.Printf("Elemen pada indeks kelipatan %d: ", x_2138)
	for i_2138 := x_2138 - 1; i_2138 < len(array_2138); i_2138 += x_2138 { // Loop untuk setiap elemen pada indeks kelipatan x_2138
		if array_2138[i_2138] != 0 { // Memeriksa apakah elemen bukan nol
			fmt.Print(array_2138[i_2138], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen array pada indeks tertentu
func hapusElemen(array_2138 []int, idx_2138 int) []int { // array_2138 - array dari tipe int, idx_2138 - indeks yang akan dihapus
	return append(array_2138[:idx_2138], array_2138[idx_2138+1:]...) // Mengembalikan array baru dengan elemen yang telah dihapus
}

// Fungsi untuk menghitung rata-rata elemen array
func tampilkanRataRata(array_2138 []int) float64 { // array_2138 - array dari tipe int
	sum_2138 := 0                           // Variabel untuk menyimpan jumlah elemen
	for _, value_2138 := range array_2138 { // Loop untuk setiap elemen dalam array
		sum_2138 += value_2138 // Menambahkan elemen ke jumlah total
	}
	return float64(sum_2138) / float64(len(array_2138)) // Mengembalikan rata-rata dalam tipe float64
}

// Fungsi untuk menghitung deviasi standar dari elemen array
func tampilkanStandarDeviasi(array_2138 []int, rata float64) float64 { // array_2138 - array dari tipe int, rata - rata-rata elemen array
	sumSquares_2138 := 0.0                  // Variabel untuk menyimpan jumlah kuadrat selisih
	for _, value_2138 := range array_2138 { // Loop untuk setiap elemen dalam array
		sumSquares_2138 += math.Pow(float64(value_2138)-rata, 2) // Menghitung kuadrat selisih setiap elemen dari rata-rata
	}
	return math.Sqrt(sumSquares_2138 / float64(len(array_2138))) // Mengembalikan akar kuadrat dari jumlah kuadrat selisih dibagi jumlah elemen
}

// Fungsi untuk menghitung frekuensi kemunculan elemen tertentu dalam array
func tampilkanFrekuensi(array_2138 []int, target_2138 int) int { // array_2138 - array dari tipe int, target_2138 - nilai yang ingin dihitung frekuensinya
	count_2138 := 0                         // Variabel untuk menghitung frekuensi
	for _, value_2138 := range array_2138 { // Loop untuk setiap elemen dalam array
		if value_2138 == target_2138 { // Memeriksa apakah elemen sama dengan target_2138
			count_2138++ // Menambahkan 1 ke frekuensi jika cocok
		}
	}
	return count_2138 // Mengembalikan frekuensi kemunculan target_2138
}

func main() {
	var n_2138, x_2138, hapusIndeks_2138, target_2138 int

	// Input jumlah elemen array dari pengguna
	fmt.Print("Masukkan jumlah elemen array (n): ")
	fmt.Scan(&n_2138) // Membaca jumlah elemen array

	// Membuat array dengan ukuran sesuai input
	dataArray_2138 := make([]int, n_2138)        // Membuat array dengan panjang n_2138
	for i_2138 := 0; i_2138 < n_2138; i_2138++ { // Loop untuk mengisi elemen array
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i_2138)
		fmt.Scan(&dataArray_2138[i_2138]) // Membaca nilai elemen dari pengguna
	}

	// Loop untuk menampilkan menu operasi array
	for {
		fmt.Println("\nPROGRAM OPERASI ARRAY BY RICO ADE PRATAMA 2311102138:")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan nilai ganjil")
		fmt.Println("3. Tampilkan elemen dengan nilai genap")
		fmt.Println("4. Tampilkan elemen pada indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata dari elemen array")
		fmt.Println("7. Tampilkan standar deviasi elemen array")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
		fmt.Println("9. Keluar")

		// Input pilihan_2138 operasi dari pengguna
		var pilihan_2138 int
		fmt.Print("Masukkan Nomor Operasi yang dipilih: ")
		fmt.Scan(&pilihan_2138) // Membaca pilihan_2138 operasi dari pengguna

		// Menjalankan operasi sesuai pilihan_2138 pengguna
		switch pilihan_2138 {
		case 1:
			tampilkanSEluruhArray(dataArray_2138) // Menampilkan seluruh isi array

		case 2:
			tampilkanIndeksGanjil(dataArray_2138) // Menampilkan elemen dengan nilai ganjil

		case 3:
			tampilkanIndeksGenap(dataArray_2138) // Menampilkan elemen dengan nilai genap

		case 4:
			fmt.Print("Masukkan nilai (x) untuk kelipatan indeks: ")
			fmt.Scan(&x_2138)                           // Membaca nilai x_2138 dari pengguna
			tampilkanKelipatanX(dataArray_2138, x_2138) // Menampilkan elemen pada indeks kelipatan x_2138

		case 5:
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&hapusIndeks_2138)                                          // Membaca indeks yang akan dihapus
			if hapusIndeks_2138 >= 0 && hapusIndeks_2138 < len(dataArray_2138) { // Memeriksa apakah indeks valid
				dataArray_2138 = hapusElemen(dataArray_2138, hapusIndeks_2138) // Menghapus elemen pada indeks tertentu
				tampilkanSEluruhArray(dataArray_2138)                          // Menampilkan array setelah elemen dihapus
			} else {
				fmt.Println("Indeks tidak valid!") // Menampilkan pesan jika indeks tidak valid
			}

		case 6:
			mean_2138 := tampilkanRataRata(dataArray_2138)          // Menghitung rata-rata elemen array
			fmt.Printf("Rata-rata elemen array: %.2f\n", mean_2138) // Menampilkan rata-rata elemen array

		case 7:
			mean_2138 := tampilkanRataRata(dataArray_2138)                            // Menghitung rata-rata elemen array
			standarDeviasi_2138 := tampilkanStandarDeviasi(dataArray_2138, mean_2138) // Menghitung deviasi standar elemen array
			fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi_2138)   // Menampilkan deviasi standar elemen array

		case 8:
			fmt.Print("Masukkan nilai untuk mencari frekuensi: ")
			fmt.Scan(&target_2138)                                                         // Membaca nilai target_2138 dari pengguna
			frekuensi_2138 := tampilkanFrekuensi(dataArray_2138, target_2138)              // Menghitung frekuensi target_2138 dalam array
			fmt.Printf("Frekuensi %d dalam array: %d kali\n", target_2138, frekuensi_2138) // Menampilkan frekuensi target_2138

		case 9:
			fmt.Println("Keluar dari program.") // Keluar dari program
			return

		default:
			fmt.Println("Pilihan tidak valid!") // Menampilkan pesan jika pilihan_2138 tidak valid
		}
	}
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo2_Part1.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo2.go_Part1.png)
![Output_TugasNo2_Part2.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo2.go_Part2.png)
![Output_TugasNo2_Part3.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo2.go_Part3.png)
![Output_TugasNo2_Part4.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo2.go_Part4.png)

Kode di atas digunakan untuk melakukan berbagai operasi terhadap array. Program ini adalah "PROGRAM OPERASI ARRAY" yang dapat, seperti menampilkan elemen, menampilkan elemen ganjil/genap, menghitung rata-rata, deviasi standar, serta menghapus dan mencari frekuensi elemen dalam array. Pengguna diberikan menu untuk memilih operasi yang ingin dijalankan pada array, dan hasilnya akan ditampilkan setelah setiap operasi. Lebih jelasnya seperti pada output diatas.

## 2. TugasNo3.go

```go
// TUGAS 3 : PRAKTIKUM 5
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM UNTUK MEREKAP SKOR PERTANDINGAN 2 BUAH KLUB BOLA YANG BERLAGA
package main

import "fmt" // Mengimpor paket fmt untuk I/O

func main() {
	// Deklarasi variabel untuk nama klub A dan klub B
	var klubA_2138, klubB_2138 string
	// Deklarasi variabel untuk skor masing-masing klub
	var scoreA_2138, scoreB_2138 int
	// Slice untuk menyimpan hasil pertandingan
	hasil_2138 := []string{}
	// Variabel untuk menghitung jumlah pertandingan
	jumlahMatch_2138 := 1

	// Input nama klub A
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA_2138)
	// Input nama klub B
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB_2138)

	// Loop untuk input skor pertandingan
	for {
		// Input skor pertandingan untuk masing-masing klub
		fmt.Printf("Pertandingan %d: ", jumlahMatch_2138)
		fmt.Scan(&scoreA_2138, &scoreB_2138)

		// Jika skor negatif, keluar dari loop
		if scoreA_2138 < 0 || scoreB_2138 < 0 {
			break
		}

		// Menentukan pemenang atau draw
		if scoreA_2138 > scoreB_2138 {
			hasil_2138 = append(hasil_2138, klubA_2138)
		} else if scoreB_2138 > scoreA_2138 {
			hasil_2138 = append(hasil_2138, klubB_2138)
		} else {
			hasil_2138 = append(hasil_2138, "Draw")
		}

		// Tambah jumlah pertandingan
		jumlahMatch_2138++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nDaftar hasil pertandingan:")
	for i_2138, poin_2138 := range hasil_2138 {
		if poin_2138 == "Draw" {
			fmt.Printf("Hasil %d: Draw\n", i_2138+1)
		} else {
			fmt.Printf("Hasil %d: %s\n", i_2138+1, poin_2138)
		}
	}

	// Menampilkan pesan bahwa pertandingan selesai
	fmt.Println("Pertandingan selesai")
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo3.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo3.go.png)

Kode di atas digunakan untuk merekap skor pertandingan dua klub sepak bola. Pengguna memasukkan nama dua klub dan skor setiap pertandingan. Program terus meminta skor pertandingan hingga skor negatif dimasukkan, yang menandakan akhir input. Program mencatat hasil setiap pertandingan, baik kemenangan untuk salah satu klub atau seri (draw), dan kemudian mencetak daftar hasil pertandingan yang telah direkap. Lebih jelasnya seperti pada output diatas.

## 3. TugasNo4.go

```go
// TUGAS 4 : PRAKTIKUM 5
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PALINDROM
package main

import (
	"bufio"   // bufio menyediakan buffered I/O, digunakan untuk membaca input dari pengguna
	"fmt"     // Mengimpor paket fmt untuk I/O
	"os"      // untuk bekerja dengan sistem operasi, seperti membaca input
	"strings" // untuk memanipulasi string, seperti perubahan huruf besar/kecil
)

const NMAX int = 127 // NMAX adalah konstanta batas dari maksimum elemen array

type tabel [NMAX]rune // tabel adalah array karakter dengan batas maksimum NMAX, dan rune berfungsi untuk menyimpan karakter unicode atau encoding karakter

// Fungsi untuk isi Array
/*
I.S. Data tersedia dalam piranti masukan
F.S. Array t berisi sejumlah n karakter yang dimasukkan user,

	Proses input selama karakter bukanlah TITIK dan n <= NMAX
*/
func isiArray(t *tabel, n *int, line string) {
	*n = 0                      // inisialisasi 0 untuk jumlah elemen array
	for _, char := range line { // untuk setiap karakter pada baris input
		if *n >= NMAX { // jika jumlah elemen array sudah mencapai batas maksimum
			break // hentikan loop
		}
		t[*n] = char // masukkan karakter ke dalam array
		*n++         // tambahkan jumlah elemen array
	}
}

// Fungsi untuk cetak Array
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. n karakter dalam array muncul di layar
*/
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ { // untuk setiap elemen pada array
		fmt.Print(string(t[i]), " ") // cetak elemen
	}
	fmt.Println() // mencetak baris baru setelah mencetak semua elemen
}

// Fungsi untuk membalikkan isi array
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. Urutan isi array t terbalik
*/
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		t[i], t[n-1-i] = t[n-1-i], t[i] // tukar elemen
	}
}

// Fungsi utama untuk memeriksa apakah array membentuk palindrom
/*
I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. Menghasilkan nilai true jika array t adalah palindrom, false jika tidak
*/
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ { // untuk setiap elemen pada setengah array
		if t[i] != t[n-1-i] { // jika elemen tidak sama
			return false // array tidak membentuk palindrom
		}
	}
	return true // array membentuk palindrom
}

// Fungsi utama
func main() {
	// si array tab dengan memanggil prosedur isiArray
	var tab tabel                         // Deklarasi variabel array
	var m int                             // Deklarasi variabel untuk jumlah elemen
	scanner := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna
	fmt.Println("Masukkan teks (ketik '.' untuk berhenti): ")

	for scanner.Scan() { // Loop untuk membaca input pengguna
		line := scanner.Text()            // Membaca satu baris input
		if strings.ToUpper(line) == "." { // Jika input adalah titik, hentikan loop
			break // keluar dari loop
		}

		// Mengisi array dengan baris input
		isiArray(&tab, &m, line) // si array tab dengan memanggil prosedur isiArray

		// Menampilkan array asli
		fmt.Print("Teks         : ")
		cetakArray(tab, m) // Memanggil fungsi untuk mencetak array

		// Balikan isi array tab dengan memanggil balkanArray
		balikanArray(&tab, m) // Balikian isi array tab
		fmt.Print("Reverse Teks : ")
		cetakArray(tab, m) // Cetak isi array tab

		// Mengecek dan menampilkan hasil palindrom
		isPalindrom := palindrom(tab, m)           // Memeriksa apakah teks merupakan palindrom
		fmt.Println("Palindrom    ?", isPalindrom) // Menampilkan hasil palindrom
		fmt.Println()                              // mencetak baris baru untuk pemisahan
	}
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo4.go.png](/PRAKTIKUM%205/Assets%20dan%20Laprak/Output_TugasNo4.go.png)

Kode di atas digunakan untuk memeriksa apakah sebuah teks yang dimasukkan pengguna adalah palindrom. Program menerima input teks, kemudian menampilkan teks asli dan teks yang dibalik. Setelah itu, program memeriksa apakah teks tersebut adalah palindrom (teks yang sama jika dibaca dari depan maupun belakang). Program menggunakan array untuk menyimpan karakter dan melakukan operasi pada array tersebut, termasuk pembalikan urutan dan pengecekan palindrom. Lebih jelasnya seperti pada output diatas.

