package myhttp

import (
	"fmt"
	"net/http"
)

type dollars float32
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type MyHandler map[string]dollars
func (m MyHandler) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range m {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
func (m MyHandler) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := m[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func Handles(){
	handler := MyHandler{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(handler.list))
	mux.Handle("/price", http.HandlerFunc(IndexHandler))
	http.ListenAndServe("localhost:9000", mux)
}

