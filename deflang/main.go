package main

import (
	"deflang.io/defcore"
	"fmt"
)

func main() {
	stash := defcore.LinuxDisk()
	output := defcore.TerminalOutput()
	session := defcore.CreateSession(&stash, &output)
	fmt.Println(session.Id)
}
