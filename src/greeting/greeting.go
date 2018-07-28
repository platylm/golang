package greeting

import (
	"encoding/json"
	"net/http"
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
