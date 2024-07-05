package rom

import "computer/multiplexer"

type Rom struct {
	M0, M1, M2, M3, M4, M5, M6, M7, M8, M9, M10, M11, M12, M13, M14, M15 multiplexer.Bool8bit
}

func (r Rom) Get(address multiplexer.Bool4bit) multiplexer.Bool8bit {
	return multiplexer.Multiplexer16to1_8bit(
		r.M0, r.M1, r.M2, r.M3, r.M4, r.M5, r.M6, r.M7, r.M8, r.M9, r.M10, r.M11, r.M12, r.M13, r.M14, r.M15,
		address,
	)
}

type RomInterface interface {
	Get(address multiplexer.Bool4bit) multiplexer.Bool8bit
}
