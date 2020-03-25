package day4

import (
	"testing"
)

func Test_hasAPairPart1(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{a: 115267},
			want: true,
		},
		{
			name: "t2",
			args: args{a: 715227},
			want: true,
		},
		{
			name: "t3",
			args: args{a: 715277},
			want: true,
		},
		{
			name: "t4",
			args: args{a: 71567},
			want: false,
		},
		{
			name: "t5",
			args: args{a: 111267},
			want: true,
		},
		{
			name: "t6",
			args: args{a: 712227},
			want: true,
		},
		{
			name: "t7",
			args: args{a: 715777},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAPairPart1(tt.args.a); got != tt.want {
				t.Errorf("hasAPairPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_hasAPairPart2(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "t1",
			args: args{a: 115267},
			want: true,
		},
		{
			name: "t2",
			args: args{a: 715227},
			want: true,
		},
		{
			name: "t3",
			args: args{a: 715277},
			want: true,
		},
		{
			name: "t4",
			args: args{a: 71567},
			want: false,
		},
		{
			name: "t5",
			args: args{a: 111267},
			want: false,
		},
		{
			name: "t6",
			args: args{a: 712227},
			want: false,
		},
		{
			name: "t7",
			args: args{a: 715777},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAPairPart2(tt.args.a); got != tt.want {
				t.Errorf("hasAPairPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_forceIncrease(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{a: 254032},
			want: 255555,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forceIncrease(tt.args.a); got != tt.want {
				t.Errorf("forceIncrease() = %v, want %v", got, tt.want)
			}
		})
	}
}
