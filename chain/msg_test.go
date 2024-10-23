package chain

import (
	"github.com/stretchr/testify/assert"
	"github.com/wdvxdr1123/ZeroBot/message"
	"testing"
)

func TestMessageChain_Line(t *testing.T) {
	var c MessageChain
	one := message.Text("test")
	two := message.Text("test1")
	c.Line(one, two)

	assert.Equal(t, c[0], one)
	assert.Equal(t, c[1], two)
	assert.Equal(t, c[2], lb)
}

func TestMessageChain_Split(t *testing.T) {
	var c MessageChain
	one := message.Text("test")
	two := message.Text("test1")
	three := message.Text("test2")
	c.Split(one, two, three)

	assert.Equal(t, c[0], one)
	assert.Equal(t, c[1], lb)
	assert.Equal(t, c[2], two)
	assert.Equal(t, c[3], lb)
	assert.Equal(t, c[4], three)
	assert.Equal(t, len(c), 5)
}
