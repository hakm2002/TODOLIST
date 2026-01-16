package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hakm2002/TODOLIST/config"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load .env file
	// Ini harus dipanggil paling pertama agar variabel seperti DB_USER, dll terbaca
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: File .env tidak ditemukan, akan menggunakan environment system jika ada.")
	}

	// 2. Inisialisasi Database
	// Fungsi ini akan membaca variabel dari .env dan melakukan koneksi ke MySQL
	config.InitDB()

	// 3. Setup Router (Gin)
	r := gin.Default()

	// --- Area Route Sementara (Hanya untuk Tes) ---
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Selamat! Server ToDo List Berjalan Lancar 🚀",
			"status":  "connected",
		})
	})
	// ----------------------------------------------

	// 4. Jalankan Server di Port 8080
	log.Println("Server sedang berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
