# go-odds

<img align="right" src="https://github.com/jsgm/go-odds/raw/master/.github/go-odds.png" alt="go-odds" title="go-odds" />

[![Go Reference](https://pkg.go.dev/badge/badge/github.com/jsgm/go-odds.svg)](https://pkg.go.dev/github.com/jsgm/go-odds)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsgm/go-odds)](https://goreportcard.com/report/github.com/jsgm/go-odds)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

__go-odds__ is a simple and fast utility library for making betting odds conversions between formats and calculating probabilities in Go.

Go gopher illustration created by [Egon Elbre](https://github.com/egonelbre), originally created by [Renee French](https://reneefrench.blogspot.com/).

- [Installation](#installation)
- [Converting odds between formats](#converting-odds-between-formats)
  * [Data types](#data-types)
  * [Odds conversions](#odds-conversions)
- [Overround and Payouts](#overround-and-payouts)
- [Arbitrage Calculation](#arbitrage-calculation)
- [Kelly Criterion](#kelly-criterion)
- [Simple Parlay Bets](#installation)

## Installation
Download the library using _go get_:
```
go get github.com/jsgm/go-odds
```

Import in your current repository:
```go
import (
    odds "github.com/jsgm/go-odds"
)
```

## Converting odds between formats
### Data types
Each odd type will return a different data type, depending on your needs.
| Odd format | Function | Type |
| ------ | ------ | ------ |
| Moneyline | .Moneyline() | int16 |
| Indonesian | .Indonesian() | float64 |
| Malay | .Malay() | float64 |
| Decimal | .Decimal() | float64 |
| HongKong | .HongKong() | float64 |
| Implied Probability | .Probability() | float64 |

### Odds conversions
```go
o := odds.NewOdd(odds.Decimal, 1.64)

// Convert to implied probability
float := o.Probability()
```

## Overround and Payouts
```go
home, _ := odds.NewOdd(odds.Implied, 65.00) // Home odds from Implied Probability
draw, _ := odds.NewOdd(odds.Decimal, 4.00)
away, _ := odds.NewOdd(odds.Decimal, 7.00)

fmt.Println("Overround ", odds.GetOverround([]odds.Odd{home, draw, away})) // 4.285714285714278
fmt.Println("Payout ", odds.GetPayout([]odds.Odd{home, draw, away})) // 95.71428571428572
```
## Arbitrage Calculation
todo

## Kelly Criterion
Kelly Criterion is simply strategy that helps calculating the proper stake that you should wager on a particular event. This doesn't just applies to betting but also for investing and calculating risks.

<img align="right" src="https://latex.codecogs.com/svg.latex?\Large&space;K%=\frac{bp-q}{b}" alt="Kelly Criterion Formula">

- b is the multiple of the stake.
- p is the probability of winning. 70% chances of winning should be 0.7.
- q is the probability of lossing. i.e. 100% - p = 30. The value would be 0.3.
- K% is the suggested percentage of the bankroll for the event.

## Simple Parlay Bets
```go	
o1, _ := odds.NewOdd(odds.Decimal, 1.25)
o2, _ := odds.NewOdd(odds.Implied, 50.7)
o3, _ := odds.NewOdd(odds.Decimal, 4.5)

multi := []odds.Odd{o1, o2, o3} 
stake := 10.0

p := odds.NewParlay(multi, stake)

p.Profit() // 67.22386587771203
p.Payout() // 77.22386587771203
p.SellectionCount() // 3
p.Probability() // 12.949364663814572 (%)
```

## Round Robin
todo
## Teaser
todo