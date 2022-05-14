package tests

import (
	"fmt"
	"testing"
)

var home_page_handler = func(msg string) string {
	msg_func := "Teste"
	return msg_func
}

func TestHomePageHandler(t *testing.T) {
	home_page_handler = func(msg string) string {
        if msg != "expected" {
            t.Fatal("good message")
        }
        return "something"
    }
}