package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ballinboys/myapp/database"
	"github.com/ballinboys/myapp/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	log.Println("Berhasil koneksi ke database")

	// Inisialisasi Gin router
	r := gin.Default()

	// Buat GraphQL server
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: database.DB}}),
	)

	r.POST("/query", gin.WrapH(srv))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))

	// Jalankan server dalam goroutine
	go func() {
		log.Println("Server berjalan di http://localhost:8080/")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Gagal menjalankan server: %v", err)
		}
	}()

	// Tunggu sinyal shutdown (Ctrl+C, docker stop, dll)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Tutup koneksi database
	if err := database.DB.Close(); err != nil {
		log.Printf("Gagal menutup koneksi database: %v", err)
	} else {
		log.Println("Koneksi database ditutup.")
	}

	// Tambahkan delay kecil untuk memastikan semua log terselesaikan
	time.Sleep(1 * time.Second)
}
