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

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{name: "easy", args: args{a: 1, b: 2}, wantC: 3},
		{name: "neg", args: args{a: -1, b: -2}, wantC: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := add(tt.args.a, tt.args.b); gotC != tt.wantC {
				t.Errorf("add() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
