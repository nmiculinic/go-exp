package deck

import "math/big"

const dim = 3

type Transform [dim][dim]int64

func (t Transform) Compose(other Transform, mod int64) Transform {
	sol := Transform{}
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for k := 0; k < dim; k++ {
				sol[i][j] += modmul(t[i][k], t[k][j], mod)
				sol[i][j] += mod
				sol[i][j] %= mod
			}
		}
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
