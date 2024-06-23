package nand

// NANDを定義する。NANDのハード的な仕組みも復習したい

type NAND struct {
	A, B bool
}

func (n NAND) Out() bool {
	return !(n.A && n.B)
}

type LogicGate interface {
	Out() bool
}
