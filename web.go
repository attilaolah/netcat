package main

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", sloooww)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}

func sloooww(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if t, err := strconv.ParseInt(q.Get("t"), 10, 32); err == nil {
		time.Sleep(time.Duration(t) * time.Second)
	}
	if s, err := strconv.ParseInt(q.Get("c"), 10, 32); err == nil {
		w.WriteHeader(int(s))
	}
}
