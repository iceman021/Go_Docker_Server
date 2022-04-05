package main

import(
"fmt"
"log"
"net/http"
"html"
)


func main(){


http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

http.HandleFunc("/Greetings", func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello")
})

log.Fatal(http.ListenAndServe(":8081", nil))

}