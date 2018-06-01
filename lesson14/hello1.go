// caseについて
package main

import (
  "fmt"
)

func main() {
  signal := "blue"
  switch signal {
  case "red":
    fmt.Println("stop")
  case "yellow":
    fmt.Println("caution")
  case "green", "blue": //　複数条件の場合はこのように記述
    fmt.Println("start")
  default: // ここがRubyと違う点
    fmt.Println("wrong signal")
  }
}
