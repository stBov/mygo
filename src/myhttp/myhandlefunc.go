package myhttp

import (
	"net/http"
)


func HandleFunc(){
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
