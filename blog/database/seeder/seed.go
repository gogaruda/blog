package seeder

import (
	"github.com/gogaruda/auth/auth/config"
)

func SeedRun() error {
	db := config.ConnectDB()

	if err := Tags(db); err != nil {
		return err
	}

	if err := Category(db); err != nil {
		return err
	}

	return nil
}
