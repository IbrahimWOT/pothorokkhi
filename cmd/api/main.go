package main

import (
	"fmt"
	"net/http"
	"pothorokkhi/internal/db" // মডিউলের নাম অনুযায়ী পাথ ঠিক করা হয়েছে

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🛡️ PothRokkhi (পথরক্ষী) Backend Engine Booting Up...")

	// ডাটাবেজ ইনিশিয়েলাইজেশন (ডকার অন থাকতে হবে)
	db.InitDatabases()

	// Gin রাউটার ডিফাইন করা
	r := gin.Default()

	// স্ট্যাটিক ফাইল পাথ লিঙ্কিং (HTML/CSS লোড করার জন্য)
	r.StaticFS("/static", http.Dir("./web"))

	// মেইন গেটওয়ে বা ল্যান্ডিং পেজ হ্যান্ডলার
	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	fmt.Println("🚀 PothRokkhi Core Core Live on http://localhost:8080")
	r.Run(":8080")
}
