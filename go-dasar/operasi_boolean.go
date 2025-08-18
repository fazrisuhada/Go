package main

import "fmt"

func main() {
	var nilaiAhir = 90
	var nilaiKehadiran = 80

	var lulusNilaiAkhir = nilaiAhir > 90
	var lulusNilaiKehadiran = nilaiKehadiran > 80

	var lulus = lulusNilaiAkhir && lulusNilaiKehadiran
	fmt.Println(lulus)
}
