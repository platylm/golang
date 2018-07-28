package greeting

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Messeage struct {
	Txt string `json:"txt"`
}

type User struct {
	Name string `json:"name"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	values := r.URL.Query()
	name := values.Get("name")
	messeage := Messeage{Txt: "hello " + name}

	encoder := json.NewEncoder(w)
	encoder.Encode(messeage)
}

func HelloPost(w http.ResponseWriter, r *http.Request) {
	var u User
	body := json.NewDecoder(r.Body)
	body.Decode(&u)
	messeage := Messeage{Txt: "hello " + u.Name}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(messeage)
}

func GreetingMessage(name string) string {
	prefix := RetrivePrefix()
	return fmt.Sprintf("%s %s", prefix, name)

}

func RetrivePrefix() string {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/bootcamp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var prefix string
	err = db.QueryRow("SELECT message FROM greeting  WHERE id = ?", 1).Scan(&prefix)
	if err != nil {
		log.Fatal(err)
	}
	return prefix
}
