package golangdatabase

import (
	"database/sql"
	"testing"
	// ketika open connection harus panggil driver
	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	// sql.DB sebenarnya bukansebuah koneksi ke database
	// init driver dalam open connection open(driver(string), data source namenya)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// gunakan DB

	// go ada database pooling merupakan management database secara otomatis yang di lakukan oleh golang
}
