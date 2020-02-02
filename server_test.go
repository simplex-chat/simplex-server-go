package main

import (
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var server *string

// TestMain initialises server url from -server parameter of "go test"
func TestMain(m *testing.M) {
	server = flag.String("server", "http://localhost:8080", "simplex messaging server URI")
	os.Exit(m.Run())
}

func catchError(t *testing.T) {
	err := recover()
	if err != nil {
		t.Error(err)
	}
}

func req(method, url string, body ...io.Reader) *http.Request {
	var reqBody io.Reader
	if body != nil {
		reqBody = body[0]
	}
	r, _ := http.NewRequest(method, *server+url, reqBody)
	return r
}

func httpRequest(t *testing.T, req *http.Request) []byte {
	defer catchError(t)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}

func testApi(t *testing.T, req *http.Request, expectedBody string) {
	body := httpRequest(t, req)
	if string(body) != expectedBody {
		t.Error(string(body))
	}
}

// TestHello tests "Hello World"
func TestHello(t *testing.T) {
	testApi(t, req("GET", ""), "Hello World\n")
}

// TestRecipientApi tests recipient REST API
func TestRecipientApi(t *testing.T) {
	testApi(t, req("POST", "/connection"),
		"createConnection not implemented\n")
	testApi(t, req("PUT", "/connection/123"),
		"secureConnection not implemented\n")
	testApi(t, req("DELETE", "/connection/123"),
		"deleteConnection not implemented\n")
	testApi(t, req("GET", "/connection/123/messages"),
		"getMessages not implemented\n")
	testApi(t, req("GET", "/connection/123/messages/456"),
		"getMessage not implemented\n")
	testApi(t, req("DELETE", "/connection/123/messages/456"),
		"deleteMessage not implemented\n")
}

// TestSenderApi tests sender REST API
func TestSenderApi(t *testing.T) {
	testApi(t, req("POST", "/connection/123/messages"),
		"sendMessage not implemented\n")
}
