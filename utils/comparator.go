package utils

type Comparator interface {
	// Compare
	// a > b  -> return 1
	// a == b -> return 0
	// a < b  -> return -1
	Compare(a, b interface{}) ComparatorResult
}

type ComparatorResult struct {
	result int8
}

func (c *ComparatorResult) IsEqual() bool {
	return c.result == 0
}

// IsLess - return true if left operand less then right
func (c *ComparatorResult) IsLess() bool {
	return c.result == -1
}

type ComparatorInt struct{}

func (c *ComparatorInt) Compare(a, b interface{}) ComparatorResult {
	aInt, bInt := a.(int), b.(int)

	if aInt > bInt {
		return ComparatorResult{1}
	} else if aInt < bInt {
		return ComparatorResult{-1}
	} else {
		return ComparatorResult{0}
	}
}
