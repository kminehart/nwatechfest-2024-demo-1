package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, int64(20), Add(10, 10))
	assert.Equal(t, int64(32), Add(40, -8))
}
