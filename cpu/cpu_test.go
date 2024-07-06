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
			"CPU Run",
			fields{
				RegisterA:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
				RegisterB:    dff.DFF4bit{DFF0: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF1: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF2: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}, DFF3: dff.DFF{RSFFInterface: &dff.RSFF{Q: false, Q_not: true}}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CPU{
				RegisterA:    tt.fields.RegisterA,
				RegisterB:    tt.fields.RegisterB,
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
