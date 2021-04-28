package sort

import "testing"

type Test [2][]int

type TestFunc func([]int) []int

var TestSet = []Test{
	Test{[]int{2, 1}, []int{1, 2}},
	Test{[]int{5, 3, 4, 6}, []int{3, 4, 5, 6}},
	Test{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	Test{[]int{9, 8, 5, 6, 3, 2, 0}, []int{0, 2, 3, 5, 6, 8, 9}},
}

func Equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	for k, fn := range map[string]TestFunc{
		"InsectionSort": InsectionSort,
		"ShellSort":     ShellSort,
		"MergeSort":     MergeSort,
		"QuickSort":     QuickSort,
		"QuickSort2":    QuickSort2,
	} {
		for i := range TestSet {
			x := fn(TestSet[i][0])
			if !Equal(TestSet[i][1], x) {
				t.Errorf("%v failed, %v, %v", k, TestSet[i][1], x)
			}
		}
	}
}
