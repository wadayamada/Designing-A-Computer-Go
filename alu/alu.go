package alu

import "computer/multiplexer"

type ALUInterface interface {
	Run(opecode, imm, a, b, ip, out, cf multiplexer.Bool4bit) (multiplexer.Bool4bit, multiplexer.Bool4bit, multiplexer.Bool4bit, multiplexer.Bool4bit, multiplexer.Bool4bit)
}

// TODO: Runの実装
