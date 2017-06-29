package enterprise

import (
	"fmt"
	"testing"
)

// BenchmarkHello tests how fast hello can be printed.
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
