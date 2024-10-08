# <h1 align="center">LAPORAN PRAKTIKUM MODUL 2 : REVIEW STRUKTUR KONTROL</h1>
<p align="center">RICO ADE PRATAMA - 2311102138</p>

## LATIHAN
### 1. Hello.go

```go
package main

import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

```
## OUTPUT SCREENSHOT
![Output_hello.go.png](/PRAKTIKUM%202/Assets/Output_hello.go.png)

### 2. Hipotesa.go

```go
package main

import "fmt"

func main() {
	var a, b, c float64
	var hipotenusa bool
	fmt.Print("Masukkan Nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan Nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan Nilai C: ")
	fmt.Scanln(&c)
	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa a, b dan c : ", hipotenusa)
}

```
## OUTPUT SCREENSHOT
![Output_hipotesa.go.png](/PRAKTIKUM%202/Assets/Output_hipotesa.go.png)

### 3. Latihan_1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukkan Input String: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan Input String: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan Input String: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal + " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("output akhir = " + satu + " " + dua + " " + tiga)
}

```
## OUTPUT SCREENSHOT
![Output_latihan1.go.png](/PRAKTIKUM%202/Assets/Output_latihan1.go.png)

### 4. Latihan_2.go

```go
package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan Tahun : ")
	fmt.Scanln(&tahun)

	if tahun%4 == 0 {
		fmt.Println("Tahun kabisat : True")
	} else {
		fmt.Println("tahun Kabisat : False")
	}
}

```
## OUTPUT SCREENSHOT
![Output_latihan2.go.png](/PRAKTIKUM%202/Assets/Output_latihan2.go.png)

## TUGAS
### 1. TUGAS_No3_PRAKTIKUM2.go

```go
// TUGAS 3 : PRAKTIKUM 2
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MENGHITUNG VOLUME DAN LUAS BOLA
package main

import (
	"fmt"
	"math"
)

func main() {
	var r_2138 float64
	fmt.Print("<<<<<<< PROGRAM MENGHITUNG VOLUME DAN LUAS BOLA BY RICO ADE PRATAMA (2311102138) >>>>>>>\n")
	fmt.Print(">> Jejari = ")
	fmt.Scanln(&r_2138)

	// Menghitung volume bola
	volume_2138 := (4.0 / 3.0) * math.Pi * math.Pow(r_2138, 3)

	// Menghitung luas permukaan bola
	luas_2138 := 4 * math.Pi * math.Pow(r_2138, 2)

	fmt.Printf("== Bola dengan jari-jari %.0f memiliki volume %.4f dan luas permukaan %.4f\n", r_2138, volume_2138, luas_2138)
}


```
## OUTPUT SCREENSHOT
![Output_TugasNo3_Praktikum2.go.png](/PRAKTIKUM%202/Assets/Output_TugasNo3_Praktikum2.go.png)

### 2. TUGAS_No4_PRAKTIKUM2.go

```go
// TUGAS 4 : PRAKTIKUM 2
// RICO ADE PRATAMA
// 2311102138
// S1IF-11-01

// PROGRAM MEMBACA NILAI TEMPERATUR DALAM DERAJAT CELSIUS
package main

import (
	"fmt"
)

func main() {
	var celsius_2138 float64
	fmt.Print("<<<<<<< PROGRAM MEMBACA NILAI TEMPERATUR DALAM DERAJAT CELSIUS BY RICO ADE PRATAMA (2311102138) >>>>>>>\n")

	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius_2138)

	// Membaca nilai temperatur dalam derajat celsius
	reamur_2138 := celsius_2138 * 4 / 5
	fahrenheit_2138 := (celsius_2138 * 9 / 5) + 32
	kelvin_2138 := celsius_2138 + 273

	fmt.Println("Derajat Reamur: ", reamur_2138)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit_2138)
	fmt.Println("Derajat Kelvin: ", kelvin_2138)
}


```
## OUTPUT SCREENSHOT
![Output_TugasNo4_Praktikum2.go.png](/PRAKTIKUM%202/Assets/Output_TugasNo4_Praktikum2.go.png)

### 3. TUGAS_No5_PRAKTIKUM2.go

```go
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

```
## OUTPUT SCREENSHOT
![Output_TugasNo5_Praktikum2.go.png](/PRAKTIKUM%202/Assets/Output_TugasNo5_Praktikum2.go.png)
