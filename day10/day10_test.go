package day10

import (
	"log"
	"reflect"
	"testing"
)

func Test_getVector(t *testing.T) {
	type args struct {
		c1 coord
		c2 coord
	}
	tests := []struct {
		name string
		args args
		want vector
	}{
		{
			name: "t1",
			args: args{
				c1: coord{x: 5, y: 3},
				c2: coord{x: 7, y: 12},
			},
			want: vector{xd: 2, yd: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVector(tt.args.c1, tt.args.c2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asteroidField_markBlockedAsteroids(t *testing.T) {
	type fields struct {
		xSize  int
		ySize  int
		field  [][]string
		sensor coord
	}
	type args struct {
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]string
	}{
		{
			name: "t1",
			fields: fields{
				xSize: 5,
				ySize: 5,
				field: [][]string{
					[]string{".", "#", ".", ".", "#"},
					[]string{".", ".", ".", ".", "."},
					[]string{"#", "#", "#", "#", "#"},
					[]string{".", ".", ".", ".", "#"},
					[]string{".", ".", ".", "#", "#"},
				},
				sensor: coord{2, 2},
			},
			args: args{},
			want: [][]string{
				[]string{".", "#", ".", ".", "#"},
				[]string{".", ".", ".", ".", "."},
				[]string{"-", "#", "#", "#", "-"},
				[]string{".", ".", ".", ".", "#"},
				[]string{".", ".", ".", "#", "#"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &asteroidField{
				xSize:  tt.fields.xSize,
				ySize:  tt.fields.ySize,
				field:  tt.fields.field,
				sensor: tt.fields.sensor,
			}
			a.markBlockedAsteroids()
			e := false
			for y := 0; y < a.ySize; y++ {
				for x := 0; x < a.xSize; x++ {
					if tt.want[y][x] != a.field[y][x] {
						e = true
						t.Errorf("Incorrect field returned: x: %v, y: %v", x, y)
					}
				}
			}
			if e {
				a.log("Got:")

				log.Println("Should be:")
				a := &asteroidField{
					xSize:  tt.fields.xSize,
					ySize:  tt.fields.ySize,
					field:  tt.want,
					sensor: tt.fields.sensor,
				}
				a.sensor = tt.fields.sensor
				a.markBlockedAsteroids()
				a.log("Want:")
			}
		})
	}
}

func Test_asteroidField_unmarkBlockedAsteroids(t *testing.T) {
	type fields struct {
		xSize  int
		ySize  int
		field  [][]string
		sensor coord
	}
	type args struct {
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]string
	}{
		{
			name: "t1",
			fields: fields{
				xSize: 5,
				ySize: 5,
				field: [][]string{
					[]string{".", "#", ".", ".", "#"},
					[]string{".", ".", ".", ".", "."},
					[]string{"-", "#", "#", "#", "-"},
					[]string{".", ".", ".", ".", "#"},
					[]string{".", ".", ".", "#", "#"},
				},
				sensor: coord{2, 2},
			},
			args: args{},
			want: [][]string{
				[]string{".", "#", ".", ".", "#"},
				[]string{".", ".", ".", ".", "."},
				[]string{"#", "#", "#", "#", "#"},
				[]string{".", ".", ".", ".", "#"},
				[]string{".", ".", ".", "#", "#"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &asteroidField{
				xSize:  tt.fields.xSize,
				ySize:  tt.fields.ySize,
				field:  tt.fields.field,
				sensor: tt.fields.sensor,
			}
			a.unmarkBlockedAsteroids()
			e := false
			for y := 0; y < a.ySize; y++ {
				for x := 0; x < a.xSize; x++ {
					if tt.want[y][x] != a.field[y][x] {
						e = true
						t.Errorf("Incorrect field returned: x: %v, y: %v", x, y)
					}
				}
			}
			if e {
				a.log("Got:")

				a := &asteroidField{
					xSize:  tt.fields.xSize,
					ySize:  tt.fields.ySize,
					field:  tt.want,
					sensor: tt.fields.sensor,
				}
				a.sensor = tt.fields.sensor
				a.markBlockedAsteroids()
				a.log("want:")
			}
		})
	}
}
func Test_findPosition(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name             string
		args             args
		wantBestLocation coord
		wantObserved     int
	}{
		{
			name: "t1",
			args: args{
				data: []string{".#..#", ".....", "#####", "....#", "...##"},
			},
			wantBestLocation: coord{x: 3, y: 4},
			wantObserved:     8,
		},
		{
			name: "t2",
			args: args{
				data: []string{"......#.#.", "#..#.#....", "..#######.", ".#.#.###..", ".#..#.....", "..#....#.#", "#..#....#.", ".##.#..###", "##...#..#.", ".#....####"},
			},
			wantBestLocation: coord{x: 5, y: 8},
			wantObserved:     33,
		},
		{
			name: "t3",
			args: args{
				data: []string{"#.#...#.#.", ".###....#.", ".#....#...", "##.#.#.#.#", "....#.#.#.", ".##..###.#", "..#...##..", "..##....##", "......#...", ".####.###."},
			},
			wantBestLocation: coord{x: 1, y: 2},
			wantObserved:     35,
		},
		{
			name: "t4",
			args: args{
				data: []string{".#..#..###", "####.###.#", "....###.#.", "..###.##.#", "##.##.#.#.", "....###..#", "..#.#..#.#", "#..#.#.###", ".##...##.#", ".....#.#.."},
			},
			wantBestLocation: coord{x: 6, y: 3},
			wantObserved:     41,
		},
		{
			name: "t5",
			args: args{
				data: []string{".#..##.###...#######", "##.############..##.", ".#.######.########.#", ".###.#######.####.#.", "#####.##.#.##.###.##", "..#####..#.#########", "####################", "#.####....###.#.#.##", "##.#################", "#####.##.###..####..", "..######..##.#######", "####.##.####...##..#", ".#####..#.######.###", "##...#.##########...", "#.##########.#######", ".####.#.###.###.#.##", "....##.##.###..#####", ".#.#.###########.###", "#.#.#.#####.####.###", "###.##.####.##.#..##"},
			},
			wantBestLocation: coord{x: 11, y: 13},
			wantObserved:     210,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBestLocation, gotBestSeen := findPosition(tt.args.data)
			if !reflect.DeepEqual(gotBestLocation, tt.wantBestLocation) {
				t.Errorf("findPosition() gotBestLocation = %v, want %v", gotBestLocation, tt.wantBestLocation)
				log.Println("Received:", gotBestSeen)
				a := newField(tt.args.data)
				a.setLocation(gotBestLocation, "\u2588")
				a.sensor = gotBestLocation
				a.markBlockedAsteroids()
				a.log("Got:")
				log.Println("Should be:", tt.wantObserved)
				a = newField(tt.args.data)
				a.setLocation(tt.wantBestLocation, "\u2588")
				a.sensor = tt.wantBestLocation
				a.markBlockedAsteroids()
				a.log("Want:")

			}
			if gotBestSeen != tt.wantObserved {
				t.Errorf("findPosition() gotBestSeen = %v, want %v", gotBestSeen, tt.wantObserved)
			}
		})
	}
}

func Test_getDirection(t *testing.T) {
	type args struct {
		v     vector
		ySize int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "vertical (Q1)",
			args: args{
				v:     vector{xd: 0, yd: -2},
				ySize: 99,
			},
			want: 300.0 + 100,
		},
		{
			name: "3/2 (Q1)",
			args: args{
				v:     vector{xd: 2, yd: -3},
				ySize: 99,
			},
			want: 300.0 + 3.0/2.0,
		},
		{
			name: "1 (Q1)",
			args: args{
				v:     vector{xd: 2, yd: -2},
				ySize: 99,
			},
			want: 300.0 + 2.0/2.0,
		},
		{
			name: "3/2 (Q1)",
			args: args{
				v:     vector{xd: 3, yd: -2},
				ySize: 99,
			},
			want: 300.0 + 2.0/3.0,
		},
		{
			name: "horizontal (Q1)",
			args: args{
				v:     vector{xd: 3, yd: 0},
				ySize: 99,
			},
			want: 300.0 + 0.0/3.0,
		},
		{
			name: "3/2 (Q2)",
			args: args{
				v:     vector{xd: 3, yd: 2},
				ySize: 99,
			},
			want: 200.0 + 3.0/2.0,
		},
		{
			name: "1 (Q2)",
			args: args{
				v:     vector{xd: 2, yd: 2},
				ySize: 99,
			},
			want: 200.0 + 2.0/2.0,
		},
		{
			name: "2/3 (Q2)",
			args: args{
				v:     vector{xd: 2, yd: 3},
				ySize: 99,
			},
			want: 200.0 + 2.0/3.0,
		},
		{
			name: "vertical (Q2)",
			args: args{
				v:     vector{xd: 0, yd: 2},
				ySize: 99,
			},
			want: 200.0 + 0.0,
		},

		{
			name: "3/2 (Q3)",
			args: args{
				v:     vector{xd: -2, yd: 3},
				ySize: 99,
			},
			want: 100.0 + 3.0/2.0,
		},
		{
			name: "1 (Q3)",
			args: args{
				v:     vector{xd: -2, yd: 2},
				ySize: 99,
			},
			want: 100.0 + 2.0/2.0,
		},
		{
			name: "2/3 (Q3)",
			args: args{
				v:     vector{xd: -3, yd: 2},
				ySize: 99,
			},
			want: 100.0 + 2.0/3.0,
		},
		{
			name: "horizontal (Q4)",
			args: args{
				v:     vector{xd: -2, yd: 0},
				ySize: 99,
			},
			want: 100.0,
		},
		{
			name: "3/2 (Q4)",
			args: args{
				v:     vector{xd: -3, yd: -2},
				ySize: 99,
			},
			want: 3.0 / 2.0,
		},
		{
			name: "1 (Q4)",
			args: args{
				v:     vector{xd: -2, yd: -2},
				ySize: 99,
			},
			want: 2.0 / 2.0,
		},
		{
			name: "2/3 (Q4)",
			args: args{
				v:     vector{xd: -2, yd: -3},
				ySize: 99,
			},
			want: 2.0 / 3.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirection(tt.args.v, tt.args.ySize); got != tt.want {
				t.Errorf("getDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_asteroidField_calculateDirections(t *testing.T) {
	type fields struct {
		xSize      int
		ySize      int
		aSize      int
		field      [][]string
		directions map[coord]float32
		sensor     coord
		targets    []coord
	}
	tests := []struct {
		name         string
		fields       fields
		wantDirCount int
		wantMap      map[coord]float32
	}{
		{
			name: "t1",
			fields: fields{
				xSize: 3,
				ySize: 3,
				aSize: 8,
				field: [][]string{
					[]string{"#", "#", "#"},
					[]string{"#", ".", "#"},
					[]string{"#", "#", "#"},
				},
				directions: make(map[coord]float32, 8),
				sensor:     coord{x: 1, y: 1},
			},
			wantDirCount: 8,
			wantMap: map[coord]float32{
				coord{x: 0, y: 0}: 1.0,
				coord{x: 1, y: 0}: 16.0,
				coord{x: 2, y: 0}: 13.0,
				coord{x: 0, y: 1}: 4.0,
				coord{x: 2, y: 1}: 12.0,
				coord{x: 0, y: 2}: 5.0,
				coord{x: 1, y: 2}: 8.0,
				coord{x: 2, y: 2}: 9.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &asteroidField{
				xSize:      tt.fields.xSize,
				ySize:      tt.fields.ySize,
				aSize:      tt.fields.aSize,
				field:      tt.fields.field,
				directions: tt.fields.directions,
				sensor:     tt.fields.sensor,
				targets:    tt.fields.targets,
			}
			a.calculateDirections()
			if len(tt.fields.directions) != 8 {
				t.Errorf("%v directions found instead of %v", len(tt.fields.directions), tt.wantDirCount)
			}
			for k, v := range tt.fields.directions {
				if tt.wantMap[k] != v {
					t.Errorf("Got %v on %+v instead of %v", v, k, tt.wantMap[k])

				}
			}
		})
	}
}
