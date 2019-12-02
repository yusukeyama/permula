package permula

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestScanToSlicebyInt(t *testing.T) {
}

func TestCombination(t *testing.T) {
	//right:3
	res := Combination(3, 2)
	assert.Equal(t, res, 3)
	//right:10
	res = Combination(5, 3)
	assert.Equal(t, res, 20)
	//fail:zero
	res = Combination(0, 0)
	assert.Equal(t, res, 0)
	//fail:min than right
	res = Combination(2, 3)
	assert.Equal(t, res, 0)
}

func TestPermutation(t *testing.T) {
	//right
	res := Permutation(3, 2)
	assert.Equal(t, res, 6)
	res = Permutation(5, 2)
	assert.Equal(t, res, 20)
	//fail:zero
	res = Permutation(0, 0)
	assert.Equal(t, res, 0)
	//fail:min than right
	res = Permutation(2, 3)
	assert.Equal(t, res, 0)
	//fail:minus
	res = Permutation(-1, -3)
	assert.Equal(t, res, 0)
	res = Permutation(3, -3)
	assert.Equal(t, res, 0)
}