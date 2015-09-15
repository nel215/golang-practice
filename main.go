package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var visitors struct {
	sync.Mutex
	n int
}

var colorRx = regexp.MustCompile(`\w*$`)

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	yourVisitNumber := visitors.n
	visitors.Unlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(yourVisitNumber) + "!"))
}

func main() {
	log.Printf("Starting on port 10080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("0.0.0.0:10080", nil))
}
