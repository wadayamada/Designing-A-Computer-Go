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
func Multiplexer4to1(a, b, c, d bool, sel bool2bit) bool {
	return Multiplexer2to1(
		Multiplexer2to1(a, b, sel.b0),
		Multiplexer2to1(c, d, sel.b0),
		sel.b1,
	)
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
func Multiplexer16to1(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 bool, sel bool4bit) bool {
	return Multiplexer4to1(
		Multiplexer4to1(i0, i1, i2, i3, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1(i4, i5, i6, i7, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1(i8, i9, i10, i11, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1(i12, i13, i14, i15, bool2bit{sel.b1, sel.b1}),
		bool2bit{sel.b3, sel.b2},
	)
}

// 2to1マルチプレクサ(4bit)

// 16to1マルチプレクサ(4bit)

// 16to1マルチプレクサ(8bit)
