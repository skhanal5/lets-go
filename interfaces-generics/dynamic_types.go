package main

// goal: write a generic add function
func add(a,b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}
	return nil
}

// Go's generic version of the above
func addGeneric[T int | float64|  string](a,b T) T {
	return a + b
}