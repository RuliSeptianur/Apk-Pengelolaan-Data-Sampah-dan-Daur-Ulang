package main

import "fmt"

type Sampah struct {
	Jenis     string
	Jumlah    float64
	Metode    string
	SudahDaur bool
}

var daftarSampah [1000]Sampah
var jumlahData int

func pilihJenis() string {
	var pil int
	fmt.Println("Pilih Jenis Sampah:")
	fmt.Println("1. Plastik")
	fmt.Println("2. Kertas/Karton")
	fmt.Println("3. Logam/Kaleng")
	fmt.Println("4. Sisa Makanan (Organik)")
	fmt.Print("Pilihan (1-4): ")
	fmt.Scan(&pil)

	switch pil {
	case 1:
		return "Plastik"
	case 2:
		return "Kertas"
	case 3:
		return "Logam"
	default:
		return "Organik"
	}
}

func pilihMetode() string {
	var pil int
	fmt.Println("Pilih Cara Olah (Daur Ulang):")
	fmt.Println("1. Dihancurkan")
	fmt.Println("2. Dijadikan Kompos")
	fmt.Println("3. Dilebur Kembali")
	fmt.Println("4. Langsung Dibuang")
	fmt.Print("Pilihan (1-4): ")
	fmt.Scan(&pil)

	switch pil {
	case 1:
		return "Hancur"
	case 2:
		return "Kompos"
	case 3:
		return "Lebur"
	default:
		return "Buang"
	}
}

func tambahData() {
	if jumlahData < 1000 {
		var s Sampah
		s.Jenis = pilihJenis()
		fmt.Print("Berapa Beratnya? (kg): ")
		fmt.Scan(&s.Jumlah)
		s.Metode = pilihMetode()

		var status int
		fmt.Print("Berhasil diolah? (1: Iya, 2: Tidak): ")
		fmt.Scan(&status)
		s.SudahDaur = (status == 1)

		daftarSampah[jumlahData] = s
		jumlahData++
		fmt.Println("✔ Data sudah disimpan!")
	}
}

func ubahData() {
	var nomor int
	tampilkanData()
	fmt.Print("Masukkan Nomor yang mau diganti: ")
	fmt.Scan(&nomor)
	if nomor > 0 && nomor <= jumlahData {
		i := nomor - 1
		fmt.Println("Masukkan Data Baru:")
		daftarSampah[i].Jenis = pilihJenis()
		fmt.Print("Berat Baru (kg): ")
		fmt.Scan(&daftarSampah[i].Jumlah)
		daftarSampah[i].Metode = pilihMetode()
		fmt.Println("✔ Data sudah diperbarui!")
	}
}

func hapusData() {
	var nomor int
	tampilkanData()
	fmt.Print("Masukkan Nomor yang mau dihapus: ")
	fmt.Scan(&nomor)
	if nomor > 0 && nomor <= jumlahData {
		for i := nomor - 1; i < jumlahData-1; i++ {
			daftarSampah[i] = daftarSampah[i+1]
		}
		jumlahData--
		fmt.Println("✔ Data sudah dihapus!")
	}
}

func urutkanData() {
	for i := 0; i < jumlahData-1; i++ {
		min := i
		for j := i + 1; j < jumlahData; j++ {
			if daftarSampah[j].Jumlah < daftarSampah[min].Jumlah {
				min = j
			}
		}
		daftarSampah[i], daftarSampah[min] = daftarSampah[min], daftarSampah[i]
	}
}

func cariData() {
	var nama string
	fmt.Print("Ketik Jenis Sampah yang dicari: ")
	fmt.Scan(&nama)
	ada := false
	for i := 0; i < jumlahData; i++ {
		if daftarSampah[i].Jenis == nama {
			fmt.Printf("- Ketemu: %s, Berat: %.2f kg\n", daftarSampah[i].Jenis, daftarSampah[i].Jumlah)
			ada = true
		}
	}
	if !ada {
		fmt.Println("Data tidak ketemu.")
	}
}

func tampilkanData() {
	if jumlahData == 0 {
		fmt.Println("\n--- Belum ada catatan sampah ---")
		return
	}
	fmt.Println("\n--- DAFTAR CATATAN SAMPAH ---")
	for i := 0; i < jumlahData; i++ {
		status := "Gagal"
		if daftarSampah[i].SudahDaur {
			status = "Sukses"
		}
		fmt.Printf("%d. [%s] | Berat: %.2fkg | Cara: %s | Status: %s\n",
			i+1, daftarSampah[i].Jenis, daftarSampah[i].Jumlah, daftarSampah[i].Metode, status)
	}
}

func lihatLaporan() {
	var total, sukses float64
	for i := 0; i < jumlahData; i++ {
		total += daftarSampah[i].Jumlah
		if daftarSampah[i].SudahDaur {
			sukses += daftarSampah[i].Jumlah
		}
	}
	fmt.Println("\n--- LAPORAN PETUGAS ---")
	fmt.Printf("Total Sampah Masuk: %.2f kg\n", total)
	fmt.Printf("Total Berhasil Diolah: %.2f kg\n", sukses)
}

func main() {
	var pilihan int
	for {
		fmt.Println("\n==============================")
		fmt.Println("  APLIKASI CATATAN SAMPAH")
		fmt.Println("==============================")
		fmt.Println("1. Tambah Catatan Baru")
		fmt.Println("2. Lihat Semua Catatan")
		fmt.Println("3. Ganti Data yang Salah")
		fmt.Println("4. Hapus Catatan")
		fmt.Println("5. Urutkan (Paling Ringan)")
		fmt.Println("6. Cari Nama Sampah")
		fmt.Println("7. Lihat Laporan Hasil")
		fmt.Println("0. Tutup Aplikasi")
		fmt.Println("==============================")
		fmt.Print("Pilih Nomor: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			tampilkanData()
		case 3:
			ubahData()
		case 4:
			hapusData()
		case 5:
			urutkanData()
			tampilkanData()
		case 6:
			cariData()
		case 7:
			lihatLaporan()
		case 0:
			fmt.Println("Aplikasi dimatikan. Terima kasih!")
			return
		default:
			fmt.Println("Nomor salah, pilih lagi.")
		}
	}
}
