package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var visitorName = r.URL.Query().Get("name")
	if (visitorName == "") {
		visitorName = "World"
	}
	fmt.Fprintln(w, fmt.Sprintf("Hello, %s!", visitorName))
}

func main() {
	http.HandleFunc("/", ServeHTTP)	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
