package selection

import (
	"reflect"
	"testing"
)

func Test_sortV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[3,2,4,1,5]",
			args: args{nums: []int{3, 2, 4, 1, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[]",
			args: args{nums: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// args := append([]int{}, tt.args.nums...)

			sortV1(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("sortV1() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
