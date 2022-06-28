package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	Stats = map[string]int{
		"John":     0,
		"Isabella": 0,
	}

	calls = []string{}
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Method Not Allowed"))

		return
	} else if name == "Isabella" {
		Stats["Isabella"] += 1
		calls = append(calls, "Isabella")

	} else {
		if name == "John" {
			Stats["John"] += 1
			calls = append(calls, "John")

		}

	}

	fmt.Fprint(w, "Hello, ", name)
	//fmt.Print(w, "Hello, ", name)
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n", Stats)

}

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//fmt.Println(Stats["Bill"], Stats["Eva"])
	// Your solution goes here. Good luck!
}
