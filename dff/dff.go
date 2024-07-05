package dff

import (
	"computer/logicgate"
	"computer/multiplexer"
	"errors"
)

type DFF4bit struct {
	DFF0 DFFInterface
	DFF1 DFFInterface
	DFF2 DFFInterface
	DFF3 DFFInterface
}

func (dff DFF4bit) Read() multiplexer.Bool4bit {
	Q0 := dff.DFF0.Read()
	Q1 := dff.DFF1.Read()
	Q2 := dff.DFF2.Read()
	Q3 := dff.DFF3.Read()
	return multiplexer.Bool4bit{B0: Q0, B1: Q1, B2: Q2, B3: Q3}
}

func (dff DFF4bit) Write(D multiplexer.Bool4bit) multiplexer.Bool4bit {
	Q0 := dff.DFF0.Write(D.B0)
	Q1 := dff.DFF1.Write(D.B1)
	Q2 := dff.DFF2.Write(D.B2)
	Q3 := dff.DFF3.Write(D.B3)
	return multiplexer.Bool4bit{B0: Q0, B1: Q1, B2: Q2, B3: Q3}
}

type DFF4bitInterface interface {
	Read() multiplexer.Bool4bit
	Write(multiplexer.Bool4bit) multiplexer.Bool4bit
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
