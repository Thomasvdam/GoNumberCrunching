package main

import (
  "fmt"
  "github.com/fatih/set"
  "math/big"
)

/* Stores an int and a string of the computations used
 */
type crunchedNumber struct {
  n *big.Int
  how string
}

/* Functions as a wrapper for the Factorial function,
 * writes output to designated channel.
 */
func AddFactorial(x *crunchedNumber, ch chan *crunchedNumber) {
  temp := &crunchedNumber{Factorial(x.n), (x.how + "Fact")}
  ch <- temp
}

/* Functions as a wrapper for the BigSqrt function,
 * writes output to designated channel.
 */
func AddSqrt(x *crunchedNumber, ch chan *crunchedNumber) {
  temp := &crunchedNumber{SqrtBig(x.n), (x.how + "Sqrt")}
  ch <- temp
}

// Where the magic happens.
func main() {
  // Create the initial set and a set to store all found numbers
  initialNumbers := set.New(4)
  allNumbers := set.New()
  crunchedNumbers := make(map[int]*crunchedNumber)

  // Create a channel and add the set to it as big.Ints
  channel := make(chan *crunchedNumber, 100)
  for !initialNumbers.IsEmpty() {
    x := initialNumbers.Pop()
    temp := x.(int)
    i := int64(temp)
    firstNumber := &crunchedNumber{big.NewInt(i), ""}
    channel <- firstNumber
  }

  found := 4

  // Loop while less then 50 values have been found
  for found < 50 {
    // Get the next value from the channel
    nextNumber := <- channel

    // If it has already been found, skip it
    if allNumbers.Has(nextNumber.n.String()) || nextNumber.n.Cmp(big.NewInt(3)) < 0 {
      continue
    }
    // Else add it to the found numbers
    allNumbers.Add(nextNumber.n.String())

    // Convert it to an int and check whether it is in the 0-100 range
    // If so, add it to the crunchedNumbers map under its string
    temp := nextNumber.n.Int64()
    if 0 < temp && temp <= 100 {
      found++
      fmt.Println(nextNumber.n)
      crunchedNumbers[int(temp)] = nextNumber
    }   

    // Start individual threads for the factorial and root
    if nextNumber.n.Cmp(big.NewInt(10000000)) <= 0 {
      go AddFactorial(nextNumber, channel)
    }
    AddSqrt(nextNumber, channel)
  }

  for x := 1; x <= 100; x++ {
    value, ok := crunchedNumbers[x]
    if ok {
      fmt.Println("Check :", x)
      fmt.Println("By doing :", value.how)
    }
  }

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
