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

type bool8bit struct {
	b7 bool
	b6 bool
	b5 bool
	b4 bool
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
func Multiplexer2to1_4bit(a, b bool4bit, sel bool) bool4bit {
	return bool4bit{
		b3: Multiplexer2to1(a.b3, b.b3, sel),
		b2: Multiplexer2to1(a.b2, b.b2, sel),
		b1: Multiplexer2to1(a.b1, b.b1, sel),
		b0: Multiplexer2to1(a.b0, b.b0, sel),
	}
}

// 4to1マルチプレクサ(4bit)
func Multiplexer4to1_4bit(a, b, c, d bool4bit, sel bool2bit) bool4bit {
	return Multiplexer2to1_4bit(
		Multiplexer2to1_4bit(a, b, sel.b0),
		Multiplexer2to1_4bit(c, d, sel.b0),
		sel.b1,
	)
}

// 16to1マルチプレクサ(4bit)
func Multiplexer16to1_4bit(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 bool4bit, sel bool4bit) bool4bit {
	return Multiplexer4to1_4bit(
		Multiplexer4to1_4bit(i0, i1, i2, i3, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1_4bit(i4, i5, i6, i7, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1_4bit(i8, i9, i10, i11, bool2bit{sel.b1, sel.b0}),
		Multiplexer4to1_4bit(i2, i13, i14, i15, bool2bit{sel.b1, sel.b1}),
		bool2bit{sel.b3, sel.b2},
	)
}

// 16to1マルチプレクサ(8bit)
func Multiplexer16to1_8bit(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 bool8bit, sel bool4bit) bool8bit {
	return bool8bit{
		b7: Multiplexer16to1(i0.b7, i1.b7, i2.b7, i3.b7, i4.b7, i5.b7, i6.b7, i7.b7, i8.b7, i9.b7, i10.b7, i11.b7, i12.b7, i13.b7, i14.b7, i15.b7, sel),
		b6: Multiplexer16to1(i0.b6, i1.b6, i2.b6, i3.b6, i4.b6, i5.b6, i6.b6, i7.b6, i8.b6, i9.b6, i10.b6, i11.b6, i12.b6, i13.b6, i14.b6, i15.b6, sel),
		b5: Multiplexer16to1(i0.b5, i1.b5, i2.b5, i3.b5, i4.b5, i5.b5, i6.b5, i7.b5, i8.b5, i9.b5, i10.b5, i11.b5, i12.b5, i13.b5, i14.b5, i15.b5, sel),
		b4: Multiplexer16to1(i0.b4, i1.b4, i2.b4, i3.b4, i4.b4, i5.b4, i6.b4, i7.b4, i8.b4, i9.b4, i10.b4, i11.b4, i12.b4, i13.b4, i14.b4, i15.b4, sel),
		b3: Multiplexer16to1(i0.b3, i1.b3, i2.b3, i3.b3, i4.b3, i5.b3, i6.b3, i7.b3, i8.b3, i9.b3, i10.b3, i11.b3, i12.b3, i13.b3, i14.b3, i15.b3, sel),
		b2: Multiplexer16to1(i0.b2, i1.b2, i2.b2, i3.b2, i4.b2, i5.b2, i6.b2, i7.b2, i8.b2, i9.b2, i10.b2, i11.b2, i12.b2, i13.b2, i14.b2, i15.b2, sel),
		b1: Multiplexer16to1(i0.b1, i1.b1, i2.b1, i3.b1, i4.b1, i5.b1, i6.b1, i7.b1, i8.b1, i9.b1, i10.b1, i11.b1, i12.b1, i13.b1, i14.b1, i15.b1, sel),
		b0: Multiplexer16to1(i0.b0, i1.b0, i2.b0, i3.b0, i4.b0, i5.b0, i6.b0, i7.b0, i8.b0, i9.b0, i10.b0, i11.b0, i12.b0, i13.b0, i14.b0, i15.b0, sel),
	}
}
