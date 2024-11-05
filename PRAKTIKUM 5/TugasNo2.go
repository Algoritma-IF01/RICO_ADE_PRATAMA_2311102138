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
