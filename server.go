package main

import (
	"net/http"
	"os"
	"fmt"
)

func main()  {
		http.HandleFunc("/", Hello)
		http.ListenAndServe(":8001", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

  fmt.Fprintf(w, "Hello! I`m %s. I`m %s", name, age)
}