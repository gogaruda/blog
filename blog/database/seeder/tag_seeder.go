package seeder

import (
	"database/sql"
	"github.com/gogaruda/pkg/utils"
)

func Tags(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO tags(id, name, slug, description, seo_title, seo_description) 
                VALUES(?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`,
		utils.NewULID(), "go", "go-lang", "belajar go lang", "belajar go lang asyik", "belajar go lang",
		utils.NewULID(), "node", "node-js", "belajar node - js", "seru banget belajar node js", "seru banget belajar node js",
		utils.NewULID(), "react js", "react-js", "yuk react-js", "Inilah belajar react js", "Inilah belajar react js",
		utils.NewULID(), "next js", "next-js", "belajar next js seru", "belajar next js seseru ini", "belajar next js seseru ini")
	if err != nil {
		return err
	}
	return nil
}
