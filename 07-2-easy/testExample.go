package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// In your Go test file
func TestExample(t *testing.T) {
	// Set up a "before" hook that will run before the test.
	t.Run("before", func(t *testing.T) {
		// Do some setup work here.
	})

	// Set up an "after" hook that will run after the test.
	t.Run("after", func(t *testing.T) {
		// Do some cleanup work here.
	})

	// Write the actual test code here.
	t.Run("test", func(t *testing.T) {
		// Use assertions to check the expected results.
		assert.Equal(t, 1+1, 2, "1+1 must be equal to 2")
	})
}
