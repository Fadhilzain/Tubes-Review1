package main

import "fmt"

const max int = 50

type Konten struct {
	Judul    string
	Platform string
	Kategori string
	Likes    int
	Komentar int
	Shares   int
}

var tabKonten [max]Konten
var totalKonten int

func main() {
	var pilihMenu int

	tabKonten[0] = Konten{"Valorant", "YouTube", "Game", 850000, 1200, 300}
	tabKonten[1] = Konten{"Gojek", "Instagram", "Transportasi", 430000, 950, 280}
	tabKonten[2] = Konten{"Shoope", "Facebook", "E-commerce", 120000, 500, 150}
	tabKonten[3] = Konten{"Netflix", "Twitter", "Film", 940000, 200, 100}
	tabKonten[4] = Konten{"Ruangguru", "TikTok", "Edukasi", 240000, 750, 400}
	totalKonten = 5

	for pilihMenu != 9 {
		menu(&pilihMenu)

		if pilihMenu == 1 {
			tambahKonten(&tabKonten, &totalKonten)
		} else if pilihMenu == 2 {
			tampilkanKonten(&tabKonten, totalKonten)
		} else if pilihMenu == 3 {
			menuCari(&tabKonten, totalKonten)
		} else if pilihMenu == 4 {
			selectionSortByLikes(&tabKonten, totalKonten)
		} else if pilihMenu == 5 {
			urutkanByKomentar(&tabKonten, totalKonten)
		} else if pilihMenu == 6 {
			urutkanByShares(&tabKonten, totalKonten)
		} else if pilihMenu == 7 {
			ubahKonten(&tabKonten, totalKonten)
		} else if pilihMenu == 8 {
			hapusKonten(&tabKonten, &totalKonten)
		} else if pilihMenu == 9 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menu(pilihMenu *int) {
	fmt.Println("\n=== APLIKASI MANAJEMEN KONTEN ===")
	fmt.Println("1. Tambah Konten Baru")
	fmt.Println("2. Lihat Semua Konten")
	fmt.Println("3. Cari Konten (Sequential / Binary)")
	fmt.Println("4. Urutkan Berdasarkan Likes")
	fmt.Println("5. Urutkan Berdasarkan Komentar")
	fmt.Println("6. Urutkan Berdasarkan Share")
	fmt.Println("7. Ubah Konten")
	fmt.Println("8. Hapus Konten")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(pilihMenu)
}

// fungsi tambah konten no 1
func tambahKonten(dataKonten *[max]Konten, jumlah *int) {
	if *jumlah >= max {
		fmt.Println("Kapasitas penyimpanan konten sudah penuh.")
		return
	}

	var k Konten
	fmt.Println("\n=== TAMBAH KONTEN BARU ===")
	fmt.Print("Judul     : ")
	fmt.Scan(&k.Judul)
	fmt.Print("Platform  : ")
	fmt.Scan(&k.Platform)
	fmt.Print("Kategori  : ")
	fmt.Scan(&k.Kategori)
	fmt.Print("Likes     : ")
	fmt.Scan(&k.Likes)
	fmt.Print("Komentar  : ")
	fmt.Scan(&k.Komentar)
	fmt.Print("Shares    : ")
	fmt.Scan(&k.Shares)

	dataKonten[*jumlah] = k
	*jumlah++
	fmt.Println("Konten berhasil ditambahkan.")
}

// fungsi lihat semuakonten no 2
func tampilkanKonten(dataKonten *[max]Konten, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data konten.")
		return
	}

	fmt.Println("\n===== DAFTAR KONTEN =====")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("\n%d. Judul: %s\n", i+1, dataKonten[i].Judul)
		fmt.Printf("   Platform: %s\n", dataKonten[i].Platform)
		fmt.Printf("   Kategori: %s\n", dataKonten[i].Kategori)
		fmt.Printf("   Likes: %d\n", dataKonten[i].Likes)
		fmt.Printf("   Komentar: %d\n", dataKonten[i].Komentar)
		fmt.Printf("   Shares: %d\n", dataKonten[i].Shares)
		fmt.Println("------------------------------")
	}
}

