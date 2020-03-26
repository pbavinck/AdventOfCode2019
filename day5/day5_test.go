package day5

import (
	"testing"
)

func Test_padWithZeros(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{s: "1001"},
			want: "01001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padWithZeros(tt.args.s); got != tt.want {
				t.Errorf("padWithZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paramMode(t *testing.T) {
	type args struct {
		opcode string
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{opcode: "101", index: 0},
			want: "1",
		},
		{
			name: "t2",
			args: args{opcode: "1001", index: 0},
			want: "0",
		},
		{
			name: "t3",
			args: args{opcode: "01", index: 2},
			want: "0",
		},
		{
			name: "t4",
			args: args{opcode: "1001", index: 1},
			want: "1",
		},
		{
			name: "t2",
			args: args{opcode: "10101", index: 2},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paramMode(tt.args.opcode, tt.args.index); got != tt.want {
				t.Errorf("paramMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueToUse(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
		paramIndex  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				code:       []string{"1", "0", "5", "4", "", "55"},
				paramIndex: 0,
			},
			want: 1,
		},
		{
			name: "t2",
			args: args{
				code:       []string{"00001", "0", "5", "4", "", "55"},
				paramIndex: 1,
			},
			want: 55,
		},
		{
			name: "t3",
			args: args{
				code:       []string{"01001", "0", "5", "4", "", "55"},
				paramIndex: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valueToUse(tt.args.code, tt.args.opcodeIndex, tt.args.paramIndex); got != tt.want {
				t.Errorf("valueToUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doInstr1(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				code:        []string{"1", "3", "5", "4", "", "55"},
				opcodeIndex: 0,
			},
			want: []string{"1", "3", "5", "4", "59", "55"},
		},
		{
			name: "t2",
			args: args{
				code:        []string{"1001", "3", "5", "4", "", "55"},
				opcodeIndex: 0,
			},
			want: []string{"1001", "3", "5", "4", "9", "55"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doInstr1(tt.args.code, tt.args.opcodeIndex)
			for i := 0; i < len(tt.args.code); i++ {
				if tt.want[i] != tt.args.code[i] {
					t.Errorf("doInstr1() = %v, want %v", tt.args.code, tt.want)
					break
				}
			}
		})
	}
}

func Test_doInstr2(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				code:        []string{"2", "3", "5", "4", "", "55"},
				opcodeIndex: 0,
			},
			want: []string{"2", "3", "5", "4", "220", "55"},
		},
		{
			name: "t2",
			args: args{
				code:        []string{"1002", "3", "5", "4", "", "55"},
				opcodeIndex: 0,
			},
			want: []string{"1002", "3", "5", "4", "20", "55"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doInstr2(tt.args.code, tt.args.opcodeIndex)
			for i := 0; i < len(tt.args.code); i++ {
				if tt.want[i] != tt.args.code[i] {
					t.Errorf("doInstr2() = %v, want %v", tt.args.code, tt.want)
					break
				}
			}
		})
	}
}

func Test_doInstr3(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
		input       string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t1",
			args: args{
				code:        []string{"3", "0"},
				opcodeIndex: 0,
				input:       "6",
			},
			want: []string{"6", "0"},
		},
		{
			name: "t2",
			args: args{
				code:        []string{"3", "4", "5", "4", "", "55"},
				opcodeIndex: 0,
				input:       "6",
			},
			want: []string{"3", "4", "5", "4", "6", "55"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doInstr3(tt.args.code, tt.args.opcodeIndex, tt.args.input)
			for i := 0; i < len(tt.args.code); i++ {
				if tt.want[i] != tt.args.code[i] {
					t.Errorf("doInstr3() = %v, want %v", tt.args.code, tt.want)
					break
				}
			}
		})
	}
}

