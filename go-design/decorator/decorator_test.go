package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	cs := NewColorSquare(sq,"red")
	got := cs.Draw()
	assert.Equal(t, "this is a square the color is red",got)
}
