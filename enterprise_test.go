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
