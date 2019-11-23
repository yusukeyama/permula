package permula

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScanToSlicebyInt(t *testing.T) {
}

func TestCombination(t *testing.T) {
	//normal:3
	res := Combination(3, 2)
	assert.Equal(t, res, 3)
}

func TestPermutation(t *testing.T) {
	//right:3
	res := Permutation(3, 2)
	assert.Equal(t, res, 3)
	//right:10
	res = Permutation(5, 2)
	assert.Equal(t, res, 10)
}