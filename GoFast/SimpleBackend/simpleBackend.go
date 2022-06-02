package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, `
	          ##         .
       ## ## ##        ==
    ## ## ## ## ##    ===
 /"""""""""""""""""\___/ ===
{                       /  ===-
 \______ O           __/
  \    \         __/
   \____\_______/

  Hello from Backend!
`)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
