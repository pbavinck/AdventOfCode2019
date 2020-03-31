package machines

import (
	"testing"
)

func Test_padWithZeros(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantCode string
	}{
		{
			name:     "t1",
			args:     args{s: "1001"},
			wantCode: "01001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := padWithZeros(tt.args.s); got != tt.wantCode {
				t.Errorf("padWithZeros() = %v, want %v", got, tt.wantCode)
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
		name     string
		args     args
		wantCode string
	}{
		{
			name:     "t1",
			args:     args{opcode: "101", index: 0},
			wantCode: "1",
		},
		{
			name:     "t2",
			args:     args{opcode: "1001", index: 0},
			wantCode: "0",
		},
		{
			name:     "t3",
			args:     args{opcode: "01", index: 2},
			wantCode: "0",
		},
		{
			name:     "t4",
			args:     args{opcode: "1001", index: 1},
			wantCode: "1",
		},
		{
			name:     "t5",
			args:     args{opcode: "10101", index: 2},
			wantCode: "1",
		},
		{
			name:     "t6",
			args:     args{opcode: "20101", index: 2},
			wantCode: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paramMode(tt.args.opcode, tt.args.index); got != tt.wantCode {
				t.Errorf("paramMode() = %v, want %v", got, tt.wantCode)
			}
		})
	}
}

