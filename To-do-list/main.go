package main

import (
	"github.com/gin-gonic/gin"
	
	// PERHATIKAN BARIS INI: Sesuaikan dengan nama module baru di langkah 1
	"github.com/hakm2002/TODOLIST/config" 
	"github.com/hakm2002/TODOLIST/routes"
)

func main() {
	// Init DB (Ingat untuk update file config/db.go agar pakai os.Getenv seperti panduan sebelumnya)
	config.InitDB()
	
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	routes.InitRoutes(r)
	
	// PENTING: Gunakan port 8080 agar sesuai dengan Jenkinsfile & Dockerfile
	r.Run(":8080") 
}
