// caseについて
// if else をswitchで記述する。

package main

import (
  "fmt"
)

func main() {
  score := 20
  switch {
  case score > 80:
    fmt.Println("great")
  default:
    fmt.Println("soso")
  }
}
