package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

func main()  {
		http.HandleFunc("/secret", Secret)
		http.HandleFunc("/configmap", ConfigMap)
		http.HandleFunc("/", Hello)
		http.ListenAndServe(":8001", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

  fmt.Fprintf(w, "Hello! I`m %s. I`m %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")

	if err != nil {
		log.Fatalf("Error reading File", err)
	}

  fmt.Fprintf(w, "My family: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

  fmt.Fprintf(w, "User: %s Password: %s", user, password)
}