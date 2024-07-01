package adder

import (
	"computer/multiplexer"
	"reflect"
	"testing"
)

func TestHalfAdder_Run(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		h    HalfAdder
		args args
		want AddResult
	}{
		{
			"HalfAdder false, false",
			HalfAdder{},
			args{false, false},
			AddResult{Cf: false, Sum: false},
		},
		{
			"HalfAdder false, true",
			HalfAdder{},
			args{false, true},
			AddResult{Cf: false, Sum: true},
		},
		{
			"HalfAdder true, false",
			HalfAdder{},
			args{true, false},
			AddResult{Cf: false, Sum: true},
		},
		{
			"HalfAdder true, true",
			HalfAdder{},
			args{true, true},
			AddResult{Cf: true, Sum: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HalfAdder{}
			if got := h.Run(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HalfAdder.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullAdder_Run(t *testing.T) {
	type fields struct {
		HalfAdderInterface HalfAdderInterface
	}
	type args struct {
		a  bool
		b  bool
		cf bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AddResult
	}{
		{
			"FullAdder false, false, false",
			fields{HalfAdder{}},
			args{false, false, false},
			AddResult{Cf: false, Sum: false},
		},
		{
			"FullAdder false, false, true",
			fields{HalfAdder{}},
			args{false, false, true},
			AddResult{Cf: false, Sum: true},
		},
		{
			"FullAdder false, true, false",
			fields{HalfAdder{}},
			args{false, true, false},
			AddResult{Cf: false, Sum: true},
		},
		{
			"FullAdder false, true, true",
			fields{HalfAdder{}},
			args{false, true, true},
			AddResult{Cf: true, Sum: false},
		},
		{
			"FullAdder true, true, false",
			fields{HalfAdder{}},
			args{true, true, false},
			AddResult{Cf: true, Sum: false},
		},
		{
			"FullAdder true, true, true",
			fields{HalfAdder{}},
			args{true, true, true},
			AddResult{Cf: true, Sum: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FullAdder{
				HalfAdderInterface: tt.fields.HalfAdderInterface,
			}
			if got := f.Run(tt.args.a, tt.args.b, tt.args.cf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FullAdder.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdder_Run(t *testing.T) {
	type fields struct {
		HalfAdderInterface HalfAdderInterface
		FullAdderInterface FullAdderInterface
	}
	type args struct {
		a multiplexer.Bool4bit
		b multiplexer.Bool4bit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   AddResult_4bit
	}{
		{
			"Adder 0000, 0000",
			fields{HalfAdder{}, FullAdder{HalfAdder{}}},
			args{multiplexer.Bool4bit{B0: false, B1: false, B2: false, B3: false}, multiplexer.Bool4bit{B0: false, B1: false, B2: false, B3: false}},
			AddResult_4bit{Cf: false, Sum: multiplexer.Bool4bit{B0: false, B1: false, B2: false, B3: false}},
		},
		{
			"Adder 0001, 0011",
			fields{HalfAdder{}, FullAdder{HalfAdder{}}},
			args{multiplexer.Bool4bit{B0: true, B1: false, B2: false, B3: false}, multiplexer.Bool4bit{B0: true, B1: true, B2: false, B3: false}},
			AddResult_4bit{Cf: false, Sum: multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: false}},
		},
		{
			"Adder 1111, 1111",
			fields{HalfAdder{}, FullAdder{HalfAdder{}}},
			args{multiplexer.Bool4bit{B0: false, B1: true, B2: true, B3: true}, multiplexer.Bool4bit{B0: true, B1: true, B2: true, B3: true}},
			AddResult_4bit{Cf: true, Sum: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: false}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adder := Adder{
				HalfAdderInterface: tt.fields.HalfAdderInterface,
				FullAdderInterface: tt.fields.FullAdderInterface,
			}
			if got := adder.Run(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adder.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
