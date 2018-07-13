// 省略

package main

import (
  "fmt"
)

func main() {
  i := 0

  for i < 10 { //終了条件のみ記述
    fmt.Println(i)
    i++
  }

}
