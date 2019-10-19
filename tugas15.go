package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 10; i++ {
		fmt.Fprintln(w, i)
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Tugas 15 GOLANG NIOMIC")
	})
	http.HandleFunc("/index", index)
	fmt.Println("starting web server on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
