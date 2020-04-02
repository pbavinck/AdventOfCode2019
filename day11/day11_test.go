package day11

import (
	"testing"

	"github.com/pbavinck/AofCode2019/machines"
)

func Test_robot_getLocation(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	type args struct {
		c coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "t1",
			args: args{
				c: coord{x: 0, y: 0},
			},
			fields: fields{
				surface: map[coord]string{
					coord{x: 0, y: 0}: ".",
				},
			},
			want: ".",
		},
		{
			name: "t2",
			args: args{
				c: coord{x: 0, y: 0},
			},
			fields: fields{
				surface: map[coord]string{},
			},
			want: blackHull,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			if got := r.getLocation(tt.args.c); got != tt.want {
				t.Errorf("robot.getLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robot_setLocation(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	type args struct {
		c     coord
		color string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "t1",
			args: args{
				c:     coord{x: 0, y: 0},
				color: "#",
			},
			fields: fields{
				surface: map[coord]string{
					coord{x: 0, y: 0}: ".",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			r.setLocation(tt.args.c, tt.args.color)
			if r.surface[tt.args.c] != tt.args.color {
				t.Errorf("robot_setLocation() got %v, want %v", r.surface[tt.args.c], tt.args.color)
			}
		})
	}
}

func Test_robot_paint(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	type args struct {
		colorCode string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantXMin  int
		wantXMax  int
		wantYMin  int
		wantYMax  int
		wantColor string
	}{
		{
			name: "t1",
			args: args{
				colorCode: "1",
			},
			fields: fields{
				position: coord{x: -5, y: 5},
			},
			wantXMin:  -5,
			wantXMax:  0,
			wantYMin:  0,
			wantYMax:  5,
			wantColor: whitePaint,
		},
		{
			name: "t1",
			args: args{
				colorCode: "0",
			},
			fields: fields{
				position: coord{x: 5, y: -5},
			},
			wantXMin:  0,
			wantXMax:  5,
			wantYMin:  -5,
			wantYMax:  0,
			wantColor: blackPaint,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			r.surface = make(map[coord]string)
			r.paint(tt.args.colorCode)
			if r.xMin != tt.wantXMin {
				t.Errorf("robot_paint() for xMin got %v, want %v", r.xMin, tt.wantXMin)
			}
			if r.xMax != tt.wantXMax {
				t.Errorf("robot_paint() for xMax got %v, want %v", r.xMax, tt.wantXMax)
			}
			if r.yMin != tt.wantYMin {
				t.Errorf("robot_paint() for yMin got %v, want %v", r.yMin, tt.wantYMin)
			}
			if r.yMax != tt.wantYMax {
				t.Errorf("robot_paint() for yMax got %v, want %v", r.yMax, tt.wantYMax)
			}
			if r.surface[tt.fields.position] != tt.wantColor {
				t.Errorf("robot_paint() for surface got %v, want %v", r.surface[tt.fields.position], tt.wantColor)
			}
		})
	}
}

func Test_robot_turn(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	type args struct {
		direction string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantDirection int
	}{
		{
			name: "t1",
			args: args{
				direction: "0", // left
			},
			fields: fields{
				direction: 0,
			},
			wantDirection: 3,
		},
		{
			name: "t2",
			args: args{
				direction: "1", // right
			},
			fields: fields{
				direction: 0,
			},
			wantDirection: 1,
		},
		{
			name: "t3",
			args: args{
				direction: "1", // right
			},
			fields: fields{
				direction: 3,
			},
			wantDirection: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			r.turn(tt.args.direction)
			if r.direction != tt.wantDirection {
				t.Errorf("robot_turn() for direction got %v, want %v", r.yMax, tt.wantDirection)
			}
		})
	}
}

func Test_robot_move(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	tests := []struct {
		name         string
		fields       fields
		wantPosition coord
	}{
		{
			name: "t1",
			fields: fields{
				position:  coord{x: -5, y: 5},
				direction: 3, //left
			},
			wantPosition: coord{x: -6, y: 5},
		},
		{
			name: "t2",
			fields: fields{
				position:  coord{x: -5, y: 5},
				direction: 2, //down
			},
			wantPosition: coord{x: -5, y: 6},
		},
		{
			name: "t3",
			fields: fields{
				position:  coord{x: -5, y: 5},
				direction: 1, //right
			},
			wantPosition: coord{x: -4, y: 5},
		},
		{
			name: "t4",
			fields: fields{
				position:  coord{x: -5, y: 5},
				direction: 0, //up
			},
			wantPosition: coord{x: -5, y: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			r.move()
			if r.position != tt.wantPosition {
				t.Errorf("robot_move() for position got %+v, want %+v", r.position, tt.wantPosition)
			}
		})
	}
}

func Test_robot_look(t *testing.T) {
	type fields struct {
		computer  *machines.Computer
		position  coord
		direction int
		surface   map[coord]string
		xMin      int
		xMax      int
		yMin      int
		yMax      int
	}
	tests := []struct {
		name               string
		fields             fields
		wantColorCode      string
		wantAlreadyPainted bool
	}{
		{
			name: "white paint",
			fields: fields{
				position: coord{x: -5, y: 5},
				surface: map[coord]string{
					coord{x: -5, y: 5}: whitePaint,
				},
			},
			wantColorCode:      whiteCode,
			wantAlreadyPainted: true,
		},
		{
			name: "black paint",
			fields: fields{
				position: coord{x: -5, y: 5},
				surface: map[coord]string{
					coord{x: -5, y: 5}: blackPaint,
				},
			},
			wantColorCode:      blackCode,
			wantAlreadyPainted: true,
		},
		{
			name: "black hull",
			fields: fields{
				position: coord{x: -5, y: 5},
				surface: map[coord]string{
					coord{x: -5, y: 5}: blackHull,
				},
			},
			wantColorCode:      blackCode,
			wantAlreadyPainted: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &robot{
				computer:  tt.fields.computer,
				position:  tt.fields.position,
				direction: tt.fields.direction,
				surface:   tt.fields.surface,
				xMin:      tt.fields.xMin,
				xMax:      tt.fields.xMax,
				yMin:      tt.fields.yMin,
				yMax:      tt.fields.yMax,
			}
			gotColorCode, gotAlreadyPainted := r.look()
			if gotColorCode != tt.wantColorCode {
				t.Errorf("robot.look() gotColorCode = %v, want %v", gotColorCode, tt.wantColorCode)
			}
			if gotAlreadyPainted != tt.wantAlreadyPainted {
				t.Errorf("robot.look() gotAlreadyPainted = %v, want %v", gotAlreadyPainted, tt.wantAlreadyPainted)
			}
		})
	}
}
