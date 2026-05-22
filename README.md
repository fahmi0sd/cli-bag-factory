# Pabrik Tas CLI E-Commerce

Aplikasi Command Line Interface (CLI) E-Commerce untuk manajemen operasional Pabrik Tas. Dibangun menggunakan Golang dan MySQL.

## ✨ Fitur Utama

**Admin Flow**
* Kelola Produk Tas (CRUD).
* Proses Pesanan Customer (Melihat riwayat, update status: Diproses, Selesai, Dibatalkan).
* Laporan Komprehensif: Laporan Penjualan (Detail tas terjual), Laporan Stok Gudang, dan Laporan Aktivitas User Aktif.

**Customer Flow**
* Autentikasi User (Login dan Register).
* Lengkapi Profil & Alamat Pengiriman.
* Lihat Katalog Tas.
* Buat Pesanan.
* Riwayat Pesanan & Pembatalan Pesanan.

## 🛠️ Tech Stack
* **Language:** Golang (Go)
* **Database:** MySQL
* **Driver & Config:** `github.com/go-sql-driver/mysql`, `github.com/joho/godotenv`
* **Testing:** `github.com/stretchr/testify`

## 📂 Struktur Direktori
```text
.
├── cli/          # Antarmuka terminal (UI, Scanln, Reader)
├── db/           # Inisialisasi koneksi database MySQL
├── entity/       # Struct Golang sebagai representasi tabel database
├── handler/      # Business logic, query database, dan file Unit Test
├── helper/       # Fungsi bantuan (misal: inisialisasi DB khusus testing)
├── sql/          # Skema tabel (database.sql) 
├── .env          # Kredensial database lokal (Diabaikan oleh Git)
├── go.mod        # Modul dan manajemen dependensi Go
└── main.go       # Entry point aplikasi
```

## Cara Menjalankan Aplikasi
### 1. Persiapan Database
* Buat database MYQSL baru di lokal 
* Eksekusi file sql/database.sql 
### 2. Konfigurasi Environment
* Buat file bernama .env di root direktori proyek dan isi dengan konfigurasi Data Source Name (DSN) database Anda:
```MYSQL_DSN=username:password@tcp(127.0.0.1:3306)/bag_factory_db```
* Gantilah `username`, `password`, dan `bag_factory_db` sesuai dengan konfigurasi database MySQL Anda.
### 3. Install Dependensi & Jalankan
* Buka terminal Anda, lalu eksekusi perintah berikut:
```bash
# Mengunduh semua package yang terdaftar di go.mod
go mod tidy

# Menjalankan aplikasi CLI
go run main.go
```
## Testing
* Untuk menjalankan unit test, gunakan perintah berikut di terminal:
```bash
# Masuk ke direktori handler
cd handler

# Jalankan tes dan tampilkan log detail
go test -v
```