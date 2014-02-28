package main

import (
  "fmt"
  "github.com/fatih/set"
  "math/big"
)

/* Functions as a wrapper for the Factorial function,
 * writes output to designated channel.
 */
func AddFactorial(x *big.Int, ch chan * big.Int) {
  ch <- Factorial(x)
}

/* Functions as a wrapper for the BigSqrt function,
 * writes output to designated channel.
 */
func AddSqrt(x *big.Int, ch chan *big.Int) {
  ch <- SqrtBig(x)
}

// Where the magic happens.
func main() {
  // Create the initial set and a set to store all found numbers
  initialNumbers := set.New(4)
  allNumbers := set.New()

  // Create a channel and add the set to it as big.Ints
  channel := make(chan *big.Int, 100)
  for !initialNumbers.IsEmpty() {
    x := initialNumbers.Pop()
    temp := x.(int)
    i := int64(temp)
    bigInt := big.NewInt(i)
    channel <- bigInt
  }

  // Loop while 5 is not in the set
  for {
    if allNumbers.Has(big.NewInt(5).String()) {
      break
    }

    // Get the next value from the channel
    nextNumber := <- channel

    // If it has already been found, skip it
    if allNumbers.Has(nextNumber.String()) || nextNumber.Cmp(big.NewInt(3)) <= 0 {
      continue
    }
    // Else add it to the found numbers
    allNumbers.Add(nextNumber.String())
    //fmt.Println(nextNumber)

    // Start individual threads for the factorial and root
    if nextNumber.Cmp(big.NewInt(100000)) <= 0 {
      go AddFactorial(nextNumber, channel)
    }
    go AddSqrt(nextNumber, channel)
  }

  fmt.Println(allNumbers)

}

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

  return
}
