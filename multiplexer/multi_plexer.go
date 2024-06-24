package multiplexer

import (
	"computer/logicgate"
)

// 入力から出力を選択するマルチプレクサを実装する

// 2to1マルチプレクサ(1bit)
// selがtrueの時はbを返し、falseの時はaを返す
func Multiplexer2to1(a, b, sel bool) bool {
	return logicgate.OR{
		A: logicgate.AND{
			A: a,
			B: logicgate.NOT{A: sel}.Out(),
		}.Out(),
		B: logicgate.AND{
			A: b,
			B: sel,
		}.Out(),
	}.Out()
}

// 4to1マルチプレクサ(1bit)
// selが00の時はaを返し、01の時はbを返し、10の時はcを返し、11の時はdを返す
func Multiplexer4to1(a, b, c, d bool, sel bool2bit) (bool, error) {
	return Multiplexer2to1(
		Multiplexer2to1(a, b, sel.b1),
		Multiplexer2to1(c, d, sel.b1),
		sel.b0,
	), nil
}

// 00->{b1=0, b0=0}, 01->{b1=0, b0=1}, 10->{b1=1, b0=0}, 11->{b1=1, b0=1}
type bool2bit struct {
	b1 bool
	b0 bool
}

// 0101->{b3=0, b2=1, b1=0, b0=1}, 1100->{b3=1, b2=1, b1=0, b0=0}
type bool4bit struct {
	b3 bool
	b2 bool
	b1 bool
	b0 bool
}

// 16to1マルチプレクサ(1bit)

// 2to1マルチプレクサ(4bit)

// 16to1マルチプレクサ(4bit)

// 16to1マルチプレクサ(8bit)
