package main

import (
  "fmt"
  "flag"
  "math/big"
  "strings"
  "strconv"
  "github.com/fatih/set"
  "github.com/Arcania0311/numberCruncher/crunch"
)

var (
  printArray bool
  printProgress bool
  goal int
)

// Where the magic happens.
func main() {
  // Parse command line flags
  flag.BoolVar(&printArray, "array", false, "Output an array of path lengths rather than all paths.")
  flag.BoolVar(&printProgress, "progress", false, "Display notification when a new int has been found.")
  flag.IntVar(&goal, "goal", 100, "How many integers should be found before aborting.")
  flag.Parse()

  // Create the initial set and a set to store all found numbers
  initialNumbers := set.New(4)
  allNumbers := set.New()

  // Use a map to store the numbers we're interested in
  crunchedNumbers := make(map[int]*crunch.CrunchedNumber)

  // Create two channels, one 'small' numbers and one for big ones
  channelBig := make(chan *crunch.CrunchedNumber, 100)
  channelSmall := make(chan *crunch.CrunchedNumber, 100)

  // Add the first number to the small channel
  for !initialNumbers.IsEmpty() {
    x := initialNumbers.Pop()
    temp := x.(int)
    i := int64(temp)
    firstNumber := &crunch.CrunchedNumber{big.NewInt(i), ""}
    channelSmall <- firstNumber
  }

  // Loop while less then 100 values have been found
  found := 0
  for found < goal {

    select {
    case nextNumber := <- channelSmall:
      // Convert it to an int and check whether it is in the 0-100 range
      // If so, add it to the crunchedNumbers map
      temp := nextNumber.N.Int64()
      if 0 < temp && temp <= 100 {
        value, ok := crunchedNumbers[int(temp)]
        if !ok {
          found++
          if printProgress {
            fmt.Println(nextNumber.N, " as element : " ,found)
          }
          crunchedNumbers[int(temp)] = nextNumber
        } else {
          oldNumber := strings.NewReader(value.How)
          newNumber := strings.NewReader(nextNumber.How)
          if oldNumber.Len() > newNumber.Len() {
            crunchedNumbers[int(temp)] = nextNumber
          }
        }
      }

      // If it has already been found, skip it, else add it to numbers found
      if allNumbers.Has(nextNumber.N.String()) {
        continue
      }
      allNumbers.Add(nextNumber.N.String())

      go crunch.AddFactorial(nextNumber, channelBig, channelSmall)
      crunch.AddSqrt(nextNumber, channelBig, channelSmall)

    case nextNumber := <- channelBig:
      // If it has already been found, skip it, else add it to numbers found
      if allNumbers.Has(nextNumber.N.String()) {
        continue
      }
      allNumbers.Add(nextNumber.N.String())

      crunch.AddSqrt(nextNumber, channelBig, channelSmall)
    }

  }

  // End of program
  foundSlice := make([]int, 100)
  for x := 0; x < 100; x++ {
    value, ok := crunchedNumbers[x + 1]
    if ok {
      path, pathLength := PrintRoute(value)
      if !printArray {
        fmt.Println(x + 1, path)
      }
      foundSlice[x] = pathLength
    } else {
      foundSlice[x] = 0
    }
  }
  
  if printArray {
    // Print all found pathlengths at the end, wee bit hacky
    fmt.Print("[")
    for i, v := range foundSlice {
      fmt.Print(v)
      if i != 99 {
        fmt.Print(", ")      
      }
    }
    fmt.Println("]")
  }
}

/* Returns a string with a human friendly version of the path, plus the length
 * of the path.
 */
func PrintRoute(r *crunch.CrunchedNumber) (string, int) {
  reader := strings.NewReader(r.How)

  sqrt := false
  sqrtCount := 0
  fact := false
  factCount := 0
  path := ": 4"
  pathLength := 0

  for i := 0; i < reader.Len(); i++ {
    switch r.How[i] {
    case 'f':
      if sqrt {
        sqrt = false
        path = path + strconv.Itoa(sqrtCount) + "s" + " 1f"
        pathLength += sqrtCount + 1
        sqrtCount = 0
      }
      if fact {
        factCount++
      } else {
        path = path + " "
        factCount = 1
        fact = true
      }

    case 's':
      if fact {
        fact = false
        path = path + strconv.Itoa(factCount) + "!"
        pathLength += factCount
        factCount = 0
      }

      if sqrt {
        sqrtCount++
      } else {
        path = path + " "
        sqrtCount = 1
        sqrt = true
      }
    }
  }
  // Append the end of the path
  if sqrt {
    path = path + strconv.Itoa(sqrtCount) + "s"

    // Hacky 'solution' for the path of 2
    pathLength += sqrtCount
    if pathLength != 1 {
      pathLength++
      path = path + " 1f"
    }
  }
  if fact {
    path = path + strconv.Itoa(factCount) + "!"
    pathLength += factCount
  }

  // Prepend the path length
  path = "(" + strconv.Itoa(pathLength) + ") " + path

  return path, pathLength
}
