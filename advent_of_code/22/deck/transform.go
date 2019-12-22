package deck

import (
	"fmt"
	"math/big"
)

const dim = 3

type Transform [dim][dim]int64

func (t Transform) GoString() string {
	a := fmt.Sprintf("%5d %5d %5d\n", t[0][0], t[0][1], t[0][2])
	a += fmt.Sprintf("%5d %5d %5d\n", t[1][0], t[1][1], t[1][2])
	a += fmt.Sprintf("%5d %5d %5d\n", t[2][0], t[2][1], t[2][2])
	return a
}

var _ fmt.GoStringer = Transform{}

func (t Transform) Compose(other Transform, mod int64) Transform {
	sol := Transform{}
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for k := 0; k < dim; k++ {
				sol[i][j] += modmul(t[i][k], other[k][j], mod)
				sol[i][j] += mod
				sol[i][j] %= mod
			}
		}
	}
	return sol
}

func (t Transform) Inverse(mod int64) Transform {
	sol := Identity

	DivRow := func(r int, div int64) {
		for i := 0; i < dim; i++ {
			muldiv := big.NewInt(0).ModInverse(big.NewInt(div), big.NewInt(mod)).Int64()
			t[r][i] = modmul(t[r][i], muldiv, mod)
		}

		for i := 0; i < dim; i++ {
			muldiv := big.NewInt(0).ModInverse(big.NewInt(div), big.NewInt(mod)).Int64()
			sol[r][i] = modmul(sol[r][i], muldiv, mod)
		}
	}

	// dr --> destination row
	// sr --> source row
	// mul --> how much to multiply source row before subbing destination row
	SubRow := func(dr, sr int, mul int64) {
		for i := 0; i < dim; i++ {
			t[dr][i] -= modmul(t[sr][i], mul, mod)
			t[dr][i] += mod
			t[dr][i] %= mod
		}

		for i := 0; i < dim; i++ {
			sol[dr][i] -= modmul(sol[sr][i], mul, mod)
			sol[dr][i] += mod
			sol[dr][i] %= mod
		}
	}

	for eliminatingRow := 0; eliminatingRow < dim; eliminatingRow++ {
		fmt.Printf("step %d [t, sol]:\n%#v\n%#v\n", eliminatingRow, t, sol)
		if t[eliminatingRow][eliminatingRow] == 0 {
			panic("just no. aaagh")
		}
		DivRow(eliminatingRow, t[eliminatingRow][eliminatingRow])
		fmt.Printf("step %d [t, sol]:\n%#v\n%#v\n", eliminatingRow, t, sol)

		// Divide the currently eliminating row
		for k := 0; k < dim; k++ {
			if k == eliminatingRow {
				continue
			}
			SubRow(k, eliminatingRow, t[k][eliminatingRow])
		}
		fmt.Printf("step %d [t, sol]:\n%#v\n%#v\n", eliminatingRow, t, sol)
		fmt.Println()
	}
	return sol
}

func (t Transform) Repeat(n, mod int64) Transform {
	x := t
	sol := Identity
	for n > 0 {
		if n&1 == 1 {
			sol = sol.Compose(x, mod)
		}
		x = x.Compose(x, mod)
		n /= 2
	}
	return sol
}

func (t Transform) Apply(c *Card) {
	c.Pos = 1*t[0][1] + modmul(c.Pos, t[1][1], c.Len) // + c.Len*t[2][1]
	c.Pos += c.Len
	c.Pos %= c.Len
	c.Pos += c.Len
	c.Pos %= c.Len
}

var Reverse Transform = [3][3]int64{
	{1, -1, 0},
	{0, -1, 0},
	{0, 1, 1},
}
var Identity Transform = [3][3]int64{
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 1},
}

func NewCutTransform(n int64) Transform {
	return [3][3]int64{
		{1, -n, 0},
		{0, 1, 0},
		{0, 1, 1},
	}
}

func NewDealWithIncrement(n int64) Transform {
	// c.Pos = (c.Pos * n) % c.Len
	return [3][3]int64{
		{1, 0, 0},
		{0, n, 0},
		{0, 0, 1},
	}
}

func modmul(a, b, mod int64) int64 {
	sol := big.NewInt(0)
	sol.Mul(big.NewInt(a), big.NewInt(b))
	sol.Mod(sol, big.NewInt(mod))
	return sol.Int64()
}
