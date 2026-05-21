package handler

import (
	"fmt"
	"log"
)

// function for update profil
func (h *CLIHandler) UpdateProfilUser(userID int, nama, hp, alamat string) {
	query := `UPDATE user_details SET full_name = ?, phone = ?, shipping_address = ? WHERE user_id = ?`
	_, err := h.DB.Exec(query, nama, hp, alamat, userID)
	if err != nil {
		log.Println("Gagal update profil:", err)
		return
	}
	fmt.Println("✅ Profil berhasil diperbarui!")
}

// function create order use SQL
func (h *CLIHandler) BuatPesanan(userID int, bagID int, qty int) {
	tx, err := h.DB.Begin()
	if err != nil {
		log.Println("Gagal memulai transaksi:", err)
		return
	}

	// Check stock from table bag
	var stock, price int
	err = tx.QueryRow(`SELECT stock, price FROM bags WHERE id = ?`, bagID).Scan(&stock, &price)
	if err != nil {
		fmt.Println("❌ Tas tidak ditemukan.")
		tx.Rollback()
		return
	}
	if stock < qty {
		fmt.Printf("❌ Stok tidak cukup! Stok tersisa: %d\n", stock)
		tx.Rollback()
		return
	}

	// Insert to table orders
	res, err := tx.Exec(`INSERT INTO orders (user_id, status) VALUES (?, 'Pending')`, userID)
	if err != nil {
		log.Println("Gagal membuat order:", err)
		tx.Rollback()
		return
	}
	orderID, _ := res.LastInsertId()

	// Insert to table order_items
	_, err = tx.Exec(`INSERT INTO order_items (order_id, bag_id, quantity) VALUES (?, ?, ?)`, orderID, bagID, qty)
	if err != nil {
		log.Println("Gagal memasukkan item pesanan:", err)
		tx.Rollback()
		return
	}

	// Reduce stock in table bags
	_, err = tx.Exec(`UPDATE bags SET stock = stock - ? WHERE id = ?`, qty, bagID)
	if err != nil {
		log.Println("Gagal memotong stok:", err)
		tx.Rollback()
		return
	}

	// Save order
	err = tx.Commit()
	if err != nil {
		log.Println("Transaksi gagal:", err)
		return
	}

	fmt.Printf("✅ Pesanan berhasil dibuat! Total tagihan: Rp%d\n", price*qty)
}

// function history order
func (h *CLIHandler) LihatRiwayat(userID int) {
	query := `
		SELECT o.id, o.status, b.name, oi.quantity, o.order_date
		FROM orders o
		JOIN order_items oi ON o.id = oi.order_id
		JOIN bags b ON oi.bag_id = b.id
		WHERE o.user_id = ?
		ORDER BY o.order_date DESC
	`
	rows, err := h.DB.Query(query, userID)
	if err != nil {
		log.Println("Gagal memuat riwayat:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n===== RIWAYAT PESANAN SAYA =====")
	for rows.Next() {
		var orderID, qty int
		var status, bagName, date string
		rows.Scan(&orderID, &status, &bagName, &qty, &date)
		fmt.Printf("[Order #%d] %s (%d pcs) - Status: %s | Tanggal: %s\n", orderID, bagName, qty, status, date)
	}
}

// function cancel order
func (h *CLIHandler) BatalkanPesananCustomer(userID int, orderID int) {
	// Query to check status order and update status order
	query := `UPDATE orders SET status = 'Dibatalkan' WHERE id = ? AND user_id = ? AND status = 'Pending'`
	res, err := h.DB.Exec(query, orderID, userID)
	if err != nil {
		log.Println("Gagal membatalkan pesanan:", err)
		return
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		fmt.Println("❌ Gagal: Pesanan tidak ditemukan, bukan milik Anda, atau sudah diproses/selesai.")
		return
	}
	fmt.Printf("✅ Pesanan #%d berhasil dibatalkan!\n", orderID)
}
