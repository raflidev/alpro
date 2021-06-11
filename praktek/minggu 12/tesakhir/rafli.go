package main

import "fmt"

const NMAX = 100

type Provinsi struct {
	nama     string
	populasi int
	tumbuh   float64
}

type dataProvinsi struct {
	tabel     [NMAX]Provinsi
	nProvinsi int
}

func cariProvinsi(data dataProvinsi, nama string) Provinsi {
	for i := 0; i < data.nProvinsi; i++ {
		if data.tabel[i].nama == nama {
			return data.tabel[i]
		}
	}
	return data.tabel[99]
}

func prediksi(data dataProvinsi, nama string) int {
	for i := 0; i < data.nProvinsi; i++ {
		if data.tabel[i].nama == nama {
			return int(float64(data.tabel[i].populasi)*float64(data.tabel[i].tumbuh)) + data.tabel[i].populasi
		}
	}
	return -1
}

func urutData(data *dataProvinsi) {
	var pass, idx, i int
	var temp Provinsi
	pass = 1
	for pass <= data.nProvinsi-1 {
		idx = pass - 1
		i = pass
		for i < data.nProvinsi {
			if data.tabel[idx].tumbuh < data.tabel[i].tumbuh {
				idx = i
			}
			i++
		}
		temp = data.tabel[pass-1]
		data.tabel[pass-1] = data.tabel[idx]
		data.tabel[idx] = temp
		pass = pass + 1
	}

}

func main() {
	var p dataProvinsi

	var nama_prov, nama_prov2 string
	i := 0
	fmt.Scanln(&p.tabel[i].nama, &p.tabel[i].populasi, &p.tabel[i].tumbuh)
	for i < NMAX && p.tabel[i].nama != "NONE" && p.tabel[i].populasi != 0 && p.tabel[i].tumbuh != 0.0 {
		p.nProvinsi++
		i++
		fmt.Scanln(&p.tabel[i].nama, &p.tabel[i].populasi, &p.tabel[i].tumbuh)
	}
	fmt.Print("Nama provinsi?\n")
	fmt.Scanln(&nama_prov)
	provinsi := cariProvinsi(p, nama_prov)
	if provinsi.nama != "" {
		fmt.Println(provinsi.nama, provinsi.populasi, provinsi.tumbuh)
	} else {
		fmt.Println("Provinsi tidak ada")
	}

	fmt.Print("Prediksi populasi tahun depan provinsi?\n")
	fmt.Scanln(&nama_prov2)
	if prediksi(p, nama_prov2) != -1 {
		fmt.Printf("Populasi Provinsi %s tahun depan: %d \n", nama_prov2, prediksi(p, nama_prov2))
	} else {
		fmt.Println("Provinsi tidak ada")
	}

	fmt.Println("Urutan data pertumbuhan provinsi berdasarkan populasi terurut menurun:")
	urutData(&p)
	for i := 0; i < p.nProvinsi; i++ {
		fmt.Println(p.tabel[i].nama, p.tabel[i].populasi, p.tabel[i].tumbuh)
	}
}

// Aceh 4906835 0.0201
// Bali 4104900 0.0133
// Banten 11704877 0.0131
// Bengkulu 1844800 0.0259
// Gorontalo 1115633 0.018
// Sumut 13766851 0.0261
// Yogyakarta 3553100 0.0195
// NONE 0 0.0
