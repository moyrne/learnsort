package bubble

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3,2,4,1,5",
			args: args{num: []int{3, 2, 4, 1, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]int{}, tt.args.num...)

			sortV1(tt.args.num)
			if !reflect.DeepEqual(tt.args.num, tt.want) {
				t.Errorf("sortV1() = %v, want %v", tt.args.num, tt.want)
			}

			tt.args.num = args
			sortV2(tt.args.num)
			if !reflect.DeepEqual(tt.args.num, tt.want) {
				t.Errorf("sortV2() = %v, want %v", tt.args.num, tt.want)
			}

			tt.args.num = args
			sortV3(tt.args.num)
			if !reflect.DeepEqual(tt.args.num, tt.want) {
				t.Errorf("sortV3() = %v, want %v", tt.args.num, tt.want)
			}
		})
	}
}
