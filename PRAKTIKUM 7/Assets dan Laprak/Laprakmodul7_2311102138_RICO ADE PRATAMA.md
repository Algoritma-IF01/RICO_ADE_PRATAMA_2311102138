# <h1 align="center">LAPORAN PRAKTIKUM MODUL 7 : PENGURUTAN DATA SELECTION SORT DAN INSERTION SORT</h1>
## <p align="center">RICO ADE PRATAMA - 2311102138</p>

# LATIHAN
## 1. Selection.go

```go
package main

import "fmt"

type arrInt [4321]int

func selectionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara ascending atau membesar dengan SELECTION SORT */
	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idx_min] {
				idx_min = j
			}
		}
		// Tukar elemen T[i] dengan T[idx_min] jika perlu
		if idx_min != i {
			T[i], T[idx_min] = T[idx_min], T[i]
		}
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 64, 34, 25, 12, 22

	fmt.Println("Array sebelum diurutkan:", T[:n])
	selectionSort1(&T, n)
	fmt.Println("Array setelah diurutkan:", T[:n])
}

```
### OUTPUT SCREENSHOT
![Output_Selection.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_Selection.go.png)

Kode di atas digunakan untuk mengimplementasikan Selection Sort dengan mengurutkan array bilangan bulat secara ascending (dari kecil ke besar). Program mencetak array sebelum dan setelah diurutkan. Fungsi utama, `selectionSort1`, mencari elemen terkecil dalam setiap iterasi dan menempatkannya pada posisi yang sesuai hingga seluruh array terurut. Lebih jelasnya seperti pada output diatas.

## 2. SelectionStruct.go

```go
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func selectionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara ascending berdasarkan ipk dengan
	   menggunakan algoritma SELECTION SORT */

	var idx_min int
	var temp mahasiswa

	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min = i

		// Cari elemen dengan IPK terkecil di subarray [i+1, n-1]
		for j := i + 1; j < n; j++ {
			if T[j].ipk < T[idx_min].ipk {
				idx_min = j
			}
		}

		// Tukar elemen di indeks i dengan elemen di idx_min jika perlu
		if idx_min != i {
			temp = T[i]
			T[i] = T[idx_min]
			T[idx_min] = temp
		}
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[1] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[2] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	selectionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}

```
### OUTPUT SCREENSHOT
![Output_SelectionStruct.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_SelectionStruct.go.png)

Kode di atas digunakan untuk mengurutkan data mahasiswa berdasarkan IPK secara ascending menggunakan Selection Sort, mencetak data sebelum dan sesudah diurutkan. Data mahasiswa meliputi nama, NIM, kelas, jurusan, dan IPK. Lebih jelasnya seperti pada output diatas.

## 3. Insertion.go

```go
package main

import "fmt"

type arrInt [4321]int

func insertionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara mengecil (descending) dengan INSERTION SORT */
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya yang lebih kecil dari temp
		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 22, 12, 34, 64, 25

	fmt.Println("Array sebelum diurutkan:", T[:n])
	insertionSort1(&T, n)
	fmt.Println("Array setelah diurutkan secara descending:", T[:n])
}

```
### OUTPUT SCREENSHOT
![Output_Insertion.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_Insertion.go.png)

Kode di atas digunakan untuk mengurutkan array bilangan bulat secara descending (dari besar ke kecil) menggunakan algoritma Insertion Sort. Array dicetak sebelum dan setelah diurutkan untuk menunjukkan hasil pengurutan. Lebih jelasnya seperti pada output diatas.

## 4. InsertionStruct.go

```go
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func insertionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara mengecil (descending) berdasarkan nama
	   dengan menggunakan algoritma INSERTION SORT */
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya
		for j > 0 && temp.nama > T[j-1].nama {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[1] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[2] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	insertionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan nama (descending):")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}

```
### OUTPUT SCREENSHOT
![Output_InsertionStruct.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_InsertionStruct.go.png)

