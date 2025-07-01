package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitDB inicializa la conexión a la base de datos MySQL
func InitDB() {
	user := "root"
	password := "admin"
	host := "localhost"
	port := "3306"
	dbname := "sanjorge_despacho_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error abriendo conexión a MySQL: %v", err)
	}

	// Verificar conexión
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error al conectar con MySQL: %v", err)
	}

	log.Println("✅ Conectado a MySQL correctamente")
}
