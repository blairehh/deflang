package main

import (
	"deflang.io/defcore"
	"fmt"
	"os"
)

func main() {
	stash := defcore.LinuxDisk()
	output := defcore.TerminalOutput()
	session, err := defcore.GetSessionFromEnv(&stash, &output)
	if err != nil {
		fmt.Println("no session")
		return
	}
	run(os.Args[1:], session)
}

func run(input []string, session *defcore.Session) {
	if len(input) == 0 {
		fmt.Println("Invalid name")
		return
	}
	variableName := input[0]
	if !defcore.IsValidVariableName(variableName) {
		fmt.Println("Invalid name " + variableName)
		return
	}

	session.CreateVariable(variableName)
}
