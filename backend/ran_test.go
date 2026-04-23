package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := add(2, 3)
	assert.Equal(t, 5, got)
}
