package main

import (
	"FirstProject/generated/repository"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type DbParams struct {
	username string
	password string
	host     string
	port     string
	database string
}

func handler(w http.ResponseWriter, r *http.Request, ctx context.Context, repo *repository.Queries) {
	if len(r.URL.Path) <= 3 {
		fmt.Fprintf(w, "No parametrs given")
		return
	}

	repo.InsertUser(ctx, repository.InsertUserParams{
		Name:     r.PathValue("name"),
		Email:    r.PathValue("email"),
		Password: r.PathValue("password"),
	})
	users, _ := repo.GetAllUsers(ctx)
	usersJson, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

func withDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func getDB(dbParams DbParams) *sql.DB {

	dsn := fmt.Sprintf("%s@(%s:%s)/%s?parseTime=true",
		dbParams.username,
		dbParams.host,
		dbParams.port,
		dbParams.database,
	)
	fmt.Printf("dsn: %s\n", dsn)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	return conn
}

func getDefaultDbParams() DbParams {
	database_flag := flag.String("db", "", "Database name")

	flag.Parse()
	database := *database_flag
	if database == "" {
		database = os.Getenv("DB_DATABASE")
	}

	if database == "" {
		log.Fatal("db flag or DB_DATABASE env is not set")
		panic("Database name is empty")
	}
	dbParams := DbParams{
		username: withDefault(os.Getenv("DB_USER"), "root"),
		password: withDefault(os.Getenv("DB_PASS"), ""),
		host:     withDefault(os.Getenv("DB_HOST"), "localhost"),
		port:     withDefault(os.Getenv("DB_PORT"), "3306"),
		database: database,
	}
	return dbParams
}

func main() {
	ctx := context.Background()
	conn := getDB(getDefaultDbParams())
	defer conn.Close()

	repo := repository.New(conn)

	http.HandleFunc("GET /{name}/{email}/{password}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx, repo)
	})
	fmt.Println("Listening on port 8080, open on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
