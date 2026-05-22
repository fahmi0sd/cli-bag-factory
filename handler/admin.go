package handler

import (
	"fmt"
	"log"
)

// Use Case: Admin - manage product, process order, generate reports
// Function Add Product
func (h *CLIHandler) AddProduct(categoryID int, name, material string, price, stock int) {
	query := `INSERT INTO bags (category_id, name, material, price, stock) VALUES (?, ?, ?, ?, ?)`
	_, err := h.DB.Exec(query, categoryID, name, material, price, stock)
	if err != nil {
		log.Println("Gagal menambah produk:", err)
		return
	}
	fmt.Println("✅ Produk berhasil ditambahkan ke database!")
}

// Function Update Product
func (h *CLIHandler) UpdateProduk(id int, name, material string, price, stock int) {
	query := `UPDATE bags SET name = ?, material = ?, price = ?, stock = ? WHERE id = ?`
	res, err := h.DB.Exec(query, name, material, price, stock, id)
	if err != nil {
		log.Println("Gagal mengupdate produk:", err)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		fmt.Println("⚠️ Produk dengan ID tersebut tidak ditemukan.")
		return
	}
	fmt.Printf("✅ Produk #%d berhasil diperbarui!\n", id)
}

// Function Delete Product
func (h *CLIHandler) HapusProduk(id int) {
	query := `DELETE FROM bags WHERE id = ?`
	res, err := h.DB.Exec(query, id)
	if err != nil {
		fmt.Println("❌ Gagal menghapus produk. Pastikan produk ini belum pernah masuk ke transaksi pesanan.")
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		fmt.Println("⚠️ Produk dengan ID tersebut tidak ditemukan.")
		return
	}
	fmt.Printf("✅ Produk #%d berhasil dihapus dari sistem!\n", id)
}

// Function view all categories
func (h *CLIHandler) ViewCategories() {
	query := `SELECT id, name FROM categories ORDER BY id ASC`
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println("Gagal mengambil daftar kategori:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- Daftar Kategori ---")
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err == nil {
			fmt.Printf("[%d] %s\n", id, name)
		}
	}
	fmt.Println("-----------------------")
}

// Function View All Product
func (h *CLIHandler) ViewAllProducts() {
	query := `
		SELECT b.id, c.name, b.name, b.material, b.price, b.stock 
		FROM bags b
		JOIN categories c ON b.category_id = c.id
		ORDER BY b.id ASC
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println("Gagal mengambil data produk:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n===== KATALOG TAS (DATABASE) =====")
	for rows.Next() {
		var id, price, stock int
		var category, name, material string
		err := rows.Scan(&id, &category, &name, &material, &price, &stock)
		if err != nil {
			continue
		}
		fmt.Printf("[%d] %s - %s (%s) | Rp%d | Stok: %d\n", id, category, name, material, price, stock)
	}
}

// Function Update Status Order
func (h *CLIHandler) UpdateOrderStatus(orderID int, status string) {
	if status != "Diproses" && status != "Selesai" && status != "Dibatalkan" {
		fmt.Println("❌ Status tidak valid! Gunakan: Diproses, Selesai, atau Dibatalkan.")
		return
	}

	query := `UPDATE orders SET status = ? WHERE id = ?`
	res, err := h.DB.Exec(query, status, orderID)
	if err != nil {
		log.Println("Gagal update status pesanan:", err)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		fmt.Println("⚠️ Pesanan dengan ID tersebut tidak ditemukan.")
		return
	}
	fmt.Printf("✅ Status pesanan #%d berhasil diubah menjadi '%s'!\n", orderID, status)
}

// Function View All Orders
func (h *CLIHandler) ViewAllOrders() {
	query := `
		SELECT o.id, u.email, o.order_date, o.status
		FROM orders o
		JOIN users u ON o.user_id = u.id
		ORDER BY o.order_date DESC
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println("Gagal mengambil data pesanan:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n===== DAFTAR PESANAN =====")
	for rows.Next() {
		var id int
		var email, date, status string
		if err := rows.Scan(&id, &email, &date, &status); err == nil {
			fmt.Printf("[ID: %d] Pembeli: %s | Tgl: %s | Status: %s\n", id, email, date, status)
		}
	}
	fmt.Println("--------------------------")
}

