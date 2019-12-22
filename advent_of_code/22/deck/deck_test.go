package deck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_Cut(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		d := New(10)
		d.Cut(3)
		assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, d.ToInts())
	})
	t.Run("negative", func(t *testing.T) {
		d := New(10)
		d.Cut(-4)
		assert.Equal(t, []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, d.ToInts())
	})
}
func TestCard_DealIntoNew(t *testing.T) {
	d := New(10)
	d.DealIntoNewStack()
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, d.ToInts())
}

func TestDeck_DealWithIncrementN(t *testing.T) {
	d := New(10)
	d.DealWithIncrementN(3)
	assert.Equal(t, []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, d.ToInts())
}
