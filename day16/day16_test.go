package day16

import (
	"testing"
)

func Test_calcDigit(t *testing.T) {
	type args struct {
		signal     []int
		digitIndex int
		pattern    []int
	}
	tests := []struct {
		name      string
		args      args
		wantDigit int
	}{
		{
			name: "digit 1",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 0,
			},
			wantDigit: 4,
		},
		{
			name: "digit 2",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 1,
			},
			wantDigit: 8,
		},
		{
			name: "digit 3",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 2,
			},
			wantDigit: 2,
		},
		{
			name: "digit 4",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 3,
			},
			wantDigit: 2,
		},
		{
			name: "digit 5",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 4,
			},
			wantDigit: 6,
		},
		{
			name: "digit 6",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 5,
			},
			wantDigit: 1,
		},
		{
			name: "digit 7",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 6,
			},
			wantDigit: 5,
		},
		{
			name: "digit 8",
			args: args{
				signal:     []int{1, 2, 3, 4, 5, 6, 7, 8},
				digitIndex: 7,
			},
			wantDigit: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.args.pattern = createPattern(tt.args.digitIndex+1, len(tt.args.signal)) //[]int{0, 1, 0, -1, 0}
			gotDigit := calcDigit(tt.args.signal, tt.args.digitIndex) //), tt.args.pattern)
			if gotDigit != tt.wantDigit {
				t.Errorf("calcDigit() gotDigit = %v, want %v", gotDigit, tt.wantDigit)
			}
		})
	}
}
