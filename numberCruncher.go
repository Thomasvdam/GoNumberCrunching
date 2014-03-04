package main

import (
  "fmt"
  "github.com/fatih/set"
  "math/big"
  "strings"
  "strconv"
)

/* Parameters declared here for easy tweaking
 *
 */
var (
  bigTwo = big.NewInt(2)
  bigFactorialThreshold = big.NewInt(20000)
)

/* Stores an int and a string of the computations used
 */
type crunchedNumber struct {
  n *big.Int
  how string
}

// Where the magic happens.
func main() {
  // Create the initial set and a set to store all found numbers
  initialNumbers := set.New(4)
  allNumbers := set.New()

  // Use a map to store the numbers we're interested in
  crunchedNumbers := make(map[int]*crunchedNumber)

  // Create two channels, one 'small' numbers and one for big ones
  channelBig := make(chan *crunchedNumber, 100)
  channelSmall := make(chan *crunchedNumber, 100)

  // Add the first number to the small channel
  for !initialNumbers.IsEmpty() {
    x := initialNumbers.Pop()
    temp := x.(int)
    i := int64(temp)
    firstNumber := &crunchedNumber{big.NewInt(i), ""}
    channelSmall <- firstNumber
  }

  // Loop while less then 100 values have been found
  found := 0
  for found < 100 {

    select {
    case nextNumber := <- channelSmall:
      // Convert it to an int and check whether it is in the 0-100 range
      // If so, add it to the crunchedNumbers map
      temp := nextNumber.n.Int64()
      if 0 < temp && temp <= 100 {
        value, ok := crunchedNumbers[int(temp)]
        if !ok {
          found++
          // fmt.Println(nextNumber.n, " as element : " ,found)
          crunchedNumbers[int(temp)] = nextNumber
        } else {
          oldNumber := strings.NewReader(value.how)
          newNumber := strings.NewReader(nextNumber.how)
          if oldNumber.Len() > newNumber.Len() {
            crunchedNumbers[int(temp)] = nextNumber
          }
        }
      }

      // If it has already been found, skip it, else add it to numbers found
      if allNumbers.Has(nextNumber.n.String()) {
        continue
      }
      allNumbers.Add(nextNumber.n.String())

      go AddFactorial(nextNumber, channelBig, channelSmall)
      AddSqrt(nextNumber, channelBig, channelSmall)

    case nextNumber := <- channelBig:
      // If it has already been found, skip it, else add it to numbers found
      if allNumbers.Has(nextNumber.n.String()) {
        continue
      }
      allNumbers.Add(nextNumber.n.String())

      AddSqrt(nextNumber, channelBig, channelSmall)
    }

  }

  // End of program, check which numbers have been found and print those.
  foundSlice := make([]int, 100)
  for x := 0; x < 100; x++ {
    value, ok := crunchedNumbers[x + 1]
    if ok {
      foundSlice[x] = x + 1
      fmt.Print(x + 1, " : ")
      PrintRoute(value)
    } else {
      foundSlice[x] = 0
    }
  }

  // Print all found numbers at the end
  //fmt.Println(foundSlice[:])
}

/* Functions as a wrapper for the Factorial function,
 * writes output to designated channel.
 */
func AddFactorial(x *crunchedNumber, chBig, chSmall chan *crunchedNumber) {
  temp := &crunchedNumber{Factorial(x.n), (x.how + "f")}
  
  if temp.n.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
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

/* Functions as a wrapper for the BigSqrt function,
 * writes output to designated channel.
 */
func AddSqrt(x *crunchedNumber, chBig, chSmall chan *crunchedNumber) {
  temp := &crunchedNumber{SqrtBig(x.n), (x.how + "s")}

  if temp.n.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
  }
}

/* Source: github.com/cznic/mathutil
 * Returns the floor of the square root of a big.Int as a big.Int
 */
func SqrtBig(n *big.Int) (result *big.Int) {
  switch n.Sign() {
  case -1:
    panic(-1)
  case 0:
    return big.NewInt(0)
  }

  var px, nx big.Int
  result = big.NewInt(0)
  result.SetBit(result, n.BitLen()/2+1, 1)
  for {
    nx.Rsh(nx.Add(result, nx.Div(n, result)), 1)
    if nx.Cmp(result) == 0 || nx.Cmp(&px) == 0 {
      break
    }
    px.Set(result)
    result.Set(&nx)
  }

  return
}

/* Functions as a wrapper for the FactorialSqrt function,
 * writes output to designated channel.
 */
func AddFactorialSqrt(x *crunchedNumber, chBig, chSmall chan *crunchedNumber) {
  temp := &crunchedNumber{FactorialSqrt(x.n), (x.how + "fs")}
  
  if temp.n.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
  }
}

/* Rather than calculating the entire factorial, calculate the sqrt immediately
 * of the result of the factorial
 */
func FactorialSqrt(n *big.Int) (result *big.Int) {
  result = new(big.Int)

  switch n.Cmp(&big.Int{}) {
  case -1, 0:
    result.SetInt64(1)
  default:
    result.Set(n)
    var one big.Int
    one.SetInt64(1)
    result.Mul(result, FactorialSqrt(SqrtBig(n.Sub(n, &one))))
  }

  return
}

/* Parses the instruction sequence of the crunchedNumber and outputs a more
 * human-friendly format.
 */
func PrintRoute(r *crunchedNumber) {
  reader := strings.NewReader(r.how)

  sqrt := false
  sqrtCount := 0
  path := "4 "

  for i := 0; i < reader.Len(); i++ {
    switch r.how[i] {
    case 'f':
      if sqrt {
        sqrt = false
        if sqrtCount > 1 {
          path = path + strconv.Itoa(sqrtCount) + "√ "
        } else {
          path = path + "√ "
        }
        sqrtCount = 0
      }
      path = path + "!"
    case 's':
      if sqrt {
        sqrtCount++
      } else {
        path = path + " "
        sqrtCount = 1
        sqrt = true
      }
    }
  }
  if sqrt {
    path = path + strconv.Itoa(sqrtCount) + "√"
  }
  fmt.Println(path)
}
