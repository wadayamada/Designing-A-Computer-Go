package dff

import "testing"

func TestRSFF_Run(t *testing.T) {
	type fields struct {
		Q     bool
		Q_not bool
	}
	type args struct {
		Reset bool
		Set   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"RSFF false, false",
			fields{Q: false, Q_not: true},
			args{false, false},
			false,
			false,
		},
		{
			"RSFF false, false",
			fields{Q: true, Q_not: false},
			args{false, false},
			true,
			false,
		},
		{
			"RSFF false, true",
			fields{Q: false, Q_not: true},
			args{false, true},
			true,
			false,
		},
		{
			"RSFF false, true",
			fields{Q: true, Q_not: false},
			args{false, true},
			true,
			false,
		},
		{
			"RSFF true, false",
			fields{Q: false, Q_not: true},
			args{true, false},
			false,
			false,
		},
		{
			"RSFF true, false",
			fields{Q: true, Q_not: false},
			args{true, false},
			false,
			false,
		},
		{
			"RSFF true, true",
			fields{Q: false, Q_not: true},
			args{true, true},
			false,
			true,
		},
		{
			"RSFF true, true",
			fields{Q: true, Q_not: false},
			args{true, true},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsff := RSFF{
				Q:     tt.fields.Q,
				Q_not: tt.fields.Q_not,
			}
			got, err := rsff.Run(tt.args.Reset, tt.args.Set)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSFF.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RSFF.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
