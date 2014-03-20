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
	if r.URL.Path == "/loop" {
		l := r.URL.Scheme
		if l == "" {
			l = "http"
		}
		l += "://" + r.Host + r.URL.String()
		w.Header().Set("Location", l)
	}
	if s, err := strconv.ParseInt(q.Get("c"), 10, 32); err == nil {
		w.WriteHeader(int(s))
		return
	}
	if r.URL.Path == "/loop" {
		w.WriteHeader(http.StatusFound)
	}
}
