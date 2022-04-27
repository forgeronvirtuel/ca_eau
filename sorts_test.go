package water

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		toSort []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal use case",
			args: args{toSort: []int{4, 5, 2, 1, 3, 9, 10, 7}},
			want: []int{1, 2, 3, 4, 5, 7, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.toSort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		toSort []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "normal use case",
			args: args{toSort: []int{4, 5, 2, 1, 3, 9, 10, 7}},
			want: []int{1, 2, 3, 4, 5, 7, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.toSort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
