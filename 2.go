package main

func g(str string) string {
	i := 0
	new_str := ""
	for i < len(str)-1 {
		new_str += string(str[i+1])
		i += 1
	}
	return new_str
}

func f(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	} else {
		return f(g(str)) + string(str[0])
	}
}

func h(n int, str string) string {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		str = f(str)
	}
	return str
}

func pow(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pow(x, y-1)
	}
}

func do_2() {
	print(h(1, "fruits"))
	print(h(2, "fruits"))
	print(h(5, "fruits"))
	// print(h(pow(2, 1000000000000000), "fruits"))
	print(h(pow(2, 10), "fruits"))
	print(h(pow(2, 37), "fruits"))
	// print(h(pow(2, 9831050005000007), "fruits"))
}
