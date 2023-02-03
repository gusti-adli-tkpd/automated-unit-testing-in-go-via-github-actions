package main

func add_even(a, b int) (res int) {
	if a % 2 == 0 && b % 2 == 0 {
		res = a + b
	} else if a % 2 == 0 {
		res = a
	} else if b % 2 == 0 {
		res = b
	}
	return
}