func Test_doInstr5(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name                string
		args                args
		wantNextOpcodeIndex int
	}{
		{
			name: "t1 position in JUMP",
			args: args{
				code:        []string{"5", "4", "4", "3", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 18,
		},
		{
			name: "t2 position in no jump",
			args: args{
				code:        []string{"5", "3", "4", "0", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 3,
		},
		{
			name: "t3 immediate in JUMP",
			args: args{
				code:        []string{"105", "1", "4", "3", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 18,
		},
		{
			name: "t4 immediate in no jump",
			args: args{
				code:        []string{"105", "0", "4", "0", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 3,
		},
		{
			name: "t5 immediate out JUMP",
			args: args{
				code:        []string{"1005", "4", "4", "3", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNextOpcodeIndex := doInstr5(tt.args.code, tt.args.opcodeIndex); gotNextOpcodeIndex != tt.wantNextOpcodeIndex {
				t.Errorf("doInstr5() = %v, want %v", gotNextOpcodeIndex, tt.wantNextOpcodeIndex)
			}
		})
	}
}

func Test_doInstr6(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name                string
		args                args
		wantNextOpcodeIndex int
	}{
		{
			name: "t1 position in JUMP",
			args: args{
				code:        []string{"5", "4", "3", "18", "0"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 18,
		},
		{
			name: "t2 position in no jump",
			args: args{
				code:        []string{"5", "3", "4", "12", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 3,
		},
		{
			name: "t3 immediate in JUMP",
			args: args{
				code:        []string{"105", "0", "4", "3", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 18,
		},
		{
			name: "t4 immediate in no jump",
			args: args{
				code:        []string{"105", "1", "4", "0", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 3,
		},
		{
			name: "t5 immediate out JUMP",
			args: args{
				code:        []string{"1005", "3", "4", "0", "18"},
				opcodeIndex: 0,
			},
			wantNextOpcodeIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNextOpcodeIndex := doInstr6(tt.args.code, tt.args.opcodeIndex); gotNextOpcodeIndex != tt.wantNextOpcodeIndex {
				t.Errorf("doInstr6() = %v, want %v", gotNextOpcodeIndex, tt.wantNextOpcodeIndex)
			}
		})
	}
}

func Test_doInstr7(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name                string
		args                args
		want                []string
		wantNextOpcodeIndex int
	}{
		{
			name: "t1 position in SET",
			args: args{
				code:        []string{"7", "4", "5", "6", "12", "21", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"7", "4", "5", "6", "12", "21", "1"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t1 position in NOT SET",
			args: args{
				code:        []string{"7", "4", "5", "6", "21", "12", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"7", "4", "5", "6", "21", "12", "0"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t1 immediate in SET",
			args: args{
				code:        []string{"107", "20", "5", "6", "12", "21", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"107", "20", "5", "6", "12", "21", "1"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t3 immediate in NOT SET",
			args: args{
				code:        []string{"107", "88", "5", "6", "21", "12", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"107", "88", "5", "6", "21", "12", "0"},
			wantNextOpcodeIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doInstr7(tt.args.code, tt.args.opcodeIndex)
			for i := 0; i < len(tt.args.code); i++ {
				if tt.want[i] != tt.args.code[i] {
					t.Errorf("doInstr7() = %v, want %v", tt.args.code, tt.want)
					break
				}
			}
			if got != tt.wantNextOpcodeIndex {
				t.Errorf("doInstr7() = %v, want %v", got, tt.wantNextOpcodeIndex)
			}
		})
	}
}

func Test_doInstr8(t *testing.T) {
	type args struct {
		code        []string
		opcodeIndex int
	}
	tests := []struct {
		name                string
		args                args
		want                []string
		wantNextOpcodeIndex int
	}{
		{
			name: "t1 position in SET",
			args: args{
				code:        []string{"7", "4", "5", "6", "12", "12", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"7", "4", "5", "6", "12", "12", "1"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t1 position in NOT SET",
			args: args{
				code:        []string{"7", "4", "5", "6", "21", "12", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"7", "4", "5", "6", "21", "12", "0"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t1 immediate in SET",
			args: args{
				code:        []string{"107", "20", "5", "6", "12", "20", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"107", "20", "5", "6", "12", "20", "1"},
			wantNextOpcodeIndex: 4,
		},
		{
			name: "t3 immediate in NOT SET",
			args: args{
				code:        []string{"107", "88", "5", "6", "21", "12", "67"},
				opcodeIndex: 0,
			},
			want:                []string{"107", "88", "5", "6", "21", "12", "0"},
			wantNextOpcodeIndex: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doInstr8(tt.args.code, tt.args.opcodeIndex)
			for i := 0; i < len(tt.args.code); i++ {
				if tt.want[i] != tt.args.code[i] {
					t.Errorf("doInstr8() = %v, want %v", tt.args.code, tt.want)
					break
				}
			}
			if got != tt.wantNextOpcodeIndex {
				t.Errorf("doInstr8() = %v, want %v", got, tt.wantNextOpcodeIndex)
			}
		})
	}
}
