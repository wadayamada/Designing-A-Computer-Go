package adder

import (
	"computer/logicgate"
	"computer/multiplexer"
)

type AdderInterface interface {
	Run(a, b multiplexer.Bool4bit) AddResult_4bit
}

type AddResult_4bit struct {
	Cf  bool
	Sum multiplexer.Bool4bit
}

type Adder struct {
	HalfAdderInterface HalfAdderInterface
	FullAdderInterface FullAdderInterface
}

type HalfAdderInterface interface {
	Run(a, b bool) AddResult
}

type AddResult struct {
	Cf, Sum bool
}

type HalfAdder struct {
}

func (h HalfAdder) Run(a, b bool) AddResult {
	cf := logicgate.AND{A: a, B: b}.Out()
	sum := logicgate.AND{A: logicgate.OR{A: a, B: b}.Out(), B: logicgate.NOT{A: cf}.Out()}.Out()
	return AddResult{Cf: cf, Sum: sum}
}

type FullAdderInterface interface {
	Run(a1, b1, cf bool) AddResult
}

type FullAdder struct {
	HalfAdderInterface HalfAdderInterface
}

func (f FullAdder) Run(a, b, cf bool) AddResult {
	add_a_b := f.HalfAdderInterface.Run(a, b)
	add_a_b_cf := f.HalfAdderInterface.Run(add_a_b.Sum, cf)
	add_cf := logicgate.OR{A: add_a_b.Cf, B: add_a_b_cf.Cf}.Out()
	return AddResult{Cf: add_cf, Sum: add_a_b_cf.Sum}
}

func (adder Adder) Run(a, b multiplexer.Bool4bit) AddResult_4bit {
	addResult0 := adder.HalfAdderInterface.Run(a.B0, b.B0)
	addResult1 := adder.FullAdderInterface.Run(a.B1, b.B1, addResult0.Cf)
	addResult2 := adder.FullAdderInterface.Run(a.B2, b.B2, addResult1.Cf)
	addResult3 := adder.FullAdderInterface.Run(a.B3, b.B3, addResult2.Cf)
	return AddResult_4bit{addResult3.Cf, multiplexer.Bool4bit{B0: addResult0.Sum, B1: addResult1.Sum, B2: addResult2.Sum, B3: addResult3.Sum}}
}