Kode di atas digunakan untuk mengurutkan data mahasiswa berdasarkan nama secara descending menggunakan algoritma Insertion Sort. Data mahasiswa mencakup nama, NIM, kelas, jurusan, dan IPK. Program mencetak data sebelum dan sesudah diurutkan untuk menunjukkan hasil pengurutan. Lebih jelasnya seperti pada output diatas.

# TUGAS
## 1. TugasNo1_SelectionSort.go

```go
// TUGAS 1 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM RUMAH KERABAT YANG AKAN MENYUSUN NOMOR-NOMOR RUMAH KERABATNYA SECARA TERURUT MEMBESAR MENGGUNAKAN ALGORITMA SELECTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		minIndex_2138 := i_2138 // Indeks elemen minimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] < array_2138[minIndex_2138] { // Bandingkan elemen saat ini dengan elemen minimum
				minIndex_2138 = j_2138 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[minIndex_2138] = array_2138[minIndex_2138], array_2138[i_2138]
	}
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna

	// Meminta jumlah daerah
	fmt.Print("[2138] Masukkan jumlah daerah: ")
	scanner_2138.Scan()
	jumlahDaerah_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah daerah ke integer

	for i_2138 := 1; i_2138 <= jumlahDaerah_2138; i_2138++ {
		// Meminta input untuk setiap daerah
		fmt.Printf(">> Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2138)
		scanner_2138.Scan()
		input_2138 := scanner_2138.Text()        // Membaca input pengguna sebagai string
		parts_2138 := strings.Fields(input_2138) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2138, _ := strconv.Atoi(parts_2138[0]) // Konversi jumlah rumah dari string ke integer
		nomorRumah_2138 := make([]int, jumlahRumah_2138)   // Buat array untuk menyimpan nomor rumah
		for j_2138 := 0; j_2138 < jumlahRumah_2138; j_2138++ {
			nomorRumah_2138[j_2138], _ = strconv.Atoi(parts_2138[j_2138+1]) // Konversi setiap nomor rumah ke integer
		}

		// Urutkan nomor rumah menggunakan Selection Sort (ascending)
		SelectionSortAsc_2138(nomorRumah_2138) // Panggil fungsi untuk mengurutkan nomor rumah

		// Tampilkan hasil untuk daerah saat ini
		fmt.Printf("== Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2138)
		fmt.Println(strings.Trim(fmt.Sprint(nomorRumah_2138), "[]")) // Cetak nomor rumah terurut tanpa tanda kurung
	}
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo1_SelectionSort.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_TugasNo1_SelectionSort.go.png)

Kode di atas digunakan untuk mengurutkan nomor rumah kerabat secara ascending (dari kecil ke besar) menggunakan algoritma Selection Sort. Pengguna memasukkan jumlah daerah dan nomor rumah masing-masing daerah. Program mencetak nomor rumah yang sudah terurut untuk setiap daerah setelah pengurutan. Lebih jelasnya seperti pada output diatas.

## 2. TugasNo2_SelectionSort.go

```go
// TUGAS 2 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM KERABAT DEKAT YANG AKAN MENAMPILKAN NOMOR RUMAH MULAI DARI NOMOR YANG GANJIL LEBIH DAHULU TERURUT MEMBESAR DAN KEMUDIAN MENAMPILKAN NOMOR RUMAH DENGAN NOMER GENAP TERURUT MENGECIL
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk memisahkan bilangan ganjil dan genap dari array_2138
func PisahkanGanjilGenap_2138(array_2138 []int) (ganjil_2138 []int, genap_2138 []int) {
	for _, angka_2138 := range array_2138 {
		if angka_2138%2 == 0 {
			genap_2138 = append(genap_2138, angka_2138) // Tambahkan ke array genap_2138
		} else {
			ganjil_2138 = append(ganjil_2138, angka_2138) // Tambahkan ke array ganjil_2138
		}
	}
	return
}

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		minIndex_2138 := i_2138 // Indeks elemen minimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] < array_2138[minIndex_2138] { // Jika elemen lebih kecil ditemukan
				minIndex_2138 = j_2138 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[minIndex_2138] = array_2138[minIndex_2138], array_2138[i_2138]
	}
}

