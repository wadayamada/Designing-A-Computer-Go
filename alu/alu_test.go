package alu

import (
	"computer/adder"
	"computer/multiplexer"
	"reflect"
	"testing"
)

func TestALU_Run(t *testing.T) {
	type fields struct {
		AdderInterface adder.AdderInterface
	}
	type args struct {
		opecode  multiplexer.Bool4bit
		imm      multiplexer.Bool4bit
		register Register
		in_a     multiplexer.Bool4bit
		in_b     multiplexer.Bool4bit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Register
	}{
		{
			"0001 MOV A, B",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0100 MOV B, A",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0011 MOV A, IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0111 MOV B, IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: true},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0010 IN A",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0110 IN B",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1001 OUT B",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: true},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				CF:  false,
			},
		},
		{
			"1011 OUT IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: false, B1: true, B0: true},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				CF:  false,
			},
		},
		{
			"0000 ADD A, IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"0101 ADD B, IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: true},
				imm:     multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1000 ADD A=4, B=2 -> A=6, B=2",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1000 ADD A=10, B=6 -> A=0, B=6, CF=1",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: true, B2: false, B1: true, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: true, B1: true, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  true,
			},
		},
		{
			"1111 JMP IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1110 JNC IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: false},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1110 JNC IMM",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: false},
				imm:     multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  true,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  true,
			},
		},
		{
			"1010 MOV C, A",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: false, B1: true, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					C:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				C:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1100 MOV A, C",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: true, B1: false, B0: false},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					C:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				C:   multiplexer.Bool4bit{B3: true, B2: true, B1: true, B0: true},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
		{
			"1101 SUB A, B",
			fields{adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			args{
				opecode: multiplexer.Bool4bit{B3: true, B2: true, B1: false, B0: true},
				imm:     multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				register: Register{
					A:   multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: true},  // 5
					B:   multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false}, // 2
					C:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
					CF:  false,
				},
				in_a: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				in_b: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
			},
			Register{
				A:   multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				B:   multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false}, // 2
				C:   multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				IP:  multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true},
				Out: multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				CF:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alu := ALU{
				AdderInterface: tt.fields.AdderInterface,
			}
			if got := alu.Run(tt.args.opecode, tt.args.imm, tt.args.register, tt.args.in_a, tt.args.in_b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ALU.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
