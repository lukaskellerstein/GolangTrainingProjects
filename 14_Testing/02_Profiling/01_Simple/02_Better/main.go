package mylib

var lookupTable = map[uint64]uint64{}

func Fibonacci(n uint64) uint64 {
	if _, resultExists := lookupTable[n]; resultExists {
		return lookupTable[n]
	}

	var result uint64

	if n == 0 {
		result = 0
	} else if n == 1 {
		result = 1
	} else {
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}

	lookupTable[n] = result
	return result
}
