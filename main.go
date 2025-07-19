package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

func ConnectToDB(envFile string) (*sql.DB, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading file %s of $v", envFile, err)
	}
	fmt.Println("DB_HOST =", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER =", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD =", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME =", os.Getenv("DB_NAME"))

	port := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	fmt.Println("DSN:", dsn)

	//dsn := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	//

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Ping ошибка: %v\\n", err)
	}
	fmt.Println("✅ Успешное подключение к базе данных!")

	fmt.Printf("Listening on port %s\n", port)
	return db, nil

}

func main() {
	env := os.Getenv("ENV")
	var envFile string
	if env == "prod" {
		envFile = ".env.prod"
	} else {
		envFile = ".env.dev"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading file %s of $v", envFile, err)
	}
	currentEnv := os.Getenv("ENV")
	adminName := os.Getenv("ADMIN_NAME")
	adminRole := os.Getenv("ADMIN_ROLE")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		data := map[string]string{
			"Env":       currentEnv,
			"AdminName": adminName,
			"AdminRole": adminRole,
		}

		tmpl.Execute(w, data)
	})
	log.Println("Сервер на :8080")
	http.ListenAndServe(":8080", nil)

}
