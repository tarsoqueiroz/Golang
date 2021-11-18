//# https://github.com/ruanbekker/go-hostname/blob/main/app.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	myhostname, _ := os.Hostname()
	fmt.Fprintln(w, "Hostname:", myhostname)
}

func main() {
	const port string = "8765"
	fmt.Println("Server listening on port", port)
	http.HandleFunc("/", hostnameHandler)
	http.ListenAndServe(":"+port, nil)
}
