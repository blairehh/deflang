package main

import (
	"deflang.io/deftest"
	"testing"
)

func TestCreateVariable(t *testing.T) {
	stash, _, session := deftest.NewTest(map[string]string{})

	run([]string{"myvar"}, session)

	if !stash.Has("myvar") {
		t.Fatal("expected stash to contain 'myvar' but did not", stash.Data)
	}
}
