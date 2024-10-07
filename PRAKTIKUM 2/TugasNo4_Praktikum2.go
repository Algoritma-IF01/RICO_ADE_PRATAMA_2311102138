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
	// Input Celsius
	var celsius_2138 float64
	fmt.Print("<<<<<<< PROGRAM MEMBACA NILAI TEMPERATUR DALAM DERAJAT CELSIUS BY RICO ADE PRATAMA (2311102138) >>>>>>>\n")
	fmt.Print(">> Temperatur Celsius: ")
	fmt.Scanln(&celsius_2138)

	// Convert ke Reamur, Fahrenheit, dan Kelvin
	reanur_2138 := celsius_2138 * 4 / 5
	fahrenheit_2138 := (celsius_2138 * 9 / 5) + 32
	kelvin_2138 := celsius_2138 + 273.15

	// Output Program pertama
	fmt.Printf("== Derajat Fahrenheit: %.0f\n", fahrenheit_2138)

	// Input Celsius Lagi
	fmt.Print("\n<<<<<<< LANJUTAN PROGRAM >>>>>>>\n")
	fmt.Print(">> Temperatur Celsius: ")
	fmt.Scanln(&celsius_2138)

	// Lanjutan Output Program
	fmt.Printf("== Derajat Reamur: %.0f\n", reanur_2138)
	fmt.Printf("== Derajat Fahrenheit: %.0f\n", fahrenheit_2138)
	fmt.Printf("== Derajat Kelvin: %.0f\n", kelvin_2138)
}
