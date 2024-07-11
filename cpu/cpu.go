package cpu

import (
	"computer/alu"
	"computer/dff"
	"computer/multiplexer"
	"computer/rom"
)

type CPU struct {
	RegisterA    dff.DFF4bitInterface
	RegisterB    dff.DFF4bitInterface
	RegisterC    dff.DFF4bitInterface
	RegisterIP   dff.DFF4bitInterface
	RegisterOut  dff.DFF4bitInterface
	RegisterCF   dff.DFFInterface
	InA          multiplexer.Bool4bit
	InB          multiplexer.Bool4bit
	ALUInterface alu.ALUInterface
}

type CPUInterface interface {
	Run(rom rom.RomInterface)
}

func (cpu *CPU) Write(a, b, c, ip, out multiplexer.Bool4bit, cf bool) {
	cpu.RegisterA.Write(a)
	cpu.RegisterB.Write(b)
	cpu.RegisterC.Write(c)
	cpu.RegisterIP.Write(ip)
	cpu.RegisterOut.Write(out)
	cpu.RegisterCF.Write(cf)
}

func (c *CPU) Run(romInterface rom.RomInterface) multiplexer.Bool4bit {
	for {
		data := romInterface.Get(c.RegisterIP.Read())
		opecode := multiplexer.Bool4bit{B3: data.B7, B2: data.B6, B1: data.B5, B0: data.B4}
		imm := multiplexer.Bool4bit{B3: data.B3, B2: data.B2, B1: data.B1, B0: data.B0}
		next_register := c.ALUInterface.Run(opecode, imm, alu.Register{A: c.RegisterA.Read(), B: c.RegisterB.Read(), IP: c.RegisterIP.Read(), Out: c.RegisterOut.Read(), CF: c.RegisterCF.Read()}, c.InA, c.InB)
		c.Write(next_register.A, next_register.B, next_register.C, next_register.IP, next_register.Out, next_register.CF)
		// TODO: OUT Bの命令実行したら終了するようにしたが、いい感じにしたい
		if (opecode == multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: true}) {
			return next_register.Out
		}
	}
}
