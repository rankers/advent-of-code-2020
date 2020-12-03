package main

import (
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  // Reads entire file in 1 go - no thought for memory
  dat, err := ioutil.ReadFile("day-1-input.txt")
  check(err)
  stringArray := strings.Split(string(dat), "\n")
  var result1, result2 int

  for _, i := range stringArray {
    factor1, err := strconv.Atoi(i)
    check(err)

    for _, j := range stringArray {
      factor2, err := strconv.Atoi(j)
      check(err)

      // No attempt to skip factors multiplying with themself. Could get bad result
      fmt.Print("Checking factor 1: ", factor1)
      fmt.Print(" with factor 2: ", factor2)
      fmt.Print(" Sum = ", factor1 + factor2, "\n")

      if (factor1 + factor2 == 2020) { // No attempt at break
        result1 = factor1
        result2 = factor2
      }
    }
  }
  fmt.Print("Hit!!", result1, result2, result1 + result2, "\n")
  fmt.Print("Result of multiplication: ", result1 * result2, "\n")
}
