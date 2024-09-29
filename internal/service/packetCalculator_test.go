package service

import (
	"reflect"
	"testing"
)

var packSizes = []int{5000, 2000, 1000, 500, 250}

func TestCalculatePacks(t *testing.T) {
	type args struct {
		order     int
		packSizes []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Test 1",
			args: args{
				order:     9,
				packSizes: []int{3, 2, 1},
			},
			want: map[int]int{
				3: 3,
			},
		},
		{
			name: "Example 1",
			args: args{
				order:     1,
				packSizes: packSizes,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Example 2",
			args: args{
				order:     250,
				packSizes: packSizes,
			},
			want: map[int]int{
				250: 1,
			},
		},
		{
			name: "Example 3",
			args: args{
				order:     251,
				packSizes: packSizes,
			},
			want: map[int]int{
				500: 1,
			},
		},
		{
			name: "Example 4",
			args: args{
				order:     501,
				packSizes: packSizes,
			},
			want: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name: "Example 5",
			args: args{
				order:     12001,
				packSizes: packSizes,
			},
			want: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePacks(tt.args.order, tt.args.packSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
