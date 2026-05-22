package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fahmi0sd/cli-bag-factory/handler"
)

var reader = bufio.NewReader(os.Stdin)

func InterfaceCLI(h *handler.CLIHandler) {
	for {
		fmt.Println("\n===========================================")
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
			login(h)
		case 2:
			register(h)
		case 3:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// this is code for login

func login(h *handler.CLIHandler) {
	fmt.Println("\n===== LOGIN =====")
	fmt.Print("Email    : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Password : ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Call function login in hendler
	user, err := h.Login(email, password)
	if err != nil {
		fmt.Println("❌ Gagal Login:", err.Error())
		return
	}

	fmt.Printf("✅ Login Berhasil! Selamat datang, %s.\n", user.Email)

	// cek role user
	if user.Role == "Admin" {
		menuAdmin(h)
	} else {
		menuCustomer(h, user.ID)
	}
}

// register

func register(h *handler.CLIHandler) {
	fmt.Println("\n===== REGISTER CUSTOMER BARU =====")

	fmt.Print("Masukkan Email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Masukkan Password : ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Print("Nama Lengkap : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("No HP : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Alamat Pengiriman : ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	// Call function Register from handler
	err := h.RegisterCustomer(email, password, nama, phone, alamat)
	if err != nil {
		fmt.Println("❌ Gagal mendaftar:", err)
		return
	}

	fmt.Println("✅ Akun berhasil dibuat! Silakan pilih menu Login.")
}

// menu customer

func menuCustomer(h *handler.CLIHandler, userID int) {
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
			updateProfil(h, userID)
		case 2:
			h.ViewAllProducts()
		case 3:
			buatPesanan(h, userID)
		case 4:
			h.LihatRiwayat(userID)
		case 5:
			batalkanPesanan(h, userID)
		case 6:
			fmt.Println("Logout berhasil!")
			return
		}
	}
}

// menu admin

func menuAdmin(h *handler.CLIHandler) {
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
			kelolaProduk(h)
		case 2:
			prosesPesanan(h)
		case 3:
			menuLaporan(h)
		case 4:
			fmt.Println("Logout berhasil!")
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

// code for feature customer

func updateProfil(h *handler.CLIHandler, userID int) {
	fmt.Println("\n===== UPDATE PROFIL =====")

	fmt.Print("Nama Lengkap : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("No HP : ")
	nohp, _ := reader.ReadString('\n')
	nohp = strings.TrimSpace(nohp)

	fmt.Print("Alamat : ")
	alamat, _ := reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat)

	h.UpdateProfilUser(userID, nama, nohp, alamat)
}

func buatPesanan(h *handler.CLIHandler, userID int) {
	h.ViewAllProducts()
	var bagID, jumlah int
	fmt.Print("\nMasukkan ID Produk yang ingin dibeli: ")
	fmt.Scanln(&bagID)

	fmt.Print("Jumlah beli: ")
	fmt.Scanln(&jumlah)

	h.BuatPesanan(userID, bagID, jumlah)
}

func batalkanPesanan(h *handler.CLIHandler, userID int) {
	fmt.Println("\n===== BATALKAN PESANAN =====")
	h.LihatRiwayat(userID)

	var id int
	fmt.Print("Masukkan ID Pesanan yang ingin dibatalkan: ")
	fmt.Scanln(&id)

	h.BatalkanPesananCustomer(userID, id)
}

// this is code for admin produk

func kelolaProduk(h *handler.CLIHandler) {
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
			fmt.Println("\n--- Tambah Produk Baru ---")
			var categoryID, price, stock int

			h.ViewCategories()

			fmt.Print("ID Kategori (1: Daypack, 2: Carrier, 3: Sling Bag) : ")
			fmt.Scanln(&categoryID)

			fmt.Print("Nama Tas : ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Material : ")
			material, _ := reader.ReadString('\n')
			material = strings.TrimSpace(material)

			fmt.Print("Harga (Rp) : ")
			fmt.Scanln(&price)

			fmt.Print("Stok Awal : ")
			fmt.Scanln(&stock)

			h.AddProduct(categoryID, name, material, price, stock)

		case 2:
			h.ViewAllProducts()

		case 3:
			h.ViewAllProducts()
			fmt.Println("\n--- Update Produk ---")
			var id, price, stock int

			fmt.Print("Masukkan ID Produk yang akan diupdate: ")
			fmt.Scanln(&id)

			fmt.Print("Nama Tas Baru : ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Material Baru : ")
			material, _ := reader.ReadString('\n')
			material = strings.TrimSpace(material)

			fmt.Print("Harga Baru (Rp) : ")
			fmt.Scanln(&price)

			fmt.Print("Stok Baru : ")
			fmt.Scanln(&stock)

			h.UpdateProduk(id, name, material, price, stock)

		case 4:
			h.ViewAllProducts()
			fmt.Println("\n--- Hapus Produk ---")
			var id int
			fmt.Print("Masukkan ID Produk yang akan dihapus: ")
			fmt.Scanln(&id)

			h.HapusProduk(id)

		case 5:
			return
		default:
			fmt.Println("Menu tidak tersedia!")
		}
	}
}

func prosesPesanan(h *handler.CLIHandler) {
	fmt.Println("\n===== PROSES PESANAN =====")
	h.ViewAllOrders()

	var id int
	var status string

	fmt.Print("Masukkan ID Pesanan : ")
	fmt.Scanln(&id)

	fmt.Print("Masukkan Status Baru (Diproses/Selesai/Dibatalkan): ")
	fmt.Scanln(&status)

	h.UpdateOrderStatus(id, status)
}

func menuLaporan(h *handler.CLIHandler) {
	h.GenerateSalesReport()
	h.GenerateStockReport()
	h.GenerateUserReport()
}