// Fungsi untuk mengurutkan array_2138 secara descending (turun) menggunakan Selection Sort
func SelectionSortDesc_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 0; i_2138 < n_2138-1; i_2138++ {
		maksIndex_2138 := i_2138 // Indeks elemen maksimum
		for j_2138 := i_2138 + 1; j_2138 < n_2138; j_2138++ {
			if array_2138[j_2138] > array_2138[maksIndex_2138] { // Jika elemen lebih besar ditemukan
				maksIndex_2138 = j_2138 // Perbarui indeks elemen maksimum
			}
		}
		// Tukar elemen
		array_2138[i_2138], array_2138[maksIndex_2138] = array_2138[maksIndex_2138], array_2138[i_2138]
	}
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna baris demi baris

	fmt.Print("[2138] Masukkan jumlah daerah: ")
	scanner_2138.Scan()
	jumlahDaerah_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah daerah ke integer

	for i_2138 := 0; i_2138 < jumlahDaerah_2138; i_2138++ {
		// Meminta input untuk setiap daerah
		fmt.Printf(">> Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2138+1)
		scanner_2138.Scan()
		baris_2138 := scanner_2138.Text()         // Membaca input pengguna sebagai string
		elemen_2138 := strings.Fields(baris_2138) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2138, _ := strconv.Atoi(elemen_2138[0]) // Konversi jumlah rumah ke integer
		nomorRumah_2138 := make([]int, jumlahRumah_2138)    // Buat array untuk menyimpan nomor rumah
		for j_2138 := 0; j_2138 < jumlahRumah_2138; j_2138++ {
			nomorRumah_2138[j_2138], _ = strconv.Atoi(elemen_2138[j_2138+1]) // Konversi setiap nomor rumah ke integer
		}

		// Pisahkan angka ganjil dan genap
		ganjil_2138, genap_2138 := PisahkanGanjilGenap_2138(nomorRumah_2138)

		// Urutkan ganjil dari kecil ke besar
		SelectionSortAsc_2138(ganjil_2138)

		// Urutkan genap dari besar ke kecil
		SelectionSortDesc_2138(genap_2138)

		// Gabungkan hasil ganjil diikuti genap
		hasil_2138 := append(ganjil_2138, genap_2138...) // Gabungkan ganjil dan genap

		// Tampilkan output untuk daerah tersebut
		fmt.Printf("== Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2138+1)
		fmt.Println(strings.Trim(fmt.Sprint(hasil_2138), "[]")) // Cetak hasil tanpa tanda kurung
	}
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo2_SelectionSort.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_TugasNo2_SelectionSort.go.png)

Kode di atas digunakan untuk memproses nomor rumah kerabat di setiap daerah dengan menampilkan nomor ganjil terlebih dahulu secara ascending (dari kecil ke besar), diikuti nomor genap secara descending (dari besar ke kecil). Pengguna memasukkan jumlah daerah dan nomor rumah untuk setiap daerah, kemudian program mencetak hasil pengurutan untuk masing-masing daerah. Lebih jelasnya seperti pada output diatas.

## 3. TugasNo1_InsertionSort.go

```go
// TUGAS 3 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM UNTUK MEMBACA DATA INTEGER YANG KEMUDIAN DIURUTKAN MENGGUNAKAN METODE INSERTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2138 secara ascending (naik) menggunakan Insertion Sort
func InsertionSort_2138(array_2138 []int) {
	n_2138 := len(array_2138) // Panjang array
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		key_2138 := array_2138[i_2138]                     // Elemen yang sedang diurutkan
		j_2138 := i_2138 - 1                               // Indeks elemen sebelumnya
		for j_2138 >= 0 && array_2138[j_2138] > key_2138 { // Geser elemen yang lebih besar
			array_2138[j_2138+1] = array_2138[j_2138]
			j_2138--
		}
		array_2138[j_2138+1] = key_2138 // Tempatkan elemen di posisi yang benar
	}
}

// Fungsi untuk memeriksa apakah array_2138 memiliki jarak tetap
func IsJarakTetap_2138(array_2138 []int) (bool_2138 bool, jarak_2138 int) {
	if len(array_2138) < 2 { // Jika elemen kurang dari 2
		return true, 0 // Secara default dianggap jarak tetap
	}
	jarak_2138 = array_2138[1] - array_2138[0] // Hitung jarak awal
	for i_2138 := 2; i_2138 < len(array_2138); i_2138++ {
		if array_2138[i_2138]-array_2138[i_2138-1] != jarak_2138 { // Jika jarak berbeda
			return false, 0 // Tidak berjarak tetap
		}
	}
	return true, jarak_2138 // Jika jarak tetap, kembalikan true dan jarak
}

func main() {
	scanner_2138 := bufio.NewScanner(os.Stdin) // Scanner untuk membaca input

	// Meminta jumlah kasus
	fmt.Print("[2138] Masukkan jumlah kasus: ")
	scanner_2138.Scan()
	jumlahKasus_2138, _ := strconv.Atoi(scanner_2138.Text()) // Konversi jumlah kasus ke integer

	for kasus_2138 := 1; kasus_2138 <= jumlahKasus_2138; kasus_2138++ {
		// Membaca input untuk setiap kasus
		fmt.Printf(">> Masukkan data untuk kasus ke-%d (akhiri dengan bilangan negatif): ", kasus_2138)
		scanner_2138.Scan()
		baris_2138 := scanner_2138.Text()         // Membaca input sebagai string
		elemen_2138 := strings.Fields(baris_2138) // Memisahkan string menjadi array
		array_2138 := make([]int, 0)              // Array untuk menyimpan bilangan non-negatif

		for _, el_2138 := range elemen_2138 { // Iterasi setiap elemen input
			num_2138, _ := strconv.Atoi(el_2138) // Konversi string ke integer
			if num_2138 < 0 {                    // Jika negatif, hentikan proses
				break
			}
			array_2138 = append(array_2138, num_2138) // Tambahkan ke array
		}

		// Urutkan array menggunakan Insertion Sort
		InsertionSort_2138(array_2138)

		// Periksa apakah array berjarak tetap
		jarakTetap_2138, jarak_2138 := IsJarakTetap_2138(array_2138)

		// Cetak hasil pengurutan
		fmt.Printf("== Hasil Pengurutan: %v\n", array_2138)
		if jarakTetap_2138 {
			fmt.Printf("=> Data berjarak %d\n", jarak_2138) // Jika jarak tetap, cetak jaraknya
		} else {
			fmt.Println("=> Data berjarak tidak tetap") // Jika tidak tetap, cetak pesan ini
		}
	}
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo1_InsertionSort.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_TugasNo1_InsertionSort.go.png)

Kode di atas digunakan untuk membaca sejumlah data integer, mengurutkannya secara ascending (dari kecil ke besar) menggunakan algoritma Insertion Sort, dan memeriksa apakah data memiliki jarak tetap antar elemen. Hasil pengurutan serta informasi tentang jarak tetap ditampilkan untuk setiap kasus yang dimasukkan pengguna. Lebih jelasnya seperti pada output diatas.

## 4. TugasNo2_InsertionSort.go

```go
// TUGAS 4 : PRAKTIKUM 7
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM PERPUSTAKAAN UNTUK MENGELOLA DATA BUKU DI DALAM SUATU PERPUSTAKAAN MENGGUNAKAN INSERTION SORT
package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

const nMax_2138 = 7919 // Menentukan jumlah maksimum buku yang dapat ditampung

// Struct Buku_2138 untuk menyimpan informasi tentang buku
type Buku_2138 struct {
	id_2138       int    // ID buku
	judul_2138    string // Judul buku
	penulis_2138  string // Penulis buku
	penerbit_2138 string // Penerbit buku
	tahun_2138    int    // Tahun terbit
	rating_2138   int    // Rating buku
}

// Fungsi untuk mengisi data buku ke dalam pustaka
func DaftarkanBuku_2138(pustaka_2138 []Buku_2138, n_2138 int) []Buku_2138 {
	reader_2138 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input
	for i_2138 := 0; i_2138 < n_2138; i_2138++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i_2138+1)

		// Membaca ID buku
		fmt.Print("ID: ")
		idInput_2138, _ := reader_2138.ReadString('\n')
		idInput_2138 = strings.TrimSpace(idInput_2138)
		id_2138, _ := strconv.Atoi(idInput_2138)

		// Membaca Judul buku
		fmt.Print("Judul: ")
		judul_2138, _ := reader_2138.ReadString('\n')
		judul_2138 = strings.TrimSpace(judul_2138)

		// Membaca Penulis buku
		fmt.Print("Penulis: ")
		penulis_2138, _ := reader_2138.ReadString('\n')
		penulis_2138 = strings.TrimSpace(penulis_2138)

		// Membaca Penerbit buku
		fmt.Print("Penerbit: ")
		penerbit_2138, _ := reader_2138.ReadString('\n')
		penerbit_2138 = strings.TrimSpace(penerbit_2138)

		// Membaca Tahun buku
		fmt.Print("Tahun: ")
		tahunInput_2138, _ := reader_2138.ReadString('\n')
		tahunInput_2138 = strings.TrimSpace(tahunInput_2138)
		tahun_2138, _ := strconv.Atoi(tahunInput_2138)

		// Membaca Rating buku
		fmt.Print("Rating: ")
		ratingInput_2138, _ := reader_2138.ReadString('\n')
		ratingInput_2138 = strings.TrimSpace(ratingInput_2138)
		rating_2138, _ := strconv.Atoi(ratingInput_2138)

		// Menambahkan data ke dalam pustaka
		pustaka_2138[i_2138] = Buku_2138{id_2138, judul_2138, penulis_2138, penerbit_2138, tahun_2138, rating_2138}
	}
	return pustaka_2138
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun menggunakan Insertion Sort
func UrutkanBuku_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		key_2138 := pustaka_2138[i_2138] // Elemen yang akan ditempatkan
		j_2138 := i_2138 - 1

		// Geser elemen yang rating-nya lebih kecil
		for j_2138 >= 0 && pustaka_2138[j_2138].rating_2138 < key_2138.rating_2138 {
			pustaka_2138[j_2138+1] = pustaka_2138[j_2138]
			j_2138--
		}
		pustaka_2138[j_2138+1] = key_2138 // Letakkan elemen di tempat yang sesuai
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	if n_2138 == 0 { // Jika tidak ada buku
		fmt.Println("Tidak ada buku di perpustakaan.") // Tampilkan pesan
		return
	}

	idxTerfavorit_2138 := 0                             // Indeks buku dengan rating tertinggi
	ratingTertinggi_2138 := pustaka_2138[0].rating_2138 // Ambil rating dari buku pertama

	// Iterasi untuk mencari buku dengan rating tertinggi
	for i_2138 := 1; i_2138 < n_2138; i_2138++ {
		if pustaka_2138[i_2138].rating_2138 > ratingTertinggi_2138 {
			ratingTertinggi_2138 = pustaka_2138[i_2138].rating_2138
			idxTerfavorit_2138 = i_2138
		}
	}

	// Tampilkan buku terfavorit dengan rating tertinggi
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		pustaka_2138[idxTerfavorit_2138].judul_2138, pustaka_2138[idxTerfavorit_2138].penulis_2138,
		pustaka_2138[idxTerfavorit_2138].penerbit_2138, pustaka_2138[idxTerfavorit_2138].tahun_2138,
		pustaka_2138[idxTerfavorit_2138].rating_2138)
}

