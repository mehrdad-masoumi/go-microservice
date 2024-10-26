package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandRange(t *testing.T) {
	minNumber := 1000
	maxNumber := 9999

	for i := 0; i < 100; i++ {
		result := RandRange(minNumber, maxNumber)

		assert.GreaterOrEqual(t, result, minNumber, "Result should be greater than or equal to min")
		assert.Less(t, result, maxNumber, "Result should be less than max")
	}
}
