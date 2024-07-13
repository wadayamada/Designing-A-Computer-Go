package alu

import (
	"computer/adder"
	"computer/logicgate"
	"computer/multiplexer"
)

type ALUInterface interface {
	// TODO: 値オブジェクトを入れたい
	Run(opecode, imm multiplexer.Bool4bit, register Register, in_a, in_b multiplexer.Bool4bit) Register
}

type ALU struct {
	AdderInterface adder.AdderInterface
}

type Register struct {
	A   multiplexer.Bool4bit
	B   multiplexer.Bool4bit
	C   multiplexer.Bool4bit
	IP  multiplexer.Bool4bit
	Out multiplexer.Bool4bit
	CF  bool
}

func (alu ALU) Run(opecode, imm multiplexer.Bool4bit, register Register, in_a, in_b multiplexer.Bool4bit) Register {
	result_add_a_b := alu.AdderInterface.Run(register.A, register.B)

	// bの2の補数を取って、a-bを計算する
	not_b := multiplexer.Bool4bit{
		B3: logicgate.NOT{A: register.B.B3}.Out(),
		B2: logicgate.NOT{A: register.B.B2}.Out(),
		B1: logicgate.NOT{A: register.B.B1}.Out(),
		B0: logicgate.NOT{A: register.B.B0}.Out(),
	}
	complement_b := alu.AdderInterface.Run(not_b, multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true})
	result_sub_a_b := alu.AdderInterface.Run(register.A, complement_b.Sum)

	// registerAの計算
	result_add_a_imm := alu.AdderInterface.Run(register.A, imm)
	next_a := multiplexer.Multiplexer16to1_4bit(
		result_add_a_imm.Sum, register.B, in_a, imm,
		register.A, register.A, register.A, register.A,
		result_add_a_b.Sum, register.A, register.A, register.A,
		register.C, result_sub_a_b.Sum, register.A, register.A,
		opecode,
	)
	// registerBの計算
	result_add_b_imm := alu.AdderInterface.Run(register.B, imm)
	next_b := multiplexer.Multiplexer16to1_4bit(
		register.B, register.B, register.B, register.B,
		register.A, result_add_b_imm.Sum, in_b, imm,
		register.B, register.B, register.B, register.B,
		register.B, register.B, register.B, register.B,
		opecode,
	)
	// registerCの計算
	next_c := multiplexer.Multiplexer16to1_4bit(
		register.C, register.C, register.C, register.C,
		register.C, register.C, register.C, register.C,
		register.C, register.C, register.A, register.C,
		register.C, register.C, register.C, register.C,
		opecode,
	)

	// キャリーフラグの計算
	next_cf := multiplexer.Multiplexer16to1(
		result_add_a_imm.Cf, register.CF, register.CF, register.CF,
		register.CF, result_add_b_imm.Cf, register.CF, register.CF,
		result_add_a_b.Cf, register.CF, register.CF, register.CF,
		register.CF, result_sub_a_b.Cf, register.CF, register.CF,
		opecode,
	)

	// 命令ポインタの計算
	result_add_ip_1 := alu.AdderInterface.Run(register.IP, multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: true})
	next_ip := multiplexer.Multiplexer16to1_4bit(
		result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum,
		result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum,
		result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum, result_add_ip_1.Sum,
		result_add_ip_1.Sum, result_add_ip_1.Sum, multiplexer.Multiplexer2to1_4bit(imm, result_add_ip_1.Sum, register.CF), imm,
		opecode,
	)
	// 出力の計算
	next_out := multiplexer.Multiplexer16to1_4bit(
		register.Out, register.Out, register.Out, register.Out,
		register.Out, register.Out, register.Out, register.Out,
		register.Out, register.B, register.Out, imm,
		register.Out, register.Out, register.Out, register.Out,
		opecode,
	)
	return Register{A: next_a, B: next_b, C: next_c, IP: next_ip, Out: next_out, CF: next_cf}
}
