package main

import (
	"bytes"
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

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	yourVisitNumber := visitors.n
	visitors.Unlock()
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	buf.WriteString("<h1 style='color: ")
	buf.WriteString(r.FormValue("color"))
	buf.WriteString("'>Welcome!</h1>You are visitor number ")
	buf.WriteString(fmt.Sprint(yourVisitNumber))
	buf.WriteString("!")
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf.Bytes())
}

func main() {
	log.Printf("Starting on port 10080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("0.0.0.0:10080", nil))
}
