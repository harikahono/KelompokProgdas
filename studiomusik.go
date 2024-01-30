package main

import (
	"fmt"
	"strings"
	"time"
)

type Studio struct {
	Nama       string
	Kapasitas  int
	HargaPerJam int
	Status     string
}

type Pemesanan struct {
	Tanggal    string
	Durasi     time.Duration
	Penyewa    string
	TotalHarga int
}

var studios = map[string]Studio{
	"A": {"Studio A", 4, 50000, "Tersedia"},
	"B": {"Studio B", 6, 80000, "Tersedia"},
}

func main() {
	menu()
}

func menu() {
	for {
		fmt.Println("=== Menu ===")
		fmt.Println("1. Lihat Informasi Studio")
		fmt.Println("2. Pesan Studio")
		fmt.Println("3. Keluar")

		fmt.Print("Pilih menu (1-3): ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatInformasiStudio()
		case 2:
			pesanStudio()
		case 3:
			fmt.Println("Terima kasih. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan menu tidak valid.")
		}
	}
}

func lihatInformasiStudio() {
	fmt.Println("Studio Information:")
	fmt.Println("----------------------------")
	for _, studio := range studios {
		fmt.Printf("Studio: %s\nKapasitas: %d orang\nHarga per jam: Rp %d\nStatus: %s\n", studio.Nama, studio.Kapasitas, studio.HargaPerJam, studio.Status)
		fmt.Println("----------------------------")
	}
}

func pesanStudio() {
	fmt.Println("Pesan Studio:")
	var pilihanStudio string
	for {
		fmt.Printf("Pilih Studio (A/B): ")
		fmt.Scanln(&pilihanStudio)
		pilihanStudio = strings.ToUpper(pilihanStudio)
		if pilihanStudio == "A" || pilihanStudio == "B" {
			break
		}
		fmt.Println("Pilihan Studio tidak valid.")
	}

	studio := studios[pilihanStudio]
	fmt.Printf("Masukkan tanggal pemesanan (format: 02/01/2006): ")
	var tanggalPemesanan string
	fmt.Scanln(&tanggalPemesanan)

	fmt.Printf("Masukkan durasi pemesanan (dalam jam): ")
	var durasiPemesanan int
	fmt.Scanln(&durasiPemesanan)

	fmt.Printf("Masukkan nama penyewa: ")
	var namaPenyewa string
	fmt.Scanln(&namaPenyewa)

	pemesanan := Pemesanan{
		Tanggal: tanggalPemesanan,
		Durasi:  time.Duration(durasiPemesanan) * time.Hour,
		Penyewa: namaPenyewa,
	}

	tampilkanDetailPemesanan(pemesanan, studio)
}

func tampilkanDetailPemesanan(pemesanan Pemesanan, studio Studio) {
	fmt.Printf("Pemesanan untuk %s\n", studio.Nama)
	fmt.Printf("Tanggal: %s\n", pemesanan.Tanggal)
	fmt.Printf("Durasi: %d JAM\n", int(pemesanan.Durasi.Hours()))
	fmt.Printf("Penyewa: %s\n", pemesanan.Penyewa)

	// Hitung total harga
	pemesanan.TotalHarga = int(pemesanan.Durasi.Hours()) * studio.HargaPerJam

	fmt.Printf("Total Harga: Rp %d\n", pemesanan.TotalHarga)
	fmt.Println("----------------------------")
}
