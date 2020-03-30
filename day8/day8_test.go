package day8

import (
	"reflect"
	"testing"
)

func Test_newImage(t *testing.T) {
	type args struct {
		data   string
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *anImage
	}{
		{
			name: "t1",
			args: args{
				data:   "123456789012",
				width:  3,
				height: 2,
			},
			want: &anImage{
				width:      3,
				height:     2,
				layerCount: 2,
				pixels: [][][]int{
					{{1, 2, 3}, {4, 5, 6}},
					{{7, 8, 9}, {0, 1, 2}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newImage(tt.args.data, tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anImage_fewest0Digits(t *testing.T) {
	type fields struct {
		width      int
		height     int
		layerCount int
		pixels     [][][]int
	}
	tests := []struct {
		name            string
		fields          fields
		wantFewestZeros int
		wantBestLayer   int
	}{
		{
			name: "t1",
			fields: fields{
				pixels: [][][]int{
					{{1, 0, 0}, {4, 0, 0}},
					{{7, 0, 9}, {0, 0, 2}},
				},
				width:      3,
				height:     2,
				layerCount: 2,
			},
			wantFewestZeros: 3,
			wantBestLayer:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &anImage{
				width:      tt.fields.width,
				height:     tt.fields.height,
				layerCount: tt.fields.layerCount,
				pixels:     tt.fields.pixels,
			}
			gotFewestZeros, gotBestLayer := img.fewest0Digits()
			if gotFewestZeros != tt.wantFewestZeros {
				t.Errorf("anImage.fewest0Digits() gotFewestZeros = %v, want %v", gotFewestZeros, tt.wantFewestZeros)
			}
			if gotBestLayer != tt.wantBestLayer {
				t.Errorf("anImage.fewest0Digits() gotBestLayer = %v, want %v", gotBestLayer, tt.wantBestLayer)
			}
		})
	}
}

func Test_anImage_multiply1x2(t *testing.T) {
	type fields struct {
		width      int
		height     int
		layerCount int
		pixels     [][][]int
	}
	type args struct {
		l int
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
				pixels: [][][]int{
					{{1, 1, 2}, {4, 1, 2}},
					{{7, 1, 9}, {1, 0, 2}},
				},
				width:      3,
				height:     2,
				layerCount: 2,
			},
			args: args{l: 0},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &anImage{
				width:      tt.fields.width,
				height:     tt.fields.height,
				layerCount: tt.fields.layerCount,
				pixels:     tt.fields.pixels,
			}
			if got := img.multiply1x2(tt.args.l); got != tt.want {
				t.Errorf("anImage.multiply1x2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_anImage_pixelColor(t *testing.T) {
	type fields struct {
		width      int
		height     int
		layerCount int
		pixels     [][][]int
	}
	type args struct {
		r int
		c int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "top left",
			fields: fields{
				pixels: [][][]int{
					{{0, 2}, {2, 2}},
					{{1, 1}, {2, 2}},
					{{2, 2}, {1, 2}},
					{{0, 0}, {0, 0}},
				},
				width:      2,
				height:     2,
				layerCount: 4,
			},
			args: args{r: 0, c: 0},
			want: 0,
		},
		{
			name: "top right",
			fields: fields{
				pixels: [][][]int{
					{{0, 2}, {2, 2}},
					{{1, 1}, {2, 2}},
					{{2, 2}, {1, 2}},
					{{0, 0}, {0, 0}},
				},
				width:      2,
				height:     2,
				layerCount: 4,
			},
			args: args{r: 0, c: 1},
			want: 1,
		},
		{
			name: "bottom left",
			fields: fields{
				pixels: [][][]int{
					{{0, 2}, {2, 2}},
					{{1, 1}, {2, 2}},
					{{2, 2}, {1, 2}},
					{{0, 0}, {0, 0}},
				},
				width:      2,
				height:     2,
				layerCount: 4,
			},
			args: args{r: 1, c: 0},
			want: 1,
		},
		{
			name: "bottom right",
			fields: fields{
				pixels: [][][]int{
					{{0, 2}, {2, 2}},
					{{1, 1}, {2, 2}},
					{{2, 2}, {1, 2}},
					{{0, 0}, {0, 0}},
				},
				width:      2,
				height:     2,
				layerCount: 4,
			},
			args: args{r: 1, c: 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &anImage{
				width:      tt.fields.width,
				height:     tt.fields.height,
				layerCount: tt.fields.layerCount,
				pixels:     tt.fields.pixels,
			}
			if got := img.pixelColor(tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("anImage.pixelColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
