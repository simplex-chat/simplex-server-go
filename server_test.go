package main

import (
	"testing"

	testsimplex "github.com/simplex-chat/test-simplex-server"
)

// TestHello tests "Hello Heroku"
func TestHello(t *testing.T) {
	ok, err := testsimplex.RunTestHello("http://localhost:8080")
	if !ok {
		t.Error(err)
	}
}
