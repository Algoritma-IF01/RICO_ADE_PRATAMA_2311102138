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
