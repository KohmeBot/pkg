package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	v1 := X(1).Y(6).Z(300)
	assert.Equal(t, "1.6.300", v1.String())
	v2 := X(1).Y(7).Z(3)
	assert.Equal(t, "1.7.3", v2.String())
	assert.True(t, v1 < v2)
	assert.Equal(t, uint16(1), v1.GetX())
	assert.Equal(t, uint16(6), v1.GetY())
	assert.Equal(t, uint32(300), v1.GetZ())
}
