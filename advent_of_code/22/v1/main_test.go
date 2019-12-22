package v1

import (
	"bufio"
	"github.com/nmiculinic/go-exp/advent_of_code/22/deck"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestV1(t *testing.T) {
	f, err := os.Open("input")
	require.NoError(t, err)
	defer f.Close()
	r := bufio.NewReader(f)
	c := deck.Card{
		Len:   10007,
		Pos:   2019,
		Value: 2019,
	}

	for {
		lineB, _, err := r.ReadLine()
		if err == io.EOF {
			t.Log(c)
			return
		}
		require.NoError(t, err)
		line := strings.TrimSpace(string(lineB))
		switch {
		case line == "deal into new stack":
			c.DealIntoNew()
			t.Log("deal into new stack", c)
		case line[:3] == "cut":
			n, err := strconv.ParseInt(line[4:], 10, 64)
			require.NoError(t, err)
			c.Cut(n)
			t.Log("cut ", n, c)
		case line[:19] == "deal with increment":
			n, err := strconv.ParseInt(line[20:], 10, 64)
			require.NoError(t, err)
			c.DealWithIncrementN(n)
			t.Log("deal with increment", n, c)
		default:
			panic(line)
		}
	}
}
