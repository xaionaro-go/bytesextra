package bytesextra

func min(args ...int) int {
	min := args[0]
	for _, cmp := range args[1:] {
		if cmp < min {
			min = cmp
		}
	}
	return min
}
