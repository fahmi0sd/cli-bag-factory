package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func InterfaceCLI() {
	for {
		fmt.Println("===========================================")
		fmt.Println("     SELAMAT DATANG DI PABRIK TAS CLI      ")
		fmt.Println("===========================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Print("Pilih menu (1-3): ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// this is code for login

func login() {
	var username, password string

	fmt.Print("Username : ")
	fmt.Scanln(&username)

	fmt.Print("Password : ")
	fmt.Scanln(&password)

	// Login sederhana
	if username == "admin" && password == "admin" {
		fmt.Println("Login Admin Berhasil!")
		menuAdmin()
	} else {
		fmt.Println("Login Customer Berhasil!")
		menuCustomer()
	}
}

// register

func register() {
	var username, password string

	fmt.Println("===== REGISTER CUSTOMER =====")
	fmt.Print("Masukkan Username : ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)

	fmt.Println("Akun berhasil dibuat!")
}

// menu customer

func menuCustomer() {
	for {
		fmt.Println("\n--- Menu Customer ---")
		fmt.Println("1. Lengkapi/Update Profil & Alamat")
		fmt.Println("2. Lihat Katalog Tas")
		fmt.Println("3. Buat Pesanan Baru")
		fmt.Println("4. Lihat Riwayat Pesanan Saya")
		fmt.Println("5. Batalkan Pesanan")
		fmt.Println("6. Logout")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			updateProfil()
		case 2:
			lihatKatalog()
		case 3:
			buatPesanan()
		case 4:
			riwayatPesanan()
		case 5:
			batalkanPesanan()
		case 6:
			fmt.Println("Logout berhasil!")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// menu admin

func menuAdmin() {
	for {
		fmt.Println("\n--- Menu Admin Pabrik ---")
		fmt.Println("1. Kelola Produk Tas (Menambah, Mengupdate, Menghapus Produk)")
		fmt.Println("2. Proses Pesanan Customer (Update Status)")
		fmt.Println("3. Menu Laporan (Reports)")
		fmt.Println("4. Logout")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			kelolaProduk()
		case 2:
			prosesPesanan()
		case 3:
			menuLaporan()
		case 4:
			fmt.Println("Logout berhasil!")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// code for feature customer

func updateProfil() {
	reader.ReadString('\n')

	fmt.Println("\n===== UPDATE PROFIL =====")

	fmt.Print("Nama Lengkap : ")
	nama, _ := reader.ReadString('\n')

	fmt.Print("Alamat : ")
	alamat, _ := reader.ReadString('\n')

	fmt.Print("No HP : ")
	nohp, _ := reader.ReadString('\n')

	fmt.Println("\nProfil berhasil diperbarui!")
	fmt.Println("Nama   :", strings.TrimSpace(nama))
	fmt.Println("Alamat :", strings.TrimSpace(alamat))
	fmt.Println("No HP  :", strings.TrimSpace(nohp))
}

func lihatKatalog() {
	fmt.Println("\n===== KATALOG TAS =====")
	fmt.Println("1. Tas Ransel - Rp150000")
	fmt.Println("2. Tas Selempang - Rp100000")
	fmt.Println("3. Tas Laptop - Rp250000")
}

func buatPesanan() {
	var pilih int
	var jumlah int

	lihatKatalog()

	fmt.Print("\nPilih produk : ")
	fmt.Scanln(&pilih)

	fmt.Print("Jumlah beli : ")
	fmt.Scanln(&jumlah)

	fmt.Println("Pesanan berhasil dibuat!")
}

func riwayatPesanan() {
	fmt.Println("\n===== RIWAYAT PESANAN =====")
	fmt.Println("1. Tas Ransel - Diproses")
	fmt.Println("2. Tas Laptop - Selesai")
}

func batalkanPesanan() {
	var id int

	fmt.Println("\n===== BATALKAN PESANAN =====")
	fmt.Print("Masukkan ID Pesanan : ")
	fmt.Scanln(&id)

	fmt.Println("Pesanan berhasil dibatalkan!")
}

// this is code for admin produk

func kelolaProduk() {
	for {
		fmt.Println("\n===== KELOLA PRODUK =====")
		fmt.Println("1. Tambah Produk")
		fmt.Println("2. Lihat Produk")
		fmt.Println("3. Update Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Tambah Produk")
		case 2:
			lihatKatalog()
		case 3:
			fmt.Println("Update Produk")
		case 4:
			fmt.Println("Hapus Produk")
		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

func prosesPesanan() {
	fmt.Println("\n===== PROSES PESANAN =====")
	fmt.Println("1. Pesanan #001 - Diproses")
	fmt.Println("2. Pesanan #002 - Menunggu")

	var id int
	fmt.Print("Pilih ID Pesanan : ")
	fmt.Scanln(&id)

	fmt.Println("Status pesanan berhasil diperbarui!")
}

func menuLaporan() {
	fmt.Println("\n===== MENU LAPORAN =====")
	fmt.Println("Total Pesanan : 10")
	fmt.Println("Pesanan Selesai : 7")
	fmt.Println("Pesanan Dibatalkan : 1")
}
