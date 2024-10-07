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
