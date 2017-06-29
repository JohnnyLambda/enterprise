package enterprise

import (
	"fmt"
	"testing"
)

// TestHello tests if hello can be printed.
func TestHello(t *testing.T) {

	fmt.Sprintf("hello")

}

// BenchmarkHello tests how fast hello can be printed.
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

var result string

// TestHelloComplete tests if hello can be printed.
func TestHelloComplete(t *testing.T) {

	// always record the result of printing to prevent
	// the compiler eliminating the function call.

	var s string
	s = fmt.Sprintf("hello")
	result = s
}

// BenchmarkHelloComplete tests how fast hello can be printed.
func BenchmarkHelloComplete(b *testing.B) {

	// always record the result of printing to prevent
	// the compiler eliminating the function call.

	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	result = s
}
