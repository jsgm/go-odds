package odds

type Parlay struct{
	Odds []Odd
	Stake float64
}

// Creates a new parlay bet.
func NewParlay(odds []Odd, stake float64) Parlay{
	p := Parlay{Odds: odds, Stake: stake}
	return p
}

// Total amount of money received for the parlay bet.
func (p *Parlay) Payout() float64{
	payout := 0.0
	for _, o := range p.Odds{
		if o.Decimal() > 0.0{
			payout = payout + (p.Stake * o.Decimal())
		}
	}
	return payout
}

// Will return a positive float64 that equals the benefits.
func (p *Parlay) Profit() float64{
	return p.Payout() - p.Stake 
}

// Return the number of bets.
func (p *Parlay) SellectionCount() int{
	return len(p.Odds)
}

// Combined probability of the parlay bet.
func (p *Parlay) Probability() float64{
	totalOdds := 0.0
	for _, o := range p.Odds{
		totalOdds = totalOdds + o.Decimal()
	}
	return 100 / totalOdds
}