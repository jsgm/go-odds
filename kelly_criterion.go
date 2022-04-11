package odds

import (
	"errors"
	"fmt"
)

func (o *Odd) KellyCriterion() (float64, error){
    if (*o == Odd{}) {
        return 0.0, errors.New("odd can't be empty")
    }

    if o.Decimal() <= 1.00{
        return 0.0, errors.New("odd is not valid")
    }
    
	b := o.Decimal() - 1.00

	fmt.Println("b: ", b)
    // Probabilities of winning.
    p := o.Probability() / 100.0
	fmt.Println("probability of winning: ", p)

    // Probabilities of lossing, 1.0 - probabilities of winning
    q := 1.0 - p
	fmt.Println("probability of lossing: ", q)
    
    return ((b*p) - q) / b, nil
}