package sqlcservice

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

var DB *sql.DB

func Connect() error {

	sqlPass := viper.Get("MYSQL_PASS")
	sqlHost := viper.Get("MYSQL_HOST")
	sqlUser := viper.Get("MYSQL_USER")

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:3306)/links_db?parseTime=True", sqlUser, sqlPass, sqlHost)
	db, _ := sql.Open("mysql", dataSourceName)

	err := db.Ping()

	if err != nil {
		return err
	}

	fmt.Println("MYSQL connected")

	DB = db
	return nil
}
