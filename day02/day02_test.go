package day02

import "testing"

func Test_testProgram(t *testing.T) {
	type args struct {
		lines []string
		in1   int
		in2   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{lines: []string{"1,0,0,0,99"}},
			want: "2",
		},
		{
			name: "test2",
			args: args{lines: []string{"2,3,0,3,99"}},
			want: "2",
		},
		{
			name: "test3",
			args: args{lines: []string{"2,4,4,5,99,0"}},
			want: "2",
		},
		{
			name: "test4",
			args: args{lines: []string{"1,1,1,4,99,5,6,0,99"}},
			want: "30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testProgram(tt.args.lines, tt.args.in1, tt.args.in2); got != tt.want {
				t.Errorf("testProgram() = %v, want %v", got, tt.want)
			}
		})
	}
}