// Function Generate Sales Report
func (h *CLIHandler) GenerateSalesReport() {
	querySummary := `
		SELECT 
			COUNT(DISTINCT o.id), 
			COALESCE(SUM(oi.quantity), 0), 
			COALESCE(SUM(oi.quantity * b.price), 0)
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		LEFT JOIN bags b ON oi.bag_id = b.id
		WHERE o.status = 'Selesai'
	`
	var totalOrder, totalItem, totalPendapatan int
	err := h.DB.QueryRow(querySummary).Scan(&totalOrder, &totalItem, &totalPendapatan)
	if err != nil {
		log.Println("Gagal memuat ringkasan penjualan:", err)
		return
	}

	fmt.Println("\n===== LAPORAN PENJUALAN =====")
	fmt.Printf("Total Transaksi Selesai : %d transaksi\n", totalOrder)
	fmt.Printf("Total Tas Terjual       : %d item\n", totalItem)
	fmt.Printf("Total Pendapatan        : Rp%d\n", totalPendapatan)

	fmt.Println("\n--- Rincian Tas Terjual ---")
	queryDetail := `
		SELECT b.name, SUM(oi.quantity), SUM(oi.quantity * b.price)
		FROM order_items oi
		JOIN orders o ON oi.order_id = o.id
		JOIN bags b ON oi.bag_id = b.id
		WHERE o.status = 'Selesai'
		GROUP BY b.id, b.name
		ORDER BY SUM(oi.quantity) DESC
	`
	rows, err := h.DB.Query(queryDetail)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var name string
			var qty, subtotal int
			if err := rows.Scan(&name, &qty, &subtotal); err == nil {
				fmt.Printf("- %s : %d pcs (Menghasilkan: Rp%d)\n", name, qty, subtotal)
			}
		}
	}
	fmt.Println("=============================")
}

// Function Generate User Report
func (h *CLIHandler) GenerateUserReport() {
	var totalCustomer int
	h.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE role = 'Customer'`).Scan(&totalCustomer)

	fmt.Println("\n===== LAPORAN USER AKTIF =====")
	fmt.Printf("Total Customer Terdaftar: %d akun\n", totalCustomer)

	fmt.Println("\n--- Rincian Aktivitas Pembeli ---")
	query := `
		SELECT u.email, COALESCE(ud.full_name, 'Belum Update Profil'), COUNT(o.id)
		FROM users u
		LEFT JOIN user_details ud ON u.id = ud.user_id
		LEFT JOIN orders o ON u.id = o.user_id
		WHERE u.role = 'Customer'
		GROUP BY u.id, u.email, ud.full_name
		ORDER BY COUNT(o.id) DESC
	`
	rows, err := h.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var email, name string
			var totalOrder int
			if err := rows.Scan(&email, &name, &totalOrder); err == nil {
				fmt.Printf("- %s (%s) | Total Melakukan Pesanan: %d kali\n", name, email, totalOrder)
			}
		}
	}
	fmt.Println("==============================")
}

// Function Generate Stock Report
func (h *CLIHandler) GenerateStockReport() {
	query := `
		SELECT c.name, COALESCE(SUM(b.stock), 0)
		FROM categories c
		LEFT JOIN bags b ON c.id = b.category_id
		GROUP BY c.id, c.name
	`
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println("Gagal memuat laporan stok:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n===== LAPORAN STOK TAS =====")
	var totalSemuaStok int
	for rows.Next() {
		var categoryName string
		var totalStock int
		if err := rows.Scan(&categoryName, &totalStock); err == nil {
			fmt.Printf("- Kategori %s : %d pcs\n", categoryName, totalStock)
			totalSemuaStok += totalStock
		}
	}
	fmt.Println("----------------------------")
	fmt.Printf("Total Seluruh Stok Gudang: %d pcs\n", totalSemuaStok)
}