func TestComputer_getParamValue(t *testing.T) {
	type fields struct {
		name         string
		program      []string
		relativeBase int
		Input        chan string
		Output       chan string
	}
	type args struct {
		line       int
		paramIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "t1",
			fields: fields{
				program:      []string{"1", "0", "5", "4", "0", "55"},
				relativeBase: 200,
			},
			args: args{
				paramIndex: 0,
			},
			want: 1,
		},
		{
			name: "t2",
			fields: fields{
				program:      []string{"01001", "0", "5", "4", "0", "55"},
				relativeBase: 200,
			},
			args: args{
				paramIndex: 1,
			},
			want: 5,
		},
		{
			name: "t3",
			fields: fields{
				program:      []string{"00001", "0", "5", "4", "0", "55"},
				relativeBase: 200,
			},
			args: args{
				paramIndex: 1,
			},
			want: 55,
		},
		{
			name: "t4",
			fields: fields{
				program:      []string{"201", "2", "5", "4", "0", "55"},
				relativeBase: 3,
			},
			args: args{
				paramIndex: 0,
			},
			want: 55,
		},
		{
			name: "t4",
			fields: fields{
				program:      []string{"2201", "2", "-5", "4", "0", "55"},
				relativeBase: 8,
			},
			args: args{
				paramIndex: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				name:         tt.fields.name,
				program:      tt.fields.program,
				relativeBase: tt.fields.relativeBase,
				Input:        tt.fields.Input,
				Output:       tt.fields.Output,
			}
			if got := c.getParamValue(tt.args.line, tt.args.paramIndex); got != tt.want {
				t.Errorf("Computer.getParamValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantCode     []string
		wantNextline int
	}{
		{
			name: "t1",
			code: "1,3,5,4,0,55",
			args: args{
				line: 0,
			},
			wantCode:     []string{"1", "3", "5", "4", "59", "55"},
			wantNextline: 4,
		},
		{
			name: "t2",
			code: "1001,3,5,4,0,55",
			args: args{
				line: 0,
			},
			wantCode:     []string{"1001", "3", "5", "4", "9", "55"},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			got := c.add(tt.args.line)
			if got != tt.wantNextline {
				t.Errorf("add() = %v, want %v", got, tt.wantNextline)
			}
			for i := 0; i < len(c.program); i++ {
				if tt.wantCode[i] != c.program[i] {
					t.Errorf("add() = %v, want %v", c.program, tt.wantCode)
					break
				}
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantCode     []string
		wantNextline int
	}{
		{
			name: "t1",
			code: "2,3,5,4,0,55",
			args: args{
				line: 0,
			},
			wantCode:     []string{"2", "3", "5", "4", "220", "55"},
			wantNextline: 4,
		},
		{
			name: "t2",
			code: "1002,3,5,4,0,55",
			args: args{
				line: 0,
			},
			wantCode:     []string{"1002", "3", "5", "4", "20", "55"},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			got := c.multiply(tt.args.line)
			if got != tt.wantNextline {
				t.Errorf("amultiplyd() = %v, want %v", got, tt.wantNextline)
			}
			for i := 0; i < len(c.program[i]); i++ {
				if tt.wantCode[i] != c.program[i] {
					t.Errorf("multiply() = %v, want %v", tt.code, tt.wantCode)
					break
				}
			}
		})
	}
}

func Test_in(t *testing.T) {
	type args struct {
		line  int
		input string
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantCode     []string
		wantNextline int
	}{
		{
			code: "3,0",
			name: "t1",
			args: args{
				line:  0,
				input: "6",
			},
			wantCode:     []string{"6", "0"},
			wantNextline: 2,
		},
		{
			name: "t2",
			code: "3,4,5,4,0,55",
			args: args{
				line:  0,
				input: "6",
			},
			wantCode:     []string{"3", "4", "5", "4", "6", "55"},
			wantNextline: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			got := c.in(tt.args.line, tt.args.input)
			if got != tt.wantNextline {
				t.Errorf("amultiplyd() = %v, want %v", got, tt.wantNextline)
			}
			for i := 0; i < len(c.program); i++ {
				if tt.wantCode[i] != c.program[i] {
					t.Errorf("doInstr3() = %v, want %v", tt.code, tt.wantCode)
					break
				}
			}
		})
	}
}

func Test_out(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantNextline int
		wantOutput   int
	}{
		{
			name: "t1 position lookup",
			code: "004,2,16",
			args: args{
				line: 0,
			},
			wantNextline: 2,
			wantOutput:   16,
		},
		{
			name: "t1 immediate lookup",
			code: "104,2,16",
			args: args{
				line: 0,
			},
			wantNextline: 2,
			wantOutput:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			gotNextline, gotOutput := c.out(tt.args.line)
			if gotNextline != tt.wantNextline {
				t.Errorf("Computer.out() gotNextline = %v, want %v", gotNextline, tt.wantNextline)
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("Computer.out() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func Test_jumpIfTrue(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantNextline int
	}{
		{
			name: "t1 position in JUMP",
			code: "5,4,4,3,18",
			args: args{
				line: 0,
			},
			wantNextline: 18,
		},
		{
			name: "t2 position in no jump",
			code: "5,3,4,0,18",
			args: args{
				line: 0,
			},
			wantNextline: 3,
		},
		{
			name: "t3 immediate in JUMP",
			code: "105,1,4,3,18",
			args: args{
				line: 0,
			},
			wantNextline: 18,
		},
		{
			name: "t4 immediate in no jump",
			code: "105,0,4,0,18",
			args: args{
				line: 0,
			},
			wantNextline: 3,
		},
		{
			name: "t5 immediate out JUMP",
			code: "1005,4,4,3,18",
			args: args{
				line: 0,
			},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			if gotNextOpcodeIndex := c.jumpIfTrue(tt.args.line); gotNextOpcodeIndex != tt.wantNextline {
				t.Errorf("jumpIfTrue() = %v, want %v", gotNextOpcodeIndex, tt.wantNextline)
			}
		})
	}
}

func Test_jumpIfFalse(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantNextline int
	}{
		{
			name: "t1 position in JUMP",
			code: "5,4,3,18,0",
			args: args{
				line: 0,
			},
			wantNextline: 18,
		},
		{
			name: "t2 position in no jump",
			code: "5,3,4,12,18",
			args: args{
				line: 0,
			},
			wantNextline: 3,
		},
		{
			name: "t3 immediate in JUMP",
			code: "105,0,4,3,18",
			args: args{
				line: 0,
			},
			wantNextline: 18,
		},
		{
			name: "t4 immediate in no jump",
			code: "105,1,4,0,18",
			args: args{
				line: 0,
			},
			wantNextline: 3,
		},
		{
			name: "t5 immediate out JUMP",
			code: "1005,3,4,0,18",
			args: args{
				line: 0,
			},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			if gotNextOpcodeIndex := c.jumpIfFalse(tt.args.line); gotNextOpcodeIndex != tt.wantNextline {
				t.Errorf("jumpIfFalse() = %v, want %v", gotNextOpcodeIndex, tt.wantNextline)
			}
		})
	}
}

func Test_lessThan(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantCode     []string
		wantNextline int
	}{
		{
			name: "t1 position in SET",
			code: "7,4,5,6,12,21,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"7", "4", "5", "6", "12", "21", "1"},
			wantNextline: 4,
		},
		{
			name: "t1 position in NOT SET",
			code: "7,4,5,6,21,12,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"7", "4", "5", "6", "21", "12", "0"},
			wantNextline: 4,
		},
		{
			name: "t1 immediate in SET",
			code: "107,20,5,6,12,21,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"107", "20", "5", "6", "12", "21", "1"},
			wantNextline: 4,
		},
		{
			name: "t3 immediate in NOT SET",
			code: "107,88,5,6,21,12,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"107", "88", "5", "6", "21", "12", "0"},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			got := c.lessThan(tt.args.line)
			if got != tt.wantNextline {
				t.Errorf("lessThan() = %v, want %v", got, tt.wantNextline)
			}
			for i := 0; i < len(c.program); i++ {
				if tt.wantCode[i] != c.program[i] {
					t.Errorf("lessThan() = %v, want %v", tt.code, tt.wantCode)
					break
				}
			}
		})
	}
}

