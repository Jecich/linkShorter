package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

const (
	DBUser     = "shorter"        // Пользователь MySQL
	DBPassword = "password"       // Пароль
	DBHost     = "localhost:3306" // Хост и порт
	DBName     = "shorter"        // Имя базы данных
)

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", DBUser, DBPassword, DBHost, DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Err conection to MySQL: %v", err)
	}

	// Проверяем соединение
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Err ping MySQL: %v", err)
	}

	log.Println("Successful connection to MySQL!")
}
