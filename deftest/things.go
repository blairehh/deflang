package deftest

import "deflang.io/defcore"

type TestStash struct {
	Data map[string]string
}

func (stash *TestStash) Has(item string) bool {
	_, ok := stash.Data[item]
	return ok
}

func (stash *TestStash) Init(sessionId string) {

}

func (stash *TestStash) ReadData(sessionId, item string) string {
	return stash.Data[item]
}

func (stash *TestStash) WriteData(sessionId, item, data string) {
	stash.Data[item] = data
}

func (stash *TestStash) Add(sessionId, item string) error {
	stash.Data[item] = ""
	return nil
}

type TestOutput struct {
	Data string
}

func (out *TestOutput) Write(value string) {
	out.Data += value
}

func NewTest(initData map[string]string) (*TestStash, *TestOutput, *defcore.Session) {
	testStash := TestStash{
		Data: initData,
	}
	var stash defcore.Stash = &testStash

	output := TestOutput{}
	var out defcore.Output = &output
	session := defcore.GetSession("id", &stash, &out)

	return &testStash, &output, session
}
