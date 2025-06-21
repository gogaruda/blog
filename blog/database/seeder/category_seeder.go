package seeder

import (
	"database/sql"
	helper "github.com/gogaruda/pkg/database"
	"github.com/gogaruda/pkg/utils"
)

func Category(db *sql.DB) error {
	return helper.WithTx(db, func(tx *sql.Tx) error {
		idUtama := utils.NewULID()
		_, err := tx.Exec(`INSERT INTO 
                    categories(id, name, slug, description, seo_title, seo_description)
                    VALUES(?, ?, ?, ?, ?, ?)`,
			idUtama,
			"Pemrograman", "pemrograman", "Belajar Bahasa Pemrograman Asyik",
			"Pemrograman", "Belajar bahasa pemrograman asyik banget",
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(`INSERT INTO
                        categories(id, name, slug, description, parent_id, seo_title, seo_description)
                        VALUES(?, ?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?, ?)`,
			utils.NewULID(), "Go Lang", "go-lang", "Belajar Bahasa Pemrograman Go", idUtama, "Belajar Bahasa Pemrograman Go", "Belajar Bahasa Pemrograman Go",
			utils.NewULID(), "React JS", "react-js", "Belajar React JS Seru Banget", idUtama, "Belajar React JS Seru Banget", "Belajar React JS Seru Banget",
			utils.NewULID(), "Next JS", "next-js", "Belajar Next JS Seru Banget", idUtama, "Belajar Next JS Seru Banget", "Belajar Next JS Seru Banget",
		)

		if err != nil {
			return err
		}

		return nil
	})
}
