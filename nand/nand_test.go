package nand

import "testing"

func TestNAND_Out(t *testing.T) {
	type fields struct {
		A bool
		B bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"NAND true, true",
			fields{true, true},
			false,
		},
		{
			"NAND true, false",
			fields{true, false},
			true,
		},
		{
			"NAND false, true",
			fields{false, true},
			true,
		},
		{
			"NAND false, false",
			fields{false, false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NAND{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := n.Out(); got != tt.want {
				t.Errorf("NAND.Out() = %v, want %v", got, tt.want)
			}
		})
	}
}
