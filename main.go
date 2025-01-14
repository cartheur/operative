package main

import (
	"fmt";
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Hey there from Go! two-hundred dollars!")
	})

	http.ListenAndServe(":8080", nil)
}
