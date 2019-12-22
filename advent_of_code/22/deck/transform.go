package deck

const dim = 3

type Transform [dim][dim]int

func (t Transform) Compose(other Transform, mod int) Transform {
	sol := Transform{}
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for k := 0; k < dim; k++ {
				sol[i][j] += t[i][k] * t[k][j]
				sol[i][j] %= mod
				sol[i][j] += mod
				sol[i][j] %= mod
			}
		}
	}
	return sol
}

func (t Transform) Apply(c *Card) {
	c.Pos = 1*t[0][1] + c.Pos*t[1][1] + c.Len*t[2][1]
	c.Pos += c.Len
	c.Pos %= c.Len
	c.Pos += c.Len
	c.Pos %= c.Len
}

var Reverse Transform = [3][3]int{
	{1, -1, 0},
	{0, -1, 0},
	{0, 1, 1},
}

//c.Pos += -n
// c.Pos += c.Len
// c.Pos %= c.Len
// c.Pos  =c.Pos + c.Len +
func NewCutTransform(n int) Transform {
	return [3][3]int{
		{1, -n, 0},
		{0, 1, 0},
		{0, 1, 1},
	}
}

func NewDealWithIncrement(n int) Transform {
	// c.Pos = (c.Pos * n) % c.Len
	return [3][3]int{
		{1, 0, 0},
		{0, n, 0},
		{0, 0, 1},
	}
}
