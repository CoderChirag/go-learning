package main

func sum(a ...any) any {
	s := 0
	for _, v := range a {
		if val, ok := v.(int); ok {
			s += val
		}
	}
	return s
}