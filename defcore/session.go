package defcore

import (
	"errors"
	"os"
)

type Session struct {
	Id     string
	stash  *Stash
	Output *Output
}

const ENV_SESSION_ID = "DEFLANG_SESSION_ID"

func CreateSession(stash *Stash, output *Output) *Session {
	id := RandStringRunes(20)
	return GetSession(id, stash, output)
}

func GetSession(id string, stash *Stash, output *Output) *Session {
	session := Session{
		Id:     id,
		stash:  stash,
		Output: output,
	}
	(*stash).Init(id)
	return &session
}

func GetSessionFromEnv(stash *Stash, output *Output) (*Session, error) {
	sessionId := os.Getenv(ENV_SESSION_ID)
	if sessionId == "" {
		return nil, errors.New("no session")
	}
	return GetSession(sessionId, stash, output), nil
}

func (session *Session) GetVariable(name string) Variable {
	return Variable{
		session: session,
		name:    name,
	}
}

func (session *Session) CreateVariable(name string) Variable {
	variable := session.GetVariable(name)
	(*session.stash).Add(session.Id, variable.name)
	return variable
}
