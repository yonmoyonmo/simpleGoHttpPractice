package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetConnector() *sql.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("USER_NAME"),
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("HOST"),
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               os.Getenv("DB_NAME"),
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(os.Getenv("TEST"))
	db := GetConnector()

	err2 := db.Ping()
	if err != nil {
		log.Println(err2)
	}

	var connectionTest string
	err = db.QueryRow("SELECT sugguestion_text FROM sugguestion WHERE id =12").Scan(&connectionTest)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(connectionTest)
	print("success")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
