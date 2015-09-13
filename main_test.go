package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleHi_Recorder(t *testing.T) {
	rw := httptest.NewRecorder()
	handleHi(rw, req(t, "GET / HTTP/1.0\r\n\r\n"))
	if !strings.Contains(rw.Body.String(), "visitor number") {
		t.Errorf("Unexpected output: %s", rw.Body)
	}
}

func req(t *testing.T, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}
