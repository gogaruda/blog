package main

import (
	"github.com/gogaruda/auth/auth/config"
	authSeed "github.com/gogaruda/auth/auth/database/seeder"
	"log"
)

func main() {
	config.LoadENV()
	config.ConnectDB()

	log.Println("Seeder modul auth...")
	if err := authSeed.SeedRun(); err != nil {
		log.Fatalf("Gagal seeder auth: %v", err)
	}

	log.Println("Semua seeder selesai")
}
