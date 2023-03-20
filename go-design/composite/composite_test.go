package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDepartNum(t *testing.T) {

	personCount := NewDepartNum().Count()

	assert.Equal(t, 10,personCount)
}
