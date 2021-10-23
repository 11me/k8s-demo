package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int64
	Name string
}

func main() {

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", "postgres", "postgres",
		os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!")
	})

	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, os.Getenv("VERSION"))
	})

	r.HandleFunc("/get/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		sid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Error(err)
			return
		}
		rows, err := db.Query(`select * from demo.user where id = $1`, sid)
		if err != nil {
			log.Error(err)
			io.WriteString(w, "Error")
			return
		}
		var users []User
		for rows.Next() {
			var user User
			rows.Scan(&user.ID, &user.Name)
			users = append(users, user)
		}
		jsonRes, err := json.Marshal(users)
		if err != nil {
			log.Error(err)
			return
		}
		io.WriteString(w, string(jsonRes))
		return
	})
	log.Info("Listening")
	http.ListenAndServe(":8080", r)
}
