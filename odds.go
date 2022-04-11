package odds

import (
	"errors"
)

const (
	Unknown = -1
	Decimal = 0 // check
	Fractional = 1
	Moneyline = 2 // check
	Implied = 3 // check
	Malay = 4
	Indonesian = 5
	HongKong = 6 // check

	// Rounding float types
	Round = 0
	Floor = 1
	Ceil = 2
)

type Odd struct{
	Type int
	DecimalValue float64

	FloatPrecision uint8
	Value interface{}
}

func NewOdd(oType int, oVal interface{}) (Odd, error){
	odd := Odd{Type: oType, Value: oVal, FloatPrecision: 2}

	switch oVal.(type) {
		case float64:
			if oType == Decimal{
				odd.DecimalValue = oVal.(float64)
			}else if oType == Implied{
				odd.DecimalValue = 100.0 / oVal.(float64)
			}else{
				return Odd{}, errors.New("float64 only accepts 'Decimal' and 'Implied' as odd format.")
			}
		default:
			return Odd{}, errors.New("Odd value data type is not correct.")
	}
	return odd, nil
}

func (odd *Odd) Precision(newPrecision uint8) Odd{
	odd.FloatPrecision = newPrecision
	return *odd
}

func (odd *Odd) Valid() bool{
	if (odd.Type == Decimal){
		return true
	}
	return false
}

// Returns the Moneyline value as int16 from Odd{}.
func (odd *Odd) Moneyline() int16{
	if odd.DecimalValue >= 2.0{
		return int16((odd.DecimalValue - 100) * 100)
	}else{
		return int16((-100) / (odd.DecimalValue - 1))
	}
}

// Returns the Implied Probability as float64 from Odd{}.
func (odd *Odd) Probability() float64{
	if odd.DecimalValue <= 0{
		return 0
	}
	return (1 / odd.DecimalValue) * 100
}

func (odd *Odd) Indonesian() float64{
	if odd.DecimalValue >= 2.0{
		return (odd.DecimalValue - 1)
	}else{
		return (-1) / (odd.DecimalValue - 1)
	}
}

func (odd *Odd) Malay() float64{
	if odd.DecimalValue >= 2.0{
		return (-1) / (odd.DecimalValue - 1)
	}else{
		return (odd.DecimalValue - 1)
	}
}

// Returns the decimal value as float64 from Odd{}.
func (odd *Odd) Decimal() float64{
	return odd.DecimalValue
}

func (odd *Odd) HongKong() float64{
	return (odd.DecimalValue - 1)
}