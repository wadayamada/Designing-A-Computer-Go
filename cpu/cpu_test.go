package cpu

import (
	"computer/adder"
	"computer/alu"
	"computer/dff"
	"computer/multiplexer"
	"computer/rom"
	"reflect"
	"testing"
)

func TestCPU_Run(t *testing.T) {
	type fields struct {
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
	type args struct {
		romInterface rom.RomInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   multiplexer.Bool4bit
	}{
		{
			"1+1を計算して、結果をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				InB:          multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					M0: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: true, B3: false, B2: false, B1: false, B0: true},   // MOV B, 1
					M1: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},  // ADD B, 1
					M2: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false}, // OUT B
				},
			},
			multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false},
		},
		{
			"16-4(IN_A)を計算して、結果の12をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: false}, // 4
				InB:          multiplexer.Bool4bit{B3: false, B2: false, B1: false, B0: false},
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					M0: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M1: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: true, B3: false, B2: false, B1: false, B0: false},   // MOV B, 0
					M2: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},   // ADD B, 1
					M3: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M4: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: false, B1: true, B0: false},    // JNC 2
					M5: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},  // OUT B
				},
			},
			multiplexer.Bool4bit{B3: true, B2: true, B1: false, B0: false}, // 12
		},
		{
			"拡張命令を使わずに、3(IN_A)+8(IN_B)を計算して、結果11をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				InB:          multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false}, // 8
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					// まず16-3(IN_A)を計算して、結果13をAに書き込む
					M0: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M1: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: true, B3: false, B2: false, B1: false, B0: false},   // MOV B, 0
					M2: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},   // ADD B, 1
					M3: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M4: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: false, B1: true, B0: false},    // JNC 2
					M5: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false}, // MOV A, B
					// 13(A), 8(B)に、Aが16になるまで1を足すことで、8(B)+3になって、結果11をOUTする
					M6:  multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					M7:  multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},   // ADD B, 1
					M8:  multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M9:  multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: true, B1: true, B0: true},      // JNC 7
					M10: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},  // OUT B
				},
			},
			multiplexer.Bool4bit{B3: true, B2: false, B1: true, B0: true}, // 11
		},
		{
			"拡張命令ADD A, Bを使って、3(IN_A)+8(IN_B)を計算して、結果11をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				InB:          multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false}, // 8
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					// まず16-3(IN_A)を計算して、結果13をAに書き込む
					M0: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M1: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					M2: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // ADD A, B
					M3: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					M4: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},  // OUT B
				},
			},
			multiplexer.Bool4bit{B3: true, B2: false, B1: true, B0: true}, // 11
		},
		{
			"8(IN_A)-3(IN_B)を計算して、結果5をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false}, // 8
				InB:          multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					// まず16-3(IN_B)を計算して、結果13をBに書き込む
					M0: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: true, B3: false, B2: false, B1: false, B0: false},  // MOV A, 0
					M1: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					M2: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M3: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},   // ADD B, 1
					M4: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: false, B1: true, B0: false},    // JNC 2
					M5: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					// 8(IN_A)+13(IN_B)を計算して、結果5(21)をOUTする
					M6: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M7: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // ADD A, B
					M8: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					M9: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},  // OUT B
				},
			},
			multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: true}, // 5
		},
		{
			"4(IN_A)*3(IN_B)を計算して、結果12をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: false, B2: true, B1: false, B0: false}, // 4
				InB:          multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					// まず16-3(IN_B)を計算して、結果13をBに書き込む
					M0: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					M1: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M2: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: true},   // ADD B, 1
					M3: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: true},    // JNC 1
					M4: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: true, B3: false, B2: false, B1: false, B0: false},   // MOV B, 0
					// 4(IN_A)を3(IN_B)回足して、結果12をBに書き込む
					M5:  multiplexer.Bool8bit{B7: true, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV C, A
					M6:  multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M7:  multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // ADD A, B
					M8:  multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					M9:  multiplexer.Bool8bit{B7: true, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV A, C
					M10: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M11: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: true, B1: false, B0: true},     // JNC 5
					// OUT B
					M12: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false}, // OUT B
				},
			},
			multiplexer.Bool4bit{B3: true, B2: true, B1: false, B0: false}, // 12
		},
		{
			"8(IN_A)/3(IN_B)を計算して、結果2をOUTする",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterC:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterIP:   dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterOut:  dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterCF:   dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}},
				InA:          multiplexer.Bool4bit{B3: true, B2: false, B1: false, B0: false}, // 8
				InB:          multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: true},  // 3
				ALUInterface: alu.ALU{AdderInterface: adder.Adder{HalfAdderInterface: adder.HalfAdder{}, FullAdderInterface: adder.FullAdder{HalfAdderInterface: adder.HalfAdder{}}}},
			},
			args{
				rom.Rom{
					// IN A, IN Bと、Cの初期化
					M0: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: true, B3: false, B2: false, B1: false, B0: false},  // MOV A, 0
					M1: multiplexer.Bool8bit{B7: true, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV C, A
					M2: multiplexer.Bool8bit{B7: false, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false}, // IN A
					M3: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					// SUB A, Bと、カウンタCをインクリメントを繰り返す
					M4:  multiplexer.Bool8bit{B7: true, B6: true, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},   // SUB A, B
					M5:  multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: true, B2: true, B1: false, B0: true},      // JNC 13
					M6:  multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					M7:  multiplexer.Bool8bit{B7: true, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV A, C
					M8:  multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: true}, // ADD A, 1
					M9:  multiplexer.Bool8bit{B7: true, B6: false, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV C, A
					M10: multiplexer.Bool8bit{B7: false, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false}, // MOV A, B
					M11: multiplexer.Bool8bit{B7: false, B6: true, B5: true, B4: false, B3: false, B2: false, B1: false, B0: false},  // IN B
					M12: multiplexer.Bool8bit{B7: true, B6: true, B5: true, B4: false, B3: false, B2: true, B1: false, B0: false},    // JNC 4
					// CをOUTする
					M13: multiplexer.Bool8bit{B7: true, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false},  // MOV A, C
					M14: multiplexer.Bool8bit{B7: false, B6: true, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false}, // MOV B, A
					M15: multiplexer.Bool8bit{B7: true, B6: false, B5: false, B4: true, B3: false, B2: false, B1: false, B0: false},  // OUT B
				},
			},
			multiplexer.Bool4bit{B3: false, B2: false, B1: true, B0: false}, // 2
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CPU{
				RegisterA:    tt.fields.RegisterA,
				RegisterB:    tt.fields.RegisterB,
				RegisterC:    tt.fields.RegisterC,
				RegisterIP:   tt.fields.RegisterIP,
				RegisterOut:  tt.fields.RegisterOut,
				RegisterCF:   tt.fields.RegisterCF,
				InA:          tt.fields.InA,
				InB:          tt.fields.InB,
				ALUInterface: tt.fields.ALUInterface,
			}
			if got := c.Run(tt.args.romInterface); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CPU.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
