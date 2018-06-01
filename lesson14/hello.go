// caseについて
package main

import (
  "fmt"
)

func main() {
  signal := "www"
  switch signal {
  case "red":
    fmt.Println("stop")
  case "yellow":
    fmt.Println("caution")
  case "green":
    fmt.Println("start")
  default: // ここがRubyと違う点
    fmt.Println("wrong signal")
  }
}
