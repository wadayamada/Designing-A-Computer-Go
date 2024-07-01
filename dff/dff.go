package dff

import "computer/multiplexer"

// TODO: DFF4bitをDFFでちゃんと実装する
type DFF4bit struct {
	D multiplexer.Bool4bit
	Q multiplexer.Bool4bit
}

type DFF4bitInterface interface {
	Run()
	Read() multiplexer.Bool4bit
	Write(multiplexer.Bool4bit)
}

// TODO: DFFを論理ゲートでちゃんと実装する
type DFF struct {
	D bool
	Q bool
}

type DFFInterface interface {
	Run()
	Read() bool
	Write(bool)
}
