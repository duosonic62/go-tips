package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		lhs  int
		rhs  int
		want int
	}{
		{name: "test1", lhs: 0, rhs: 1, want: 1},
		{name: "test2", lhs: 1, rhs: -1, want: 0},
		{name: "test3", lhs: 2, rhs: 1, want: 3},
	}

	for _, test := range tests {
		test := test // 並行テストのためにインスタンスをforの中でのみ生存するようにする

		// parallel
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Add(test.lhs, test.rhs)
			if got != test.want {
				t.Errorf("%s expected: %v, got: %v", test.name, test.want, got)
			}
		})
	}
}
