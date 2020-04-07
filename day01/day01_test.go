package day01

import "testing"

func TestSolvePart1(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: []int{12}},
			want: 2,
		},
		{
			name: "test2",
			args: args{s: []int{14}},
			want: 2,
		},
		{
			name: "test3",
			args: args{s: []int{1969}},
			want: 654,
		},
		{
			name: "test4",
			args: args{s: []int{100756}},
			want: 33583,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart1(tt.args.s); got != tt.want {
				t.Errorf("SolvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{data: []int{12}},
			want: 2,
		},
		{
			name: "test2",
			args: args{data: []int{14}},
			want: 2,
		},
		{
			name: "test3",
			args: args{data: []int{1969}},
			want: 966,
		},
		{
			name: "test4",
			args: args{data: []int{100756}},
			want: 50346,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart2(tt.args.data); got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
