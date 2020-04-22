package day17

import (
	"reflect"
	"testing"
)

func Test_robot_countOccurences(t *testing.T) {
	type fields struct {
		path []string
	}
	type args struct {
		path []string
		s    []string
	}
	tests := []struct {
		name        string
		path        []string
		args        args
		wantCount   int
		wantIndeces []int
	}{
		{
			name: "t1",
			args: args{
				path: []string{"L10", "L10", "R6", "L10", "L10", "R6", "R12", "L12", "L12", "R12", "L12", "L12", "L6", "L10", "R12", "R12", "R12", "L12", "L12", "L6", "L10", "R12", "R12", "R12", "L12", "L12", "L6", "L10", "R12", "R12", "L10", "L10", "R6"},
				s:    []string{"R12", "L12", "L12"},
			},
			wantCount:   4,
			wantIndeces: []int{6, 9, 16, 23},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotIndeces := countOccurences(tt.args.path, tt.args.s)
			if gotCount != tt.wantCount {
				t.Errorf("countOccurences(), Count = %v, want %v", gotCount, tt.wantCount)
			}
			if !reflect.DeepEqual(gotIndeces, tt.wantIndeces) {
				t.Errorf("countOccurences() Indeces= %v, want %v", gotIndeces, tt.wantIndeces)
			}

		})
	}
}
