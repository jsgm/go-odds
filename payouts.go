package odds

// Get the overround as float64 from an array of odds.
func GetOverround(odds []Odd) float64{
	sum := 0.0
	for _, o := range odds{
		sum = sum + o.Probability()
	}

	fmt.Println(sum)
	return sum - 100.0
}

// Get the payout as float64 from an array of odds.
func GetPayout(odds []Odd) float64{
	return 100.0 - GetOverround(odds)
}
