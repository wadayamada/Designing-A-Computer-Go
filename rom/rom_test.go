package rom

import (
	"computer/multiplexer"
	"reflect"
	"testing"
)

func TestRom_Get(t *testing.T) {
	type fields struct {
		m0  multiplexer.Bool8bit
		m1  multiplexer.Bool8bit
		m2  multiplexer.Bool8bit
		m3  multiplexer.Bool8bit
		m4  multiplexer.Bool8bit
		m5  multiplexer.Bool8bit
		m6  multiplexer.Bool8bit
		m7  multiplexer.Bool8bit
		m8  multiplexer.Bool8bit
		m9  multiplexer.Bool8bit
		m10 multiplexer.Bool8bit
		m11 multiplexer.Bool8bit
		m12 multiplexer.Bool8bit
		m13 multiplexer.Bool8bit
		m14 multiplexer.Bool8bit
		m15 multiplexer.Bool8bit
	}
	type args struct {
		address multiplexer.Bool4bit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   multiplexer.Bool8bit
	}{
		{
			"Get 0",
			fields{
				m0: multiplexer.Bool8bit{
					B7: true, B6: true, B5: true, B4: true, B3: true, B2: true, B1: true, B0: true,
				},
				m1: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m2: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m3: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m4: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m5: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m6: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m7: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m8: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m9: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m10: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m11: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m12: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m13: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m14: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m15: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
			},
			args{
				address: multiplexer.Bool4bit{
					B3: false, B2: false, B1: false, B0: false,
				},
			},
			multiplexer.Bool8bit{
				B7: true, B6: true, B5: true, B4: true, B3: true, B2: true, B1: true, B0: true,
			},
		},
		{
			"Get 5",
			fields{
				m0: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m1: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m2: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m3: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m4: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m5: multiplexer.Bool8bit{
					B7: true, B6: true, B5: true, B4: true, B3: true, B2: true, B1: true, B0: true,
				},
				m6: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m7: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m8: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m9: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m10: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m11: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m12: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m13: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m14: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
				m15: multiplexer.Bool8bit{
					B7: false, B6: false, B5: false, B4: false, B3: false, B2: false, B1: false, B0: false,
				},
			},
			args{
				address: multiplexer.Bool4bit{
					B3: false, B2: true, B1: false, B0: true,
				},
			},
			multiplexer.Bool8bit{
				B7: true, B6: true, B5: true, B4: true, B3: true, B2: true, B1: true, B0: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rom{
				M0:  tt.fields.m0,
				M1:  tt.fields.m1,
				M2:  tt.fields.m2,
				M3:  tt.fields.m3,
				M4:  tt.fields.m4,
				M5:  tt.fields.m5,
				M6:  tt.fields.m6,
				M7:  tt.fields.m7,
				M8:  tt.fields.m8,
				M9:  tt.fields.m9,
				M10: tt.fields.m10,
				M11: tt.fields.m11,
				M12: tt.fields.m12,
				M13: tt.fields.m13,
				M14: tt.fields.m14,
				M15: tt.fields.m15,
			}
			if got := r.Get(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rom.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
