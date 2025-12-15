package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("🌐 LinkedIn Automation Lab — Web Dashboard")
	fmt.Println("Listening on http://localhost:8080")

	fs := http.FileServer(http.Dir("./ui/web"))
	http.Handle("/", fs)

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
