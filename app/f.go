package main

type struct0 struct {
	s string
	i int
}

func f0(x int) struct0 {
	if x == 0 {
		panic("wololo")
	}
	res := struct0{
		s: "abc",
		i: 123,
	}
	return res
}
