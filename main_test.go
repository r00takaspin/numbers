package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, 21, Gcd(861, 798))
	assert.Equal(t, 7, Gcd(1547, 560))

	assert.Equal(t, 42, Gcd(3528, 672, 1890))
	assert.Equal(t, 5, Gcd(195, 455, 1190))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 32718, Lcm(861, 798))
	assert.Equal(t, 123760, Lcm(1547, 560))

	assert.Equal(t, 211680, Lcm(3528, 672, 1890))
	assert.Equal(t, 46410, Lcm(195, 455, 1190))
	assert.Equal(t, 1116, Lcm(279, 372))
	assert.Equal(t, 3276, Lcm(252, 468))
}
