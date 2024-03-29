package rotation

func StringRightRotation(val string, d int) string {
	if d < 0 || len(val) == 0 {
		return val
	}
	// break string into runes
	r := len(val) - d%len(val)
	left := val[r:]
	right := val[:r]
	val = left + right
	return val
}
func StringLeftRotation(val string, d int) string {
	if d < 0 || len(val) == 0 {
		return val
	}
	// break string into runes
	r := d
	left := val[r:]
	right := val[:r]
	val = left + right
	return val
}

func StringLeftRotationMatch(x, y string) bool {
	matched := false
	left := y[:1]
	right := y[1:]
	xy := right + left
	if xy == x {
		return true
	}

	return matched
}