// Fungsi untuk mencetak lima buku dengan rating tertinggi
func CetakSetTerbaru_2138(pustaka_2138 []Buku_2138, n_2138 int) {
	fmt.Println("\nTOP 5 buku dengan rating tertinggi:")
	// Pastikan buku diurutkan sebelum mencetak
	UrutkanBuku_2138(pustaka_2138, n_2138)
	for i_2138 := 0; i_2138 < 5 && i_2138 < n_2138; i_2138++ {
		fmt.Printf("Judul: %s, Rating: %d\n", pustaka_2138[i_2138].judul_2138, pustaka_2138[i_2138].rating_2138)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku_2138(pustaka_2138 []Buku_2138, n_2138 int, ratingCari_2138 int) {
	ditemukan_2138 := false // Penanda jika buku ditemukan

	// Cari buku dengan rating yang sesuai
	fmt.Printf("\nBuku dengan rating %d:\n", ratingCari_2138)
	for i_2138 := 0; i_2138 < n_2138; i_2138++ {
		if pustaka_2138[i_2138].rating_2138 == ratingCari_2138 {
			// Tampilkan data buku
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n\n",
				pustaka_2138[i_2138].judul_2138, pustaka_2138[i_2138].penulis_2138,
				pustaka_2138[i_2138].penerbit_2138, pustaka_2138[i_2138].tahun_2138,
				pustaka_2138[i_2138].rating_2138)
			ditemukan_2138 = true
		}
	}

	// Jika tidak ada buku yang ditemukan
	if !ditemukan_2138 {
		fmt.Printf("Tidak ada buku dengan rating %d\n", ratingCari_2138)
	}
}

// Program utama
func main() {
	reader_2138 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input

	// Membaca jumlah buku di perpustakaan
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	nInput_2138, _ := reader_2138.ReadString('\n')
	nInput_2138 = strings.TrimSpace(nInput_2138)
	nPustaka_2138, _ := strconv.Atoi(nInput_2138)

	// Membuat slice untuk menampung data buku
	var Pustaka_2138 = make([]Buku_2138, nPustaka_2138)

	// Memanggil fungsi untuk mengisi data buku
	Pustaka_2138 = DaftarkanBuku_2138(Pustaka_2138, nPustaka_2138)
	fmt.Println()

	// Menampilkan buku dengan rating tertinggi
	CetakTerfavorit_2138(Pustaka_2138, nPustaka_2138)

	// Menampilkan lima buku dengan rating tertinggi
	CetakSetTerbaru_2138(Pustaka_2138, nPustaka_2138)

	// Mencari buku berdasarkan rating tertentu
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	ratingCariInput_2138, _ := reader_2138.ReadString('\n')
	ratingCariInput_2138 = strings.TrimSpace(ratingCariInput_2138)
	ratingCari_2138, _ := strconv.Atoi(ratingCariInput_2138)
	CariBuku_2138(Pustaka_2138, nPustaka_2138, ratingCari_2138)
}

```
### OUTPUT SCREENSHOT
![Output_TugasNo2_InsertionSort.go.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_TugasNo2_InsertionSort.go.png)
![Output_TugasNo2_InsertionSort.go.2.png](/PRAKTIKUM%207/Assets%20dan%20Laprak/Output_TugasNo2_InsertionSort.go.2.png)

Kode di atas digunakan untuk mengelola data buku di perpustakaan dengan menampilkan, mengurutkan, dan mencari buku berdasarkan rating menggunakan Insertion Sort. Buku diurutkan secara descending (dari rating tertinggi ke terendah), menampilkan buku terfavorit, lima buku terbaik, dan memungkinkan pencarian berdasarkan rating tertentu. Data buku meliputi ID, judul, penulis, penerbit, tahun, dan rating. Lebih jelasnya seperti pada output diatas.