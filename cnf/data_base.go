package cnf

import (
	"database/sql"
	"fmt"
)

func InitDB() (*sql.DB, error) {

	config := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DataBase.User, config.DataBase.Pass, config.DataBase.Host, config.DataBase.Port, config.DataBase.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Test connection to the database
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging the database: %w", err)
	}

	return db, nil
}
