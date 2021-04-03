package main

import "fmt"

func main() {
	var budget, wisata, inap int
	var tiket, biaya_inap, konsumsi, total int
	fmt.Print("Masukan budget dalam dollar: ")
	fmt.Scanln(&budget)
	fmt.Print("Masukan jumlah tempat wisata yang dikunjungi: ")
	fmt.Scanln(&wisata)
	fmt.Print("Lama menginap: ")
	fmt.Scanln(&inap)

	hitungBiayaSatuan(wisata, inap, &konsumsi, &tiket, &biaya_inap)
	hitungBiayaTotal(&tiket, &biaya_inap, &konsumsi, &total)

	fmt.Printf("Biaya tiket masuk: $%d \n", tiket)
	fmt.Printf("Biaya menginap: $%d (Rp. %d) \n", biaya_inap, USDtoDollar(biaya_inap))
	fmt.Printf("Biaya konsumsi: $%d \n", konsumsi)
	fmt.Printf("Total biaya perjalanan: $%d (Rp. %d) \n", total, USDtoDollar(total))
	fmt.Printf("Total Kekurangan/Kelebihan: $%d (Rp. %d) \n", budget-total, USDtoDollar(budget-total))
}

func USDtoDollar(x int) int {
	return x * 13000
}

func hitungBiayaSatuan(wisata, inap int, konsumsi, tiket, biaya_inap *int) {
	*tiket = 13 * wisata
	*biaya_inap = 514 * inap
	*konsumsi = 2 * inap * 20
}
func hitungBiayaTotal(tiket, biaya_inap, konsumsi, total *int) {
	*total = *tiket + *biaya_inap + *konsumsi + 450
}
