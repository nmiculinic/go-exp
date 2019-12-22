package deck

type Card struct {
	Len   int
	Pos   int
	Value int
}

type Deck []Card

func (c *Card) DealIntoNew() {
	Reverse.Apply(c)
}

func (c *Card) Cut(n int) {
	c.Pos += -n
	c.Pos += c.Len
	c.Pos %= c.Len
}

func (c *Card) DealWithIncrementN(n int) {
	c.Pos = (c.Pos * n) % c.Len
	if gcd(n, c.Len) != 1 {
		panic("gcd != 1")
	}
}

func gcd(a, b int) int {
	if b > a {
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func New(n int) Deck {
	sol := make([]Card, n)
	for i := range sol {
		sol[i] = Card{
			Len:   n,
			Pos:   i,
			Value: i,
		}
	}
	return sol
}

func (deck Deck) ToInts() []int {
	sol := make([]int, deck[0].Len)
	for _, c := range deck {
		sol[c.Pos] = c.Value
	}
	return sol
}

func (data Deck) DealIntoNewStack() {
	for i := range data {
		data[i].DealIntoNew()
	}
}

func (data Deck) Cut(n int) {
	for i := range data {
		data[i].Cut(n)
	}
}

func (data Deck) DealWithIncrementN(n int) {
	for i := range data {
		data[i].DealWithIncrementN(n)
	}
}
