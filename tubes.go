package main

import "fmt"
type Konten struct {
    Judul    string
    Platform string
    Kategori string
    Likes    int
    Komentar int
    Shares   int
}

var tabKonten [max]Konten
var totalKonten int = 0
const max int = 50

func main() {

	var pilihMenu int

	for pilihMenu != 7 {
		menu(&pilihMenu) 

		if pilihMenu == 1 {
			tambahKonten()
		} else if pilihMenu == 2 {
			tampilkanKonten()
		} else if pilihMenu == 3 {
			fmt.Println("Fitur cari konten biasa belum tersedia.")
		} else if pilihMenu == 4 {
			fmt.Println("Fitur cari konten cepat belum tersedia.")
		} else if pilihMenu == 5 {
			fmt.Println("Fitur urutkan berdasarkan likes belum tersedia.")
		} else if pilihMenu == 6 {
			fmt.Println("Fitur urutkan berdasarkan popularitas belum tersedia.")
		} else if pilihMenu == 7 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}


func menu(pilihMenu *int){
    fmt.Println("\n=== APLIKASI MANAJEMEN KONTEN ===")
    fmt.Println("1. Tambah Konten Baru")
    fmt.Println("2. Lihat Semua Konten")
    fmt.Println("3. Cari Konten (Pencarian Biasa)")
    fmt.Println("4. Cari Konten (Pencarian Cepat)")
    fmt.Println("5. Urutkan Berdasarkan Likes")
    fmt.Println("6. Urutkan Berdasarkan Popularitas")
    fmt.Println("7. Keluar")
    fmt.Print("Pilih menu: ")
    fmt.Scan(pilihMenu)
}

func tambahKonten() {
	if totalKonten >= max {
		fmt.Println("Kapasitas penyimpanan konten sudah penuh.")
		return
	}
    var tambahkonten Konten

	fmt.Println("\n=== TAMBAH KONTEN BARU ===")
	fmt.Print("Judul     : ")
	fmt.Scan(&tambahkonten.Judul)
	fmt.Print("Platform  : ")
	fmt.Scan(&tambahkonten.Platform)
	fmt.Print("Kategori  : ")
	fmt.Scan(&tambahkonten.Kategori)
	fmt.Print("Likes     : ")
	fmt.Scan(&tambahkonten.Likes)
	fmt.Print("Komentar  : ")
	fmt.Scan(&tambahkonten.Komentar)
	fmt.Print("Shares    : ")
	fmt.Scan(&tambahkonten.Shares)

	tabKonten[totalKonten] = tambahkonten
	totalKonten++
	fmt.Println("Konten berhasil ditambahkan.")
}

func tampilkanKonten() {
	var i int
	if totalKonten == 0 {
		fmt.Println("Belum ada data konten.")
		return
	}

	fmt.Println("\n===== DAFTAR KONTEN =====")
	for i = 0; i < totalKonten; i++ {
		fmt.Printf("\n%d. Judul: %s\n", i+1, tabKonten[i].Judul)
		fmt.Printf("   Platform: %s\n", tabKonten[i].Platform)
		fmt.Printf("   Kategori: %s\n", tabKonten[i].Kategori)
		fmt.Printf("   Likes: %d\n", tabKonten[i].Likes)
		fmt.Printf("   Komentar: %d\n", tabKonten[i].Komentar)
		fmt.Printf("   Shares: %d\n", tabKonten[i].Shares)
		fmt.Println("------------------------------")
	}
}


func binarySearch(tabKonten []Konten, n int, targetLikes int) int {
	var tengah, kiri, kanan, ketemu int
	ketemu = -1
	kiri = 0
	kanan = n - 1
	tengah = (kiri + kanan) / 2

	for kiri <= kanan && ketemu == -1 {
		if targetLikes < tabKonten[tengah].Likes {
			kanan = tengah - 1
		} else if targetLikes > tabKonten[tengah].Likes {
			kiri = tengah + 1
		} else {
			ketemu = tengah
		}
		tengah = (kiri + kanan) / 2
	}

	return ketemu
}

