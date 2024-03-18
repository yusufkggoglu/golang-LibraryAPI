package repository

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

var Conn *sql.DB
var err error
var Dbs *Queries
var Store *session.Store

func Connect() {
	dbinfo := "host=go_library_db port=5432 user=admin password=admin123 dbname=postgres sslmode=disable"

	Conn, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalln("bağlanılamadı : ", err)
	}
	if err := Conn.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Bağlantı oluşturuldu.")
	}
	createAdminSql := `
	INSERT INTO users (username, password, role)
	VALUES ('admin', 'admin123', 'Admin')
	ON CONFLICT (username) DO NOTHING;`

	_, err = Conn.Exec(createAdminSql)
	if err != nil {
		log.Fatal("Tablo oluşturma hatası:", err)
	}

	Dbs = New(Conn)
	Store = session.New()
}
