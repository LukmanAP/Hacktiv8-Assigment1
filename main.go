package main

import (
	"fmt"
	"os"
)

type Teman struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var arg = os.Args[1:]

	if len(arg) == 0 {
		fmt.Println("Anda Tidak Memasukan Nama")
		return
	}

	namaTeman := arg[0]
	ketemu := false

	for _, teman := range bioTeman {
		if teman.nama == namaTeman {
			ketemu = true
			mencariNamaTeman(teman.nama)
			break

		}
	}

	if !ketemu {
		fmt.Println("Nama Teman anda ", namaTeman, "tidak ditemukan.")
	}

}

var bioTeman []Teman = []Teman{
	{
		id:        1,
		nama:      "Jericho",
		alamat:    "purwokerto",
		pekerjaan: "Backend Developer",
		alasan:    "Karena Golang adalah bahasa yang lagi trend saat ini",
	},
	{
		id:        2,
		nama:      "Winsen",
		alamat:    "Banyumanik",
		pekerjaan: "Devops Developer",
		alasan:    "mau switch karir menjadi backand golang",
	},
	{
		id:        3,
		nama:      "Firman",
		alamat:    "Kendal",
		pekerjaan: "web Proraming",
		alasan:    "Menjadi expert dalam Bahasa Golang",
	},
	{
		id:        4,
		nama:      "Zeni",
		alamat:    "Demak",
		pekerjaan: "Mahasiswa",
		alasan:    "ingin mempunya karir menjadi backend developer",
	},
}

func mencariNamaTeman(nama string) {
	for _, pilihTeman := range bioTeman {
		if nama == pilihTeman.nama {
			fmt.Println("id:", pilihTeman.id)
			fmt.Println("nama:", pilihTeman.nama)
			fmt.Println("alamat:", pilihTeman.alamat)
			fmt.Println("pekerjaan:", pilihTeman.pekerjaan)
			fmt.Println("alasa:", pilihTeman.alasan)
		}
	}
}
