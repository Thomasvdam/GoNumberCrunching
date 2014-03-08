package crunch

import "math/big"

/* Parameters declared here for easy tweaking
 */
var (
  bigFactorialThreshold = big.NewInt(20000)
)

/* Stores an int and a string of the computations used
 */
type CrunchedNumber struct {
  N *big.Int
  How string
}

/* Functions as a wrapper for the Factorial function,
 * writes output to designated channel.
 */
func AddFactorial(x *CrunchedNumber, chBig, chSmall chan *CrunchedNumber) {
  temp := &CrunchedNumber{Factorial(x.N), (x.How + "f")}
  
  if temp.N.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
  }
}

/* Source: peterSO on stackoverflow
 * returns the factorial of a big.Int as a big.Int.
 */
func Factorial(N *big.Int) (result *big.Int) {
  result = new(big.Int)

  switch N.Cmp(&big.Int{}) {
  case -1, 0:
    result.SetInt64(1)
  default:
    result.Set(N)
    var one big.Int
    one.SetInt64(1)
    result.Mul(result, Factorial(N.Sub(N, &one)))
  }

  return
}

/* Functions as a wrapper for the BigSqrt function,
 * writes output to designated channel.
 */
func AddSqrt(x *CrunchedNumber, chBig, chSmall chan *CrunchedNumber) {
  temp := &CrunchedNumber{SqrtBig(x.N), (x.How + "s")}

  if temp.N.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
  }
}

/* Source: github.com/cznic/mathutil
 * Returns the floor of the square root of a big.Int as a big.Int
 */
func SqrtBig(N *big.Int) (result *big.Int) {
  switch N.Sign() {
  case -1:
    panic(-1)
  case 0:
    return big.NewInt(0)
  }

  var px, nx big.Int
  result = big.NewInt(0)
  result.SetBit(result, N.BitLen()/2+1, 1)
  for {
    nx.Rsh(nx.Add(result, nx.Div(N, result)), 1)
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
func AddFactorialSqrt(x *CrunchedNumber, chBig, chSmall chan *CrunchedNumber) {
  temp := &CrunchedNumber{FactorialSqrt(x.N), (x.How + "fs")}
  
  if temp.N.Cmp(bigFactorialThreshold) <= 0 {
    chSmall <- temp
  } else {
    chBig <- temp
  }
}

/* Rather than calculating the entire factorial, calculate the sqrt immediately
 * of the result of the factorial
 */
func FactorialSqrt(N *big.Int) (result *big.Int) {
  result = new(big.Int)

  switch N.Cmp(&big.Int{}) {
  case -1, 0:
    result.SetInt64(1)
  default:
    result.Set(N)
    var one big.Int
    one.SetInt64(1)
    result.Mul(result, FactorialSqrt(SqrtBig(N.Sub(N, &one))))
  }

  return
}