// fungsi cari konten
func menuCari(dataKonten *[max]Konten, jumlah int) {
	var metode int
	fmt.Println("\n=== MENU PENCARIAN KONTEN ===")
	fmt.Println("1. Pencarian Sequential")
	fmt.Println("2. Pencarian Binary")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&metode)

	if metode == 1 {
		cariKontenSequential(dataKonten, jumlah)
	} else if metode == 2 {
		cariKontenBinary(dataKonten, jumlah)
	} else {
		fmt.Println("Pilihan metode tidak valid.")
	}
}

// fungsi cari konten sequential
func cariKontenSequential(dataKonten *[max]Konten, jumlah int) {
	var judul string
	var ketemu bool
	var i int
	fmt.Print("Masukkan Judul atau Platform yang ingin dicari (Sequential): ")
	fmt.Scan(&judul)

	ketemu = false
	for i = 0; i < jumlah; i++ {
		if dataKonten[i].Judul == judul || dataKonten[i].Platform == judul {
			ketemu = true
			fmt.Println("\nKonten ditemukan:")
			fmt.Printf("Judul: %s\n", dataKonten[i].Judul)
			fmt.Printf("Platform: %s\n", dataKonten[i].Platform)
			fmt.Printf("Kategori: %s\n", dataKonten[i].Kategori)
			fmt.Printf("Likes: %d\n", dataKonten[i].Likes)
			fmt.Printf("Komentar: %d\n", dataKonten[i].Komentar)
			fmt.Printf("Shares: %d\n", dataKonten[i].Shares)
			fmt.Println("------------------------------")
		}
	}
	if !ketemu {
		fmt.Println("Konten tidak ditemukan.")
	}
}

// fungsi mengurutkan berdasarkan like selection sort
func selectionSortByLikes(data *[max]Konten, n int) {
    var i, j, minIdx int
    var temp Konten

    for i = 0; i < n-1; i++ {
        minIdx = i
        for j = i + 1; j < n; j++ {
            if data[j].Likes < data[minIdx].Likes {
                minIdx = j
            }
        }

        // Tukar data[i] dengan data[minIdx]
        temp = data[i]
        data[i] = data[minIdx]
        data[minIdx] = temp
    }

    fmt.Println("Konten berhasil diurutkan berdasarkan Likes.")
}

// menggunakan insertion sort
func urutkanByKomentar(dataKonten *[max]Konten, jumlah int) {
	var i, j int
	var temp Konten

	for i = 1; i < jumlah; i++ {
		temp = dataKonten[i]
		j = i - 1
		for j >= 0 && dataKonten[j].Komentar < temp.Komentar {
			dataKonten[j+1] = dataKonten[j]
			j = j - 1
		}
		dataKonten[j+1] = temp
	}
	fmt.Println("Konten berhasil diurutkan berdasarkan Komentar.")
	tampilkanKonten(dataKonten, jumlah)
}

func urutkanByShares(dataKonten *[max]Konten, jumlah int) {
	var i, j, maxIdx int
	for i = 0; i < jumlah-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumlah; j++ {
			if dataKonten[j].Shares > dataKonten[maxIdx].Shares {
				maxIdx = j
			}
		}
		dataKonten[i], dataKonten[maxIdx] = dataKonten[maxIdx], dataKonten[i]
	}
	fmt.Println("Konten berhasil diurutkan berdasarkan Shares.")
	tampilkanKonten(dataKonten, jumlah)
}



// fungsi urutkan berdasarkan judul
func urutkanByJudul(dataKonten *[max]Konten, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		for j := i + 1; j < jumlah; j++ {
			if dataKonten[i].Judul > dataKonten[j].Judul {
				dataKonten[i], dataKonten[j] = dataKonten[j], dataKonten[i]
			}
		}
	}
}

