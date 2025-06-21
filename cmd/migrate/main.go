package main

import (
	"github.com/gogaruda/auth/auth/config"
	authDB "github.com/gogaruda/auth/auth/database"
	blogDB "github.com/gogaruda/blog/blog/database"
	"log"
)

func main() {
	config.LoadENV()

	log.Println("Migrasi modul auth...")
	if err := authDB.RunMigration(); err != nil {
		log.Fatalf("Gagal migrasi auth: %v", err)
	}

	log.Println("Migrasi modul blog...")
	if err := blogDB.RunMigration(); err != nil {
		log.Fatalf("Gagal migrasi blog: %v", err)
	}

	log.Println("Semua migrasi selesai")
}
