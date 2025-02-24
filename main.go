package main

import (
	"FirstProject/src/go/_generated/repository"
	"FirstProject/src/go/db"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

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

func main() {
	ctx := context.Background()
	conn := db.GetDB(db.GetDefaultDbParams())
	defer conn.Close()

	repo := repository.New(conn)

	http.HandleFunc("GET /{name}/{email}/{password}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ctx, repo)
	})
	fmt.Println("Listening on port 8080, open on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
