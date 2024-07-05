package dff

import (
	"computer/logicgate"
	"computer/multiplexer"
	"errors"
)

// TODO: DFF4bitをDFFでちゃんと実装する
type DFF4bit struct {
	D multiplexer.Bool4bit
	Q multiplexer.Bool4bit
}

type DFF4bitInterface interface {
	Read() multiplexer.Bool4bit
	Write(multiplexer.Bool4bit)
}

type DFFInterface interface {
	Read() bool
	Write(bool) bool
}

type DFF struct {
	RSFFInterface RSFFInterface
}

func (dff DFF) Read() bool {
	return dff.RSFFInterface.Read()
}

func (dff DFF) Write(D bool) bool {
	clock := true
	S := logicgate.AND{A: D, B: clock}.Out()
	R := logicgate.AND{A: logicgate.NOT{A: S}.Out(), B: clock}.Out()
	Q, _ := dff.RSFFInterface.Run(R, S)
	return Q
}

type RSFFInterface interface {
	Run(Reset bool, Set bool) (bool, error)
	Read() bool
}

type RSFF struct {
	Q     bool
	Q_not bool
}

func (rsff RSFF) Read() bool {
	return rsff.Q
}

func (rsff RSFF) Run(Reset bool, Set bool) (bool, error) {
	if Reset && Set {
		return false, errors.New("Reset and Set are both true")
	}

	next_Q := logicgate.NOT{A: logicgate.OR{A: Reset, B: rsff.Q_not}.Out()}.Out()
	next_Q_not := logicgate.NOT{A: logicgate.OR{A: Set, B: next_Q}.Out()}.Out()

	next_next_Q := logicgate.NOT{A: logicgate.OR{A: Reset, B: next_Q_not}.Out()}.Out()
	next_next_Q_not := logicgate.NOT{A: logicgate.OR{A: Set, B: next_next_Q}.Out()}.Out()

	rsff.Q = next_next_Q
	rsff.Q_not = next_next_Q_not

	return rsff.Q, nil
}
