package main

func foo(x []int, a, b, i, j int) int {
	k := j
	ct := 0
	for k > i-1 {
		if x[k] <= b && !(x[k] <= a) {
			ct += 1
		}
		k -= 1
	}
	return ct
}

func do_1() {
	// x := []int{11, 10, 10, 5, 10, 15, 20, 10, 7, 11}
	// print(foo(x, 8, 18, 3, 6))
	// print(foo(x, 10, 20, 0, 9))
	// print(foo(x, 8, 18, 6, 3))
	// print(foo(x, 20, 10, 0, 9))
	// print(foo(x, 6, 7, 8, 8))
}
