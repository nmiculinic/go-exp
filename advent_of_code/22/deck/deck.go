package deck

type Card struct {
	Len   int64
	Pos   int64
	Value int64
}

type Deck []Card

func (c *Card) DealIntoNew() {
	Reverse.Apply(c)
}

func (c *Card) Cut(n int64) {
	NewCutTransform(n).Apply(c)
}

func (c *Card) DealWithIncrementN(n int64) {
	NewDealWithIncrement(n).Apply(c)
}

func New(n int64) Deck {
	sol := make([]Card, n)
	for i := range sol {
		sol[i] = Card{
			Len:   n,
			Pos:   int64(i),
			Value: int64(i),
		}
	}
	return sol
}

func (deck Deck) ToInts() []int {
	sol := make([]int, deck[0].Len)
	for _, c := range deck {
		sol[c.Pos] = int(c.Value)
	}
	return sol
}

func (data Deck) DealIntoNewStack() {
	for i := range data {
		data[i].DealIntoNew()
	}
}

func (data Deck) Cut(n int64) {
	for i := range data {
		data[i].Cut(n)
	}
}

func (data Deck) DealWithIncrementN(n int64) {
	for i := range data {
		data[i].DealWithIncrementN(n)
	}
}
