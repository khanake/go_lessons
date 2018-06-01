package main

import (
  "fmt"
)

func main() {
  score := 30
  if score > 80 { //条件の部分にかっこはいらない
    fmt.Println("great")
  } else if score > 60 {
    fmt.Println("good")
  } else {
    fmt.Println("soso")
  }
}
