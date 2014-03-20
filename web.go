package main

import (
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	pCode    = "c"
	pTimeout = "t"
)

func main() {
	http.HandleFunc("/", sloooww)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}

func sloooww(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if t, err := strconv.ParseInt(q.Get(pTimeout), 10, 32); err == nil {
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
	for key, vals := range q {
		for i, val := range vals {
			if i == 0 && (key == pCode || key == pTimeout) {
				continue
			}
			w.Header().Add(key, val)
		}
	}
	if s, err := strconv.ParseInt(q.Get(pCode), 10, 32); err == nil {
		w.WriteHeader(int(s))
		return
	}
	if r.URL.Path == "/loop" {
		w.WriteHeader(http.StatusFound)
	}
}
