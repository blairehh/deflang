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
	run(os.Args[0], session, os.Args[1:])
}

func run(name string, session *defcore.Session, input []string) {
	variable := session.GetVariable(name)

	if len(input) == 0 {
		(*session.Output).Write(variable.GetValue())
		return
	}

	if isVariableAssignment(input[0]) {
		fmt.Println("is assigning data " + valueToAssign(input))
		variable.SetValue(valueToAssign(input))
	} else {
		fmt.Println("what am i doing?")
	}

}

func isVariableAssignment(value string) bool {
	return value == "="
}

func valueToAssign(input []string) string {
	return input[1]
}
