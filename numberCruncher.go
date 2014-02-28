package main

import (
  "fmt"
  "github.com/fatih/set"
  "math/big"
)

/* Source: peterSO on stackoverflow
 * returns the factorial of a big.Int as a big.Int.
 */
func Factorial(n *big.Int) (result *big.Int) {
  result = new(big.Int)

  switch n.Cmp(&big.Int{}) {
  case -1, 0:
    result.SetInt64(1)
  default:
    result.Set(n)
    var one big.Int
    one.SetInt64(1)
    result.Mul(result, Factorial(n.Sub(n, &one)))
  }

  return
}

/* Source: github.com/cznic/mathutil
 * Returns the floor of the square root of a big.Int as a big.Int
 */
func SqrtBig(n *big.Int) (x *big.Int) {
  switch n.Sign() {
  case -1:
    panic(-1)
  case 0:
    return big.NewInt(0)
  }

  var px, nx big.Int
  x = big.NewInt(0)
  x.SetBit(x, n.BitLen()/2+1, 1)
  for {
    nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
    if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
      break
    }
    px.Set(x)
    x.Set(&nx)
  }
fmt.Println(x)
  return
}

func main() {
  allNumbers := set.New(4)

  channel := make(chan *big.Int, 100)
  channel <- big.NewInt(4)

  for !allNumbers.Has(5) {
    nextNumber := <- channel
    go AddFactorial(nextNumber, channel)
    go AddSqrt(nextNumber, channel)
  }

  fmt.Println(allNumbers)

}

func AddFactorial(x *big.Int, ch chan * big.Int) {
  ch <- Factorial(x)
}

func AddSqrt(x *big.Int, ch chan *big.Int) {
  ch <- SqrtBig(x)
}