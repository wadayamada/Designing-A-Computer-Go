package rom

import "computer/multiplexer"

type Rom struct {
	m0, m1, m2, m3, m4, m5, m6, m7, m8, m9, m10, m11, m12, m13, m14, m15 multiplexer.Bool8bit
}

func (r Rom) Get(address multiplexer.Bool4bit) multiplexer.Bool8bit {
	return multiplexer.Multiplexer16to1_8bit(
		r.m0, r.m1, r.m2, r.m3, r.m4, r.m5, r.m6, r.m7, r.m8, r.m9, r.m10, r.m11, r.m12, r.m13, r.m14, r.m15,
		address,
	)
}
