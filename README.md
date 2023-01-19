# `jexl-go`

A small library to evaluate jexl expressions in Go

## Getting Started

Get the pacakge like so

```bash
go get github.com/manfromth3m0oN/jexl-go
```

Then import the package and call `Eval()`

```go
import "github.com/manfromth3m0oN/jexl-go"

jexlExpr := "6 * 12 + 5 / 2.6"
eval, err := jexl.Eval(jexlExpr)
```
