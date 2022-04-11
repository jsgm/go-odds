package odds

import (
	_ "fmt"
)
// Get the overround as float64 from an array of odds.
func GetOverround(odds []Odd) float64{
	sum := 0.0
	for _, o := range odds{
		sum = sum + o.Probability()
	}
	return sum - 100
}

// Get the payout as float64 from an array of odds.
func GetPayout(odds []Odd) float64{
	return 100.0 - GetOverround(odds)
}
