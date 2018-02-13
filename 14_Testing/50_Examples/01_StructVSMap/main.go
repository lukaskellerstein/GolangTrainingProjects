package structvsmap

func AddToMap() {
	m := map[int]int{0: 0, 1: 1}
	_ = m[0] + m[1]
}

func AddToStruct() {
	m := struct{ a, b int }{0, 1}
	_ = m.a + m.b
}
