package database

import "database/sql"

func Migrations(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	return err
}
