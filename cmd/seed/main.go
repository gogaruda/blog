package main

import (
	"github.com/gogaruda/auth/auth/config"
	authSeed "github.com/gogaruda/auth/auth/database/seeder"
	blogSeed "github.com/gogaruda/blog/blog/database/seeder"
	"log"
)

func main() {
	config.LoadENV()
	config.ConnectDB()

	log.Println("Seeder modul auth...")
	if err := authSeed.SeedRun(); err != nil {
		log.Fatalf("Gagal seeder auth: %v", err)
	}

	log.Println("Seeder modul blog...")
	if err := blogSeed.SeedRun(); err != nil {
		log.Fatalf("Gagal seeder blog: %v", err)
	}

	log.Println("Semua seeder selesai")
}
