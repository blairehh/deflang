package main

import (
	"deflang.io/deftest"
	"testing"
)

func TestGetValue(t *testing.T) {
	_, output, session := deftest.NewTest(map[string]string{"myvar": "mydata"})

	run("myvar", session, []string{})

	if output.Data != "mydata" {
		t.Fatal("expected output to be 'mydata', but was '", output.Data, "'")
	}
}

func TestSetValue(t *testing.T) {
	stash, _, session := deftest.NewTest(map[string]string{})
	run("myvar", session, []string{"=", "mydata"})

	actual := stash.Data["myvar"]
	if actual != "mydata" {
		t.Fatal("expected output to be 'mydata', but was '", actual, "'")
	}
}