// fungsi cari konten binary
func cariKontenBinary(dataKonten *[max]Konten, jumlah int) {
	var judul string
	var kiri, kanan, tengah int
	var ketemu bool

	fmt.Print("Masukkan Judul Konten yang ingin dicari (Binary): ")
	fmt.Scan(&judul)

	urutkanByJudul(dataKonten, jumlah)

	kiri = 0
	kanan = jumlah - 1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if dataKonten[tengah].Judul == judul {
			ketemu = true
		} else if dataKonten[tengah].Judul < judul {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu {
		fmt.Println("\nKonten ditemukan:")
		fmt.Printf("Judul: %s\n", dataKonten[tengah].Judul)
		fmt.Printf("Platform: %s\n", dataKonten[tengah].Platform)
		fmt.Printf("Kategori: %s\n", dataKonten[tengah].Kategori)
		fmt.Printf("Likes: %d\n", dataKonten[tengah].Likes)
		fmt.Printf("Komentar: %d\n", dataKonten[tengah].Komentar)
		fmt.Printf("Shares: %d\n", dataKonten[tengah].Shares)
	} else {
		fmt.Println("Konten tidak ditemukan.")
	}
}

// ubah konten
func ubahKonten(dataKonten *[max]Konten, jumlah int) {
	var pilihData int
	var pilihan string
	var kiri, kanan, tengah, idx int

	idx = -1
	kiri = 0
	kanan = jumlah - 1

	tampilkanKonten(dataKonten, jumlah)
	fmt.Print("Data konten keberapa yang ingin diubah: ")
	fmt.Scan(&pilihData)
	pilihData-- 

	if pilihData >= jumlah || pilihData < 0 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2
		if pilihData == tengah {
			idx = tengah
		} else if pilihData < tengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if idx == pilihData {
		fmt.Println("\n=== DATA KONTEN KE-", idx+1, "===")
		fmt.Println("1. Judul    :", dataKonten[idx].Judul)
		fmt.Println("2. Platform :", dataKonten[idx].Platform)
		fmt.Println("3. Kategori :", dataKonten[idx].Kategori)
		fmt.Println("4. Likes    :", dataKonten[idx].Likes)
		fmt.Println("5. Komentar :", dataKonten[idx].Komentar)
		fmt.Println("6. Shares   :", dataKonten[idx].Shares)
		fmt.Println("====================================")

		fmt.Print("Pilih data yang ingin diubah (1-6): ")
		fmt.Scan(&pilihan)

		if pilihan == "1" {
			fmt.Print("Masukkan Judul baru    : ")
			fmt.Scan(&dataKonten[idx].Judul)
		} else if pilihan == "2" {
			fmt.Print("Masukkan Platform baru : ")
			fmt.Scan(&dataKonten[idx].Platform)
		} else if pilihan == "3" {
			fmt.Print("Masukkan Kategori baru : ")
			fmt.Scan(&dataKonten[idx].Kategori)
		} else if pilihan == "4" {
			fmt.Print("Masukkan Likes baru    : ")
			fmt.Scan(&dataKonten[idx].Likes)
		} else if pilihan == "5" {
			fmt.Print("Masukkan Komentar baru : ")
			fmt.Scan(&dataKonten[idx].Komentar)
		} else if pilihan == "6" {
			fmt.Print("Masukkan Shares baru   : ")
			fmt.Scan(&dataKonten[idx].Shares)
		} else {
			fmt.Println("Pilihan tidak valid.")
			return
		}

		fmt.Println("Data konten berhasil diubah!")
	}
}

func hapusKonten(dataKonten *[max]Konten, jumlah *int) {
	var idx int

	if *jumlah == 0 {
		fmt.Println("Belum ada data konten.")
		return
	}

	tampilkanKonten(dataKonten, *jumlah)
	fmt.Print("\nPilih nomor konten yang akan dihapus: ")
	fmt.Scan(&idx)
	idx--

	if idx < 0 || idx >= *jumlah {
		fmt.Println("Nomor konten tidak valid.")
		return
	}

	for i := idx; i < *jumlah-1; i++ {
		dataKonten[i] = dataKonten[i+1]
	}
	*jumlah--
	fmt.Println("Data konten berhasil dihapus!")
}
