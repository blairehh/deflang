package defcore

type VariableType string

const (
	String = "String"
	Number = "Number"
	Bool   = "Bool"
)

type Variable struct {
	session *Session
	name    string
}

func (variable *Variable) GetValue() string {
	return (*variable.session.stash).ReadData(variable.session.Id, variable.name)
}

func (variable *Variable) SetValue(value string) {
	(*variable.session.stash).WriteData(variable.session.Id, variable.name, value)
}
