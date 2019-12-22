package v1

import (
	"bufio"
	"github.com/nmiculinic/go-exp/advent_of_code/22/deck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func ParseInput(t *testing.T, mod int64) deck.Transform {
	f, err := os.Open("input")
	require.NoError(t, err)
	defer f.Close()
	r := bufio.NewReader(f)
	sol := deck.Identity
	for {
		lineB, _, err := r.ReadLine()
		if err == io.EOF {
			return sol
		}
		require.NoError(t, err)
		line := strings.TrimSpace(string(lineB))
		switch {
		case line == "deal into new stack":
			sol = sol.Compose(deck.Reverse, mod)
		case line[:3] == "cut":
			n, err := strconv.ParseInt(line[4:], 10, 64)
			require.NoError(t, err)
			sol = sol.Compose(deck.NewCutTransform(n), mod)
		case line[:19] == "deal with increment":
			n, err := strconv.ParseInt(line[20:], 10, 64)
			require.NoError(t, err)
			sol = sol.Compose(deck.NewDealWithIncrement(n), mod)
		default:
			panic(line)
		}
	}

}

func TestV1(t *testing.T) {
	var (
		mod int64 = 10007
		pos int64 = 2019
		n   int64 = 1
	)
	_ = n
	sol := ParseInput(t, 10007)
	c := &deck.Card{
		Len:   mod,
		Pos:   pos,
		Value: pos,
	}
	sol.Apply(c)
	assert.Equal(t, int64(2322), c.Pos)
	t.Log(c)

}

func TestV2(t *testing.T) {
	var (
		mod int64 = 10007
		pos int64 = 2019
		n   int64 = 1
	)
	_ = n
	sol := ParseInput(t, 10007)
	c := &deck.Card{
		Len:   mod,
		Pos:   pos,
		Value: pos,
	}
	sol.Apply(c)
	assert.Equal(t, 2322, c.Pos)
	t.Log(c)
}
