package main

import (
	"fmt"
	"net/http"
	"os"
)

func linuxVersionHandler(w http.ResponseWriter, r *http.Request) {
	// Read the Linux version from /proc/version
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		http.Error(w, "Failed to read version: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Linux Version: %s", string(data))
}

func main() {
	http.HandleFunc("/version", linuxVersionHandler)

	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
