package day12

import (
	"testing"
)

func Test_aMoon_gravityFrom(t *testing.T) {
	type fields struct {
		position coord
		velocity vector
		gravity  vector
	}
	type args struct {
		o aMoon
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantVelocity vector
	}{
		{
			name: "less than",
			fields: fields{
				position: coord{0, 0, 0},
			},
			args: args{
				o: aMoon{
					position: coord{1, 2, 3},
				},
			},
			wantVelocity: vector{1, 1, 1},
		},
		{
			name: "greater than",
			fields: fields{
				position: coord{1, 2, 3},
			},
			args: args{
				o: aMoon{
					position: coord{0, 0, 0},
				},
			},
			wantVelocity: vector{-1, -1, -1},
		},
		{
			name: "equal",
			fields: fields{
				position: coord{1, 2, 3},
				velocity: vector{4, 5, 6},
			},
			args: args{
				o: aMoon{
					position: coord{1, 2, 3},
				},
			},
			wantVelocity: vector{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &aMoon{
				position: tt.fields.position,
				velocity: tt.fields.velocity,
			}
			m.gravityFrom(tt.args.o)
			if m.velocity != tt.wantVelocity {
				t.Errorf("aMoon_gravityFrom() got %v velocity, want %v", m.velocity, tt.wantVelocity)
			}
		})
	}
}

func Test_jupiterMoons_doGravity(t *testing.T) {
	type fields struct {
		ticks int
		moons []aMoon
	}
	tests := []struct {
		name      string
		fields    fields
		wantMoons []aMoon
	}{
		{
			name: "t1",
			fields: fields{
				moons: []aMoon{
					{ // moon1
						id:       0,
						position: coord{1, 2, 3},
						velocity: vector{4, 5, 6},
					},
					{ //moon 2
						id:       1,
						position: coord{3, 2, 1},
						velocity: vector{6, 5, 4},
					},
				},
			},
			wantMoons: []aMoon{
				{ // moon1
					id:       0,
					position: coord{1, 2, 3},
					velocity: vector{5, 5, 5},
				},
				{ //moon 2
					id:       1,
					position: coord{3, 2, 1},
					velocity: vector{5, 5, 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jupiterMoons{
				ticks: tt.fields.ticks,
				moons: tt.fields.moons,
			}
			j.doGravity()
			for i, m := range j.moons {
				if j.moons[i] != tt.wantMoons[i] {
					t.Errorf("Moon %v velocity is inccorrect, got %+v, want %+v", i, m, tt.wantMoons[i])
				}
			}
		})
	}
}

func Test_jupiterMoons_addVelocity(t *testing.T) {
	type fields struct {
		ticks int
		moons []aMoon
	}
	tests := []struct {
		name      string
		fields    fields
		wantMoons []aMoon
	}{
		{
			name: "t1",
			fields: fields{
				moons: []aMoon{
					{ // moon1
						id:       0,
						position: coord{1, 2, 3},
						velocity: vector{3, 5, 7},
					},
					{ //moon 2
						id:       1,
						position: coord{3, 2, 1},
						velocity: vector{7, 5, 3},
					},
				},
			},
			wantMoons: []aMoon{
				{ // moon1
					id:       0,
					position: coord{4, 7, 10},
					velocity: vector{3, 5, 7},
				},
				{ //moon 2
					id:       1,
					position: coord{10, 7, 4},
					velocity: vector{7, 5, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jupiterMoons{
				ticks: tt.fields.ticks,
				moons: tt.fields.moons,
			}
			j.addVelocity()
			for i, m := range j.moons {
				if j.moons[i] != tt.wantMoons[i] {
					t.Errorf("Moon %v position is inccorrect, got %+v, want %+v", i, m, tt.wantMoons[i])
				}
			}
		})
	}
}

func Test_aMoon_energy(t *testing.T) {
	type fields struct {
		id       int
		position coord
		velocity vector
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "pos",
			fields: fields{
				position: coord{1, 2, 3},
				velocity: vector{3, 2, 1},
			},
			want: 36,
		},
		{
			name: "mix",
			fields: fields{
				position: coord{1, -2, 3},
				velocity: vector{-3, 2, -1},
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &aMoon{
				id:       tt.fields.id,
				position: tt.fields.position,
				velocity: tt.fields.velocity,
			}
			if got := m.energy(); got != tt.want {
				t.Errorf("aMoon.energy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startCalculation(t *testing.T) {
	type args struct {
		data   []string
		cycles int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantMoons []aMoon
	}{
		{
			name: "c1",
			args: args{
				data:   []string{"<x=-1, y=0, z=2>", "<x=2, y=-10, z=-7>", "<x=4, y=-8, z=8>", "<x=3, y=5, z=-1>"},
				cycles: 1,
			},
			want: 229,
			wantMoons: []aMoon{
				{id: 0, position: coord{2, -1, 1}, velocity: vector{3, -1, -1}},
				{id: 1, position: coord{3, -7, -4}, velocity: vector{1, 3, 3}},
				{id: 2, position: coord{1, -7, 5}, velocity: vector{-3, 1, -3}},
				{id: 3, position: coord{2, 2, 0}, velocity: vector{-1, -3, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := newMoons(tt.args.data)
			if got := j.startCalculation(tt.args.cycles); got != tt.want {
				t.Errorf("startCalculation() = %v, want %v", got, tt.want)
			}
			for i, m := range j.moons {
				if j.moons[i].position != tt.wantMoons[i].position {
					t.Errorf("Moon %v position is inccorrect, got %+v, want %+v", m.id, m.position, tt.wantMoons[i].position)
				}
			}
			for i, m := range j.moons {
				if j.moons[i].velocity != tt.wantMoons[i].velocity {
					t.Errorf("Moon %v velocity is inccorrect, got %+v, want %+v", m.id, m.velocity, tt.wantMoons[i].velocity)
				}
			}
		})
	}
}
