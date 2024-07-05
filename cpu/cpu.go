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
	RegisterIP   dff.DFF4bitInterface
	RegisterOut  dff.DFF4bitInterface
	RegisterCF   dff.DFFInterface
	In           multiplexer.Bool4bit
	ALUInterface alu.ALUInterface
}

type CPUInterface interface {
	Run(rom rom.RomInterface)
}

func (c *CPU) Write(a, b, ip, out multiplexer.Bool4bit, cf bool) {
	c.RegisterA.Write(a)
	c.RegisterB.Write(b)
	c.RegisterIP.Write(ip)
	c.RegisterOut.Write(out)
	c.RegisterCF.Write(cf)
}

// TODO: テストも実装する
func (c *CPU) Run(romInterface rom.RomInterface) multiplexer.Bool4bit {
	for {
		data := romInterface.Get(c.RegisterIP.Read())
		opecode := multiplexer.Bool4bit{B3: data.B7, B2: data.B6, B1: data.B5, B0: data.B4}
		imm := multiplexer.Bool4bit{B3: data.B3, B2: data.B2, B1: data.B1, B0: data.B0}
		next_register := c.ALUInterface.Run(opecode, imm, alu.Register{A: c.RegisterA.Read(), B: c.RegisterB.Read(), IP: c.RegisterIP.Read(), Out: c.RegisterOut.Read(), CF: c.RegisterCF.Read()}, c.In)
		c.Write(next_register.A, next_register.B, next_register.IP, next_register.Out, next_register.CF)
		if (opecode == multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: true}) {
			return next_register.Out
		}
	}
}
