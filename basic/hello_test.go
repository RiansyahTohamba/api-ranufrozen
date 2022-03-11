package basic_test

import (
	"api-ranufrozen/basic"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestHelloMain(m *testing.M) {
// 	fmt.Println("Before Test")
// 	m.Run()
// 	fmt.Println("After Test")
// }

func TestHelloWorldTable(t *testing.T) {
	// tipe data struct array?
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "basic.HelloWorld(bio)",
			request:  "bio",
			expected: "Hello bio",
		},
		{
			name:     "basic.HelloWorld(brad)",
			request:  "brad",
			expected: "Hello brad",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := basic.HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
func TestHelloWorld(t *testing.T) {
	result := basic.HelloWorld("bio")
	assert.Equal(t, "Hello bio", result, "Result must be 'Hello bio'")
}

func TestHelloSubTest(t *testing.T) {
	t.Run("bio", func(t *testing.T) {
		result := basic.HelloWorld("bio")
		assert.Equal(t, "Hello bio", result, "result must be 'Hello bio'")
	})
	t.Run("brad", func(t *testing.T) {
		result := basic.HelloWorld("brad")
		assert.Equal(t, "Hello brad", result, "result must be 'Hello brad'")
	})
}
