package day9

import (
	"testing"
)

func Test_execute(t *testing.T) {
	type args struct {
		data  []string
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				data:  []string{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"},
				input: "-1",
			},
			want: "99",
		},
		{
			name: "t2",
			args: args{
				data:  []string{"1102,34915192,34915192,7,4,7,99,0"},
				input: "-1",
			},
			want: "1219070632396864",
		},
		{
			name: "t3",
			args: args{
				data:  []string{"104,1125899906842624,99"},
				input: "-1",
			},
			want: "1125899906842624",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := execute(tt.args.data, tt.args.input); got != tt.want {
				t.Errorf("execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
