package handler

import (
	"testing"

	"github.com/fahmi0sd/cli-bag-factory/helper"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db := helper.SetupTestDB()
	defer db.Close()

	h := NewCLIHandler(db)

	testEmail := "admin@gmail.com"
	testPassword := "admin123"

	t.Run("Skenario 1: Login Berhasil", func(t *testing.T) {
		user, err := h.Login(testEmail, testPassword)

		assert.NoError(t, err, "Seharusnya tidak ada error saat login")
		assert.NotNil(t, user, "Data user tidak boleh nil")

		if user != nil {
			assert.Equal(t, testEmail, user.Email, "Email yang dikembalikan harus sama")
			assert.Equal(t, "Admin", user.Role, "Role harus Admin")
		}
	})

	t.Run("Skenario 2: Login Gagal (Password Salah)", func(t *testing.T) {
		user, err := h.Login(testEmail, "passwordsalah1234")

		assert.Error(t, err, "Seharusnya menghasilkan error")
		assert.Nil(t, user, "Data user harus nil jika gagal")
		assert.Equal(t, "email atau password salah", err.Error(), "Pesan error harus sesuai")
	})
}
