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
	sol := ParseInput(t, mod).Repeat(n, mod)
	c := &deck.Card{
		Len:   mod,
		Pos:   pos,
		Value: pos,
	}
	sol.Apply(c)
	assert.Equal(t, int64(2322), c.Pos)

	t.Run("inv", func(t *testing.T) {
		g := sol.Inverse(mod)
		t.Logf("\n%#v", g)
		inv := sol.Compose(g, mod)
		t.Logf("inv\n%#v", inv)
		assert.Equal(t, deck.Identity, inv)
	})
}

func TestV2(t *testing.T) {
	var (
		mod int64 = 119315717514047
		pos int64 = 2020
		n   int64 = 101741582076661
	)
	sol := ParseInput(t, mod).Repeat(n, mod).Inverse(mod)
	c := &deck.Card{
		Len:   mod,
		Pos:   pos,
		Value: pos,
	}
	sol.Apply(c)
	assert.Equal(t, int64(49283089762689), c.Pos)
	t.Log(c)
}
