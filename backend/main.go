package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 连接字符串
	connStr := "host=localhost port=5432 user=vivi password=vivi#2024 dbname=Litlog sslmode=disable search_path=public"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Printf("Successfully connected to the database!!")

	// 初始化路由
	router := InitializeRouter(db) // 使用 db

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
