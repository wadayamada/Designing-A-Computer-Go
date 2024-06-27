package computer

import (
	"computer/cpu"
	"computer/rom"
)

type Computer struct {
	CPUInterface cpu.CPUInterface
	ROMInterface rom.RomInterface
}

func (c *Computer) Run() {
	c.CPUInterface.Run(c.ROMInterface)
}
