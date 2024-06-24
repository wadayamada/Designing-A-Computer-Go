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
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Multiplexer4to1 00の時はaを返す",
			args{true, false, false, false, bool2bit{false, false}},
			true,
			false,
		},
		{
			"Multiplexer4to1 01の時はbを返す",
			args{false, false, true, false, bool2bit{false, true}},
			true,
			false,
		},
		{
			"Multiplexer4to1 10の時はcを返す",
			args{false, true, false, false, bool2bit{true, false}},
			true,
			false,
		},
		{
			"Multiplexer4to1 11の時はdを返す",
			args{false, false, false, true, bool2bit{true, true}},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Multiplexer4to1(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.sel)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiplexer4to1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multiplexer4to1() = %v, want %v", got, tt.want)
			}
		})
	}
}
