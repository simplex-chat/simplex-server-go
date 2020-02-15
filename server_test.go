package main

import (
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func req(method, url string, body ...string) *http.Request {
	var reqBody io.Reader
	if body != nil {
		reqBody = strings.NewReader(body[0])
	}
	r, _ := http.NewRequest(method, *server+url, reqBody)
	return r
}

func httpRequest(t *testing.T, req *http.Request) (resp *http.Response, respBody []byte) {
	defer catchError(t)
	client := &http.Client{}
	resp, _ = client.Do(req)
	defer resp.Body.Close()
	respBody, _ = ioutil.ReadAll(resp.Body)
	return resp, respBody
}

func testApi(t *testing.T, req *http.Request, expectedStatus int, expectedBody string) {
	resp, body := httpRequest(t, req)
	if string(body) != expectedBody {
		t.Error(string(body))
	}
	if resp.StatusCode != expectedStatus {
		t.Error(resp.StatusCode)
	}
}

// TestHello tests "Hello World"
func TestHello(t *testing.T) {
	testApi(t, req("GET", ""), 200, "Hello World\n")
}

// TestCreateConnection
func TestCreateConnection(t *testing.T) {
	testApi(t, req("POST", "/connection", `{"recipient": "123"}`),
		200, "Ok")
	testApi(t, req("POST", "/connection", `1`),
		400, "Bad Request")
	testApi(t, req("POST", "/connection", `"recipient"`),
		400, "Bad Request")
	testApi(t, req("POST", "/connection", `[]`),
		400, "Bad Request")
	testApi(t, req("POST", "/connection", `{}`),
		400, "Bad Request")
	testApi(t, req("POST", "/connection", `{"recipient": 1}`),
		400, "Bad Request")
	testApi(t, req("POST", "/connection", `{"recipient": "123", "unknown": "123"}`),
		400, "Bad Request")
}

// TestRecipientApi tests recipient REST API
func TestRecipientApi(t *testing.T) {
	testApi(t, req("PUT", "/connection/123"),
		200, "secureConnection not implemented\n")
	testApi(t, req("DELETE", "/connection/123"),
		200, "deleteConnection not implemented\n")
	testApi(t, req("GET", "/connection/123/messages"),
		200, "getMessages not implemented\n")
	testApi(t, req("GET", "/connection/123/messages/456"),
		200, "getMessage not implemented\n")
	testApi(t, req("DELETE", "/connection/123/messages/456"),
		200, "deleteMessage not implemented\n")
}

// TestSenderApi tests sender REST API
func TestSenderApi(t *testing.T) {
	testApi(t, req("POST", "/connection/123/messages"),
		200, "sendMessage not implemented\n")
}
