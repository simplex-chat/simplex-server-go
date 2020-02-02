package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var server *string

func TestMain(m *testing.M) {
	server = flag.String("server", "http://localhost:8080", "simplex messaging server URI")
	os.Exit(m.Run())
}

// TestHello tests "Hello Heroku"
func TestHello(t *testing.T) {
	fmt.Println(*server)
	resp, err := http.Get(*server)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "Hello World\n" {
		t.Error(string(body))
	}
}