func Test_equal(t *testing.T) {
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		code         string
		args         args
		wantCode     []string
		wantNextline int
	}{
		{
			name: "t1 position in SET",
			code: "7,4,5,6,12,12,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"7", "4", "5", "6", "12", "12", "1"},
			wantNextline: 4,
		},
		{
			name: "t1 position in NOT SET",
			code: "7,4,5,6,21,12,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"7", "4", "5", "6", "21", "12", "0"},
			wantNextline: 4,
		},
		{
			name: "t1 immediate in SET",
			code: "107,20,5,6,12,20,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"107", "20", "5", "6", "12", "20", "1"},
			wantNextline: 4,
		},
		{
			name: "t3 immediate in NOT SET",
			code: "107,88,5,6,21,12,67",
			args: args{
				line: 0,
			},
			wantCode:     []string{"107", "88", "5", "6", "21", "12", "0"},
			wantNextline: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComputer("Tester", []string{tt.code}, 0)
			got := c.equal(tt.args.line)
			if got != tt.wantNextline {
				t.Errorf("equal() = %v, want %v", got, tt.wantNextline)
			}
			for i := 0; i < len(c.program); i++ {
				if tt.wantCode[i] != c.program[i] {
					t.Errorf("equal() = %v, want %v", tt.code, tt.wantCode)
					break
				}
			}
		})
	}
}

func TestComputer_base(t *testing.T) {
	type fields struct {
		name         string
		program      []string
		relativeBase int
		Input        chan string
		Output       chan string
	}
	type args struct {
		line int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNextLine int
		wantNextBase int
	}{
		{
			name: "t1",
			fields: fields{
				name:         "T",
				program:      []string{"9", "3", "5", "6", "21", "12", "67"},
				relativeBase: 200,
			},
			args: args{
				line: 0,
			},
			wantNextLine: 2,
			wantNextBase: 206,
		},
		{
			name: "t2",
			fields: fields{
				name:         "T",
				program:      []string{"109", "14", "5", "6", "21", "12", "67"},
				relativeBase: 200,
			},
			args: args{
				line: 0,
			},
			wantNextLine: 2,
			wantNextBase: 214,
		},
		{
			name: "t2",
			fields: fields{
				name:         "T",
				program:      []string{"209", "-196", "5", "6", "21", "12", "67"},
				relativeBase: 200,
			},
			args: args{
				line: 0,
			},
			wantNextLine: 2,
			wantNextBase: 221,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				name:         tt.fields.name,
				program:      tt.fields.program,
				relativeBase: tt.fields.relativeBase,
				Input:        tt.fields.Input,
				Output:       tt.fields.Output,
			}
			if gotNextLine := c.base(tt.args.line); gotNextLine != tt.wantNextLine {
				t.Errorf("Computer.base() return = %v, want %v", gotNextLine, tt.wantNextLine)
			}
			if tt.wantNextBase != c.relativeBase {
				t.Errorf("Computer.base() base = %v, want %v", c.relativeBase, tt.wantNextBase)
			}
		})
	}
}
