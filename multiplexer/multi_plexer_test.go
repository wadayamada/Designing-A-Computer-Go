package multiplexer

import (
	"testing"
)

func TestMultiplexer2to1(t *testing.T) {
	type args struct {
		a   bool
		b   bool
		sel bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Multiplexer2to1 sel=trueの時はbを返す",
			args{true, false, true},
			false,
		},
		{
			"Multiplexer2to1 sel=falseの時はaを返す",
			args{true, false, false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplexer2to1(tt.args.a, tt.args.b, tt.args.sel); got != tt.want {
				t.Errorf("Multiplexer2to1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplexer4to1(t *testing.T) {
	type args struct {
		a   bool
		b   bool
		c   bool
		d   bool
		sel bool2bit
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Multiplexer4to1 00の時はaを返す",
			args{true, false, false, false, bool2bit{false, false}},
			true,
		},
		{
			"Multiplexer4to1 01の時はbを返す",
			args{false, true, false, false, bool2bit{false, true}},
			true,
		},
		{
			"Multiplexer4to1 10の時はcを返す",
			args{false, false, true, false, bool2bit{true, false}},
			true,
		},
		{
			"Multiplexer4to1 11の時はdを返す",
			args{false, false, false, true, bool2bit{true, true}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplexer4to1(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.sel); got != tt.want {
				t.Errorf("Multiplexer4to1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplexer16to1(t *testing.T) {
	type args struct {
		i0  bool
		i1  bool
		i2  bool
		i3  bool
		i4  bool
		i5  bool
		i6  bool
		i7  bool
		i8  bool
		i9  bool
		i10 bool
		i11 bool
		i12 bool
		i13 bool
		i14 bool
		i15 bool
		sel bool4bit
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Multiplexer16to1 0000の時はi0を返す",
			args{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, bool4bit{false, false, false, false}},
			true,
		},
		{
			"Multiplexer16to1 0101の時はi5を返す",
			args{false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, bool4bit{false, true, false, true}},
			true,
		},
		{
			"Multiplexer16to1 1100の時はi10を返す",
			args{false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, bool4bit{true, false, true, false}},
			true,
		},
		{
			"Multiplexer16to1 1111の時はi15を返す",
			args{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, bool4bit{true, true, true, true}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplexer16to1(tt.args.i0, tt.args.i1, tt.args.i2, tt.args.i3, tt.args.i4, tt.args.i5, tt.args.i6, tt.args.i7, tt.args.i8, tt.args.i9, tt.args.i10, tt.args.i11, tt.args.i12, tt.args.i13, tt.args.i14, tt.args.i15, tt.args.sel); got != tt.want {
				t.Errorf("Multiplexer16to1() = %v, want %v", got, tt.want)
			}
		})
	}
}
