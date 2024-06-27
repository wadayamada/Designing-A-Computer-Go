package cpu

import (
	"computer/alu"
	"computer/dff"
	"computer/multiplexer"
	"computer/rom"
)

type CPU struct {
	RegisterA    dff.DFF
	RegisterB    dff.DFF
	RegisterIP   dff.DFF
	RegisterOut  dff.DFF
	RegisterCF   dff.DFF
	ALUInterface alu.ALUInterface
}

type CPUInterface interface {
	Run(rom rom.RomInterface)
}

func (c *CPU) Write(a, b, ip, out, cf multiplexer.Bool4bit) {
	// TODO: registerに書き込む
}

func (c *CPU) Run(romInterface rom.RomInterface) {
	for {
		data := romInterface.Get(c.RegisterIP.D)
		// TODO: Bool8bitからBool4bit, Bool4bitに変換する
		opecode, imm := multiplexer.Bool4bit{B3: data.B7, B2: data.B6, B1: data.B5, B0: data.B4}, multiplexer.Bool4bit{B3: data.B3, B2: data.B2, B1: data.B1, B0: data.B0}
		a, b, ip, out, cf := c.ALUInterface.Run(opecode, imm, c.RegisterA.D, c.RegisterB.D, c.RegisterIP.D, c.RegisterOut.D, c.RegisterCF.D)
		c.Write(a, b, ip, out, cf)
	}
}
