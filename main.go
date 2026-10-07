package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// PENTING: Menghubungkan folder static agar CSS & JS terbaca
	r.Static("/static", "./static")

	// Meload semua file HTML di folder templates
	r.LoadHTMLGlob("templates/*")

	// Route Navigasi
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "QURBAKALA"})
	})

	r.GET("/detail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "detail.html", nil)
	})

	r.GET("/checkout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "checkout.html", nil)
	})

	println("Akses Qurbakala di: http://localhost:8080")
	r.Run(":8080")
}