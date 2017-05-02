package stringutil

// Reverse returns its argument string reversed
func Reverse(s string) string {
	return reverseRuneSlice([]rune(s), 0, len(s)-1)
}

// rotates a string  num chars to the left
func Rotate(s string, num int) string {
	return rotateRuneSlice([]rune(s), num)
}

// Reverses a rune slice  from the start to the end
func reverseRuneSlice(r []rune, start int, end int) string {
	for i := 0; i < (end-start+1)/2; i++ {
		r[i+start], r[end-i] = r[end-i], r[i+start]
	}

	return string(r)
}

// Rotates a rune slice by num steps in place i.e without using any extra space
// Consider the input slice as a sum of two slice. (A'B')' => BA
func rotateRuneSlice(r []rune, num int) string {
	if len(r) == 0 {
		return string(r)
	}
	reverseRuneSlice(r, 0, num-1)
	reverseRuneSlice(r, num, len(r)-1)
	reverseRuneSlice(r, 0, len(r)-1)
	return string(r)
}
