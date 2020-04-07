package day06

import (
	"testing"
)

func Test_countOrbits(t *testing.T) {
	type args struct {
		u aUniverse
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				u: aUniverse{
					"B": {name: "B", orbits: "COM"},
					"C": {name: "C", orbits: "B"},
					"D": {name: "D", orbits: "C"},
					"E": {name: "E", orbits: "D"},
					"F": {name: "F", orbits: "E"},
					"G": {name: "G", orbits: "B"},
					"H": {name: "H", orbits: "G"},
					"I": {name: "I", orbits: "D"},
					"J": {name: "J", orbits: "E"},
					"K": {name: "K", orbits: "J"},
					"L": {name: "L", orbits: "K"},
				},
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args.u).countOrbits(); got != tt.want {
				t.Errorf("countOrbits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countJumpsToSanta(t *testing.T) {
	type args struct {
		u aUniverse
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t2",
			args: args{
				u: aUniverse{
					"B":   {name: "B", orbits: "COM", jumpsToSanta: -1},
					"C":   {name: "C", orbits: "B", jumpsToSanta: -1},
					"D":   {name: "D", orbits: "C", jumpsToSanta: -1},
					"E":   {name: "E", orbits: "D", jumpsToSanta: -1},
					"F":   {name: "F", orbits: "E", jumpsToSanta: -1},
					"G":   {name: "G", orbits: "B", jumpsToSanta: -1},
					"H":   {name: "H", orbits: "G", jumpsToSanta: -1},
					"I":   {name: "I", orbits: "D", jumpsToSanta: -1},
					"J":   {name: "J", orbits: "E", jumpsToSanta: -1},
					"K":   {name: "K", orbits: "J", jumpsToSanta: -1},
					"L":   {name: "L", orbits: "K", jumpsToSanta: -1},
					"YOU": {name: "YOU", orbits: "K", jumpsToSanta: -1},
					"SAN": {name: "SAN", orbits: "I", jumpsToSanta: -1},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args.u).countJumpsToSanta(); got != tt.want {
				t.Errorf("countJumpsToSanta() = %v, want %v", got, tt.want)
			}
		})
	}
}
