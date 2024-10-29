package main

import (
	"github.com/michell179/greetings"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Michell"
	exp := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello("Michell")
	if !exp.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("juan") = %q , %v quiere coincidecia para  %#q, nil`, msg, err, exp)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello() = %q, %v, quiere "", error`, msg, err)
	}
}
