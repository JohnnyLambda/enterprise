package enterprise

import "errors"
import "math"

// Add Two numbers.
func Add(a int, b int) (c int) {
	c = a + b
	return c
}

// PiOverFour returns pi over 4 in binary
func PiOverFour() (s string) {
	piOverFour := math.Pi / 4
	s = ""
	x := piOverFour
	for index := 0; index < 10; index++ {
		x = piOverFour * 2
		if x > 1 {
			s = s + "1"
			x = x - 1
		} else {
			s = s + "0"
		}
	}
	//s = strconv.FormatFloat(piOverFour, 'f', 2, 64)
	return
}

// Subtract two numbers.
func Subtract(a int, b int) (c int) { // hi subract 1
	c = a - b // hi subtract 2
	return c  // hi subtract 3
} // hi subtract 4

// add Two numbers private.
func add(a int, b int) (c int) {
	c = a + b
	return c
}

// Yolo returns the Yolo string.
func Yolo() string {
	return "Yolo"
}

// Yolo42 returns 42.
func Yolo42() (awwYeah int) {
	return 42
}

// Divide two numbers, returning an error if the denominator is zero.
func Divide(numerator int, denominator int) (result int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return
	}

	result = numerator / denominator
	return
}
