package basicCalculator

import (
	"testing"
)

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "0",
			args:    args{s: "1+(2-3)"},
			wantAns: 0,
		},
		{
			name:    "1",
			args:    args{s: "2147483647"},
			wantAns: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := calculate(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("calculate() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
