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
		Multiplexer2to1(a, b, sel.B0),
		Multiplexer2to1(c, d, sel.B0),
		sel.B1,
	)
}

// 00->{B1=0, B0=0}, 01->{B1=0, B0=1}, 10->{B1=1, B0=0}, 11->{B1=1, B0=1}
type bool2bit struct {
	B1 bool
	B0 bool
}

// 0101->{B3=0, B2=1, B1=0, B0=1}, 1100->{B3=1, B2=1, B1=0, B0=0}
type Bool4bit struct {
	B3, B2, B1, B0 bool
}

type Bool8bit struct {
	B7, B6, B5, B4, B3, B2, B1, B0 bool
}

// 16to1マルチプレクサ(1bit)
func Multiplexer16to1(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 bool, sel Bool4bit) bool {
	return Multiplexer4to1(
		Multiplexer4to1(i0, i1, i2, i3, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1(i4, i5, i6, i7, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1(i8, i9, i10, i11, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1(i12, i13, i14, i15, bool2bit{sel.B1, sel.B1}),
		bool2bit{sel.B3, sel.B2},
	)
}

// 2to1マルチプレクサ(4bit)
func Multiplexer2to1_4bit(a, b Bool4bit, sel bool) Bool4bit {
	return Bool4bit{
		B3: Multiplexer2to1(a.B3, b.B3, sel),
		B2: Multiplexer2to1(a.B2, b.B2, sel),
		B1: Multiplexer2to1(a.B1, b.B1, sel),
		B0: Multiplexer2to1(a.B0, b.B0, sel),
	}
}

// 4to1マルチプレクサ(4bit)
func Multiplexer4to1_4bit(a, b, c, d Bool4bit, sel bool2bit) Bool4bit {
	return Multiplexer2to1_4bit(
		Multiplexer2to1_4bit(a, b, sel.B0),
		Multiplexer2to1_4bit(c, d, sel.B0),
		sel.B1,
	)
}

// 16to1マルチプレクサ(4bit)
func Multiplexer16to1_4bit(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 Bool4bit, sel Bool4bit) Bool4bit {
	return Multiplexer4to1_4bit(
		Multiplexer4to1_4bit(i0, i1, i2, i3, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1_4bit(i4, i5, i6, i7, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1_4bit(i8, i9, i10, i11, bool2bit{sel.B1, sel.B0}),
		Multiplexer4to1_4bit(i2, i13, i14, i15, bool2bit{sel.B1, sel.B1}),
		bool2bit{sel.B3, sel.B2},
	)
}

// 16to1マルチプレクサ(8bit)
func Multiplexer16to1_8bit(i0, i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15 Bool8bit, sel Bool4bit) Bool8bit {
	return Bool8bit{
		B7: Multiplexer16to1(i0.B7, i1.B7, i2.B7, i3.B7, i4.B7, i5.B7, i6.B7, i7.B7, i8.B7, i9.B7, i10.B7, i11.B7, i12.B7, i13.B7, i14.B7, i15.B7, sel),
		B6: Multiplexer16to1(i0.B6, i1.B6, i2.B6, i3.B6, i4.B6, i5.B6, i6.B6, i7.B6, i8.B6, i9.B6, i10.B6, i11.B6, i12.B6, i13.B6, i14.B6, i15.B6, sel),
		B5: Multiplexer16to1(i0.B5, i1.B5, i2.B5, i3.B5, i4.B5, i5.B5, i6.B5, i7.B5, i8.B5, i9.B5, i10.B5, i11.B5, i12.B5, i13.B5, i14.B5, i15.B5, sel),
		B4: Multiplexer16to1(i0.B4, i1.B4, i2.B4, i3.B4, i4.B4, i5.B4, i6.B4, i7.B4, i8.B4, i9.B4, i10.B4, i11.B4, i12.B4, i13.B4, i14.B4, i15.B4, sel),
		B3: Multiplexer16to1(i0.B3, i1.B3, i2.B3, i3.B3, i4.B3, i5.B3, i6.B3, i7.B3, i8.B3, i9.B3, i10.B3, i11.B3, i12.B3, i13.B3, i14.B3, i15.B3, sel),
		B2: Multiplexer16to1(i0.B2, i1.B2, i2.B2, i3.B2, i4.B2, i5.B2, i6.B2, i7.B2, i8.B2, i9.B2, i10.B2, i11.B2, i12.B2, i13.B2, i14.B2, i15.B2, sel),
		B1: Multiplexer16to1(i0.B1, i1.B1, i2.B1, i3.B1, i4.B1, i5.B1, i6.B1, i7.B1, i8.B1, i9.B1, i10.B1, i11.B1, i12.B1, i13.B1, i14.B1, i15.B1, sel),
		B0: Multiplexer16to1(i0.B0, i1.B0, i2.B0, i3.B0, i4.B0, i5.B0, i6.B0, i7.B0, i8.B0, i9.B0, i10.B0, i11.B0, i12.B0, i13.B0, i14.B0, i15.B0, sel),
	}
}
