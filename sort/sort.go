package sort

func InsectionSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	if len(x) < 3 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	if len(x) < 4 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		if x[1] > x[2] {
			x[2], x[1] = x[1], x[2]
		}
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	for i := 1; i < len(x); i++ {
		for j := i; j > 0; j-- {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
	return x
}

func ShellSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	if len(x) < 3 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	if len(x) < 4 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		if x[1] > x[2] {
			x[2], x[1] = x[1], x[2]
		}
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	for l := len(x) / 2; l > 0; l = l / 2 {
		for i := l; i < len(x); i++ {
			for j := i; j > l; j -= l {
				if x[j] < x[j-l] {
					x[j], x[j-l] = x[j-l], x[j]
				}
			}
		}
	}
	return x
}

func MergeSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	if len(x) < 3 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	if len(x) < 4 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		if x[1] > x[2] {
			x[2], x[1] = x[1], x[2]
		}
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	l := len(x) / 2
	y := MergeSort(x[:l])
	z := MergeSort(x[l:])
	r := make([]int, len(x))
	i := 0
	j := 0
	for k := 0; k < len(r); k++ {
		if i == len(y) {
			copy(r[k:], z[j:])
			break
		}
		if j == len(z) {
			copy(r[k:], y[i:])
			break
		}
		if y[i] < z[j] {
			r[k] = y[i]
			i++
		} else {
			r[k] = z[j]
			j++
		}
	}
	return r
}

func QuickSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	if len(x) < 3 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	if len(x) < 4 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		if x[1] > x[2] {
			x[2], x[1] = x[1], x[2]
		}
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	i := 0
	for i < len(x)-1 {
		if len(x)-i == 3 {
			QuickSort(x[i:])
			break
		}
		if x[i+1] < x[i] {
			x[i+1], x[i] = x[i], x[i+1]
			i++
			continue
		}
		j := len(x) - 1
		for j > i+1 {
			if x[j] < x[i] {
				x[j], x[i+1] = x[i+1], x[j]
				x[i+1], x[i] = x[i], x[i+1]
				i++
				break
			}
			j--
		}
		if j == i+1 {
			break
		}
	}
	QuickSort(x[:i])
	QuickSort(x[i+1:])
	return x
}

func QuickSort2(x []int) []int {
	if len(x) < 2 {
		return x
	}
	if len(x) < 3 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	if len(x) < 4 {
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		if x[1] > x[2] {
			x[2], x[1] = x[1], x[2]
		}
		if x[0] > x[1] {
			x[0], x[1] = x[1], x[0]
		}
		return x
	}
	i := 0
	j := len(x) - 1
	for i < j {
		for i < j {
			if x[i] > x[0] {
				break
			}
			i++
		}
		for i < j {
			if x[j] < x[0] {
				break
			}
			j--
		}
		if i < j {
			x[i], x[j] = x[j], x[i]
			continue
		}
		if i == j {
			if x[i] > x[0] {
				i--
			}
			x[i], x[0] = x[0], x[i]
			break
		}
	}
	QuickSort2(x[:i])
	QuickSort2(x[i+1:])
	return x
}
