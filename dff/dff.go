package dff

import "computer/multiplexer"

type DFF struct {
	D multiplexer.Bool4bit
	Q multiplexer.Bool4bit
}
