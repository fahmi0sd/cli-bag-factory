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

// Function View All Product
func (h *CLIHandler) ViewAllProducts() {
	query := `
		SELECT b.id, c.name, b.name, b.material, b.price, b.stock 
		FROM bags b
		JOIN categories c ON b.category_id = c.id
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
	// Validasi status Enum
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

// Function Generate Sales Report
func (h *CLIHandler) GenerateSalesReport() {
	query := `
		SELECT COALESCE(SUM(oi.quantity), 0), COALESCE(SUM(oi.quantity * b.price), 0)
		FROM order_items oi
		JOIN orders o ON oi.order_id = o.id
		JOIN bags b ON oi.bag_id = b.id
		WHERE o.status = 'Selesai'
	`
	var totalItem, totalPendapatan int
	err := h.DB.QueryRow(query).Scan(&totalItem, &totalPendapatan)
	if err != nil {
		log.Println("Gagal memuat laporan penjualan:", err)
		return
	}

	fmt.Println("\n===== LAPORAN PENJUALAN =====")
	fmt.Printf("Total Tas Terjual : %d item\n", totalItem)
	fmt.Printf("Total Pendapatan  : Rp%d\n", totalPendapatan)
}

// Function Generate User Report
func (h *CLIHandler) GenerateUserReport() {
	query := `SELECT COUNT(*) FROM users WHERE role = 'Customer'`
	var totalCustomer int
	err := h.DB.QueryRow(query).Scan(&totalCustomer)
	if err != nil {
		log.Println("Gagal memuat laporan user:", err)
		return
	}

	fmt.Println("\n===== LAPORAN USER =====")
	fmt.Printf("Total Customer Terdaftar: %d akun\n", totalCustomer)
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
