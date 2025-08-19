package main

import "fmt"

func main() {
	fruits := []string{"Apel", "Alpukat", "Salak", "Mangga", "Jeruk", "Buah Naga", "Sirsak", "Pir", "Pisang"}

	fmt.Println(fruits)

	sliceFruits1 := fruits[3:8]
	fmt.Println(sliceFruits1)

	sliceFruits2 := fruits[:3]
	fmt.Println(sliceFruits2)

	sliceFruits3 := len(fruits[3:])
	fmt.Println(sliceFruits3)

	sliceFruits4 := fruits[:]
	fmt.Println(sliceFruits4)

	// append slice
	days := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Saturday" //saat kita mengubah data slice, data di array ikut berubah
	daysSlice1[1] = "Sunday"

	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Tanggal Merah")
	fmt.Println(daysSlice2)
	fmt.Println(daysSlice1)
	fmt.Println(days)

	// make slice
	var makanan = make([]string, 2, 5)
	makanan[0] = "Rendang"
	makanan[1] = "Dendeng"

	fmt.Println(makanan)
	fmt.Println(len(makanan))
	fmt.Println(cap(makanan))

	makananNew := append(makanan, "Sate")
	fmt.Println(makananNew)

	makananNew[0] = "Tunjang"
	fmt.Println(makananNew)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}
