package gnum

func RoundTo(val int, to int) int {
	if val == 0 {
		return to
	} else if val%to == 0 {
		return val
	}
	return (val + (to - val%to))
}
