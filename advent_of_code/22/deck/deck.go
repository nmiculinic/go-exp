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
	NewCutTransform(n).Apply(c)
}

func (c *Card) DealWithIncrementN(n int) {
	NewDealWithIncrement(n).Apply(c)
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
