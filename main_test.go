package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_div(t *testing.T) {
	var a float64
	var b float64
	var expected float64
	var res float64
	var err error

	a = 5.0
	b = 10.0
	expected = 0.5
	res, _ = div(a, b)
	assert.Equal(t, expected, res)

	a = 10.0
	b = 5.0
	expected = 2
	res, _ = div(a, b)
	assert.Equal(t, expected, res)

	a = 5.0
	b = 0.0
	expected = 0
	res, err = div(a, b)
	assert.NotNil(t, err)
	assert.Equal(t, expected, res)
}
