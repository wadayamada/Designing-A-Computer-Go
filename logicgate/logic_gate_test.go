package logicgate

import "testing"

func TestNOT_Out(t *testing.T) {
	type fields struct {
		A bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"NOT true",
			fields{true},
			false,
		},
		{
			"NOT false",
			fields{false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NOT{
				A: tt.fields.A,
			}
			if got := n.Out(); got != tt.want {
				t.Errorf("NOT.Out() = %v, want %v", got, tt.want)
			}
		})
	}
}
