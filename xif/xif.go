package xif

func ifb(cond bool, ifTrue, ifFalse bool) bool {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func ifi(cond bool, ifTrue, ifFalse int) int {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func Ifs(cond bool, ifTrue, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}

func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func IfEquals(condition bool) bool {
	if condition {
		return true
	}
	return false
}

func IfEqualsMultipleInt(value int, values ...int) bool {
	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			if value == values[i] {
				return true
			}
		}
	}
	return false
}

func IfEqualsMultipleString(value string, values ...string) bool {
	if len(values) > 0 {
		for i := 0; i < len(values); i++ {
			if value == values[i] {
				return true
			}
		}
	}
	return false
}
