package defcore

import "fmt"

type Output interface {
	Write(value string)
}

type terminalOutput struct {
}

func TerminalOutput() Output {
	return &terminalOutput{}
}

func (output *terminalOutput) Write(value string) {
	fmt.Println(value)
}
