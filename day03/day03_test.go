package day03

import (
	"testing"
)

func Test_solvePart1(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83"}},
			want: 159,
		},
		{
			name: "test2",
			args: args{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}},
			want: 135,
		},
		{
			name: "test3",
			args: args{[]string{"R8,U5,L5,D3",
				"U7,R6,D4,L4"}},
			want: 6,
		},
		{
			name: "test4",
			args: args{[]string{"R10,U3,L5,D20",
				"D4,R10"}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.s); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83"}},
			want: 610,
		},
		{
			name: "test2",
			args: args{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}},
			want: 410,
		},
		{
			name: "test3",
			args: args{[]string{"R10,U3,L5,D20",
				"D4,R10"}},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.s); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
