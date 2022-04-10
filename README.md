# go-odds

<img align="right" src="https://github.com/jsgm/go-odds/raw/master/.github/go-odds.png" alt="go-odds" title="go-odds" />

[![Go Reference](https://pkg.go.dev/badge/badge/github.com/jsgm/go-odds.svg)](https://pkg.go.dev/github.com/jsgm/go-odds)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsgm/go-odds)](https://goreportcard.com/report/github.com/jsgm/go-odds)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

__go-odds__ is a simple and fast utility library for making betting odds conversions between formats and calculating probabilities in Go.

## Installation
Download the library using _go get_:
```go
go get https://github.com/jsgm/go-odds
```

Import in your current repository:
```go
import (
    odds "https://github.com/jsgm/go-odds"
)
```

## Data types
Each odd type will return a different data type, depending on your needs.
| Odd format | Function | README |
| ------ | ------ | ------ |
| Moneyline | .Moneyline() | int16 |
| Indonesian | .Indonesian() | float64 |
| Malay | .Malay() | float64 |
| Decimal | .Decimal() | float64 |
| HongKong | .HongKong() | float64 |
| Implied Probability | .Probability() | float64 |

## Convert odds between different formats
```go
o := NewOdd(Decimal, 1.64)

// Convert to implied probability
float := o.Probability()
```

## License
MIT