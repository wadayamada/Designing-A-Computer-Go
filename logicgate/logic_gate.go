package logicgate

import "computer/nand"

// NANDでAND, OR, NOTを実装する

type NOT struct {
	A bool
}

type AND struct {
	A, B bool
}

type OR struct {
	A, B bool
}

func (n NOT) Out() bool {
	return nand.NAND{A: n.A, B: n.A}.Out()
}

func (a AND) Out() bool {
	return NOT{nand.NAND{A: a.A, B: a.B}.Out()}.Out()
}

func (o OR) Out() bool {
	return nand.NAND{A: NOT{o.A}.Out(), B: NOT{o.B}.Out()}.Out()
}
