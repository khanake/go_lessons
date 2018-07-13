// 省略形

package main

import (
  "fmt"
)

func main() {
  s := []int{2,3,8}
  for _, v := range s { // indexを消すとエラーになるなるブランクスコア(_)を使用する
    fmt.Println(v)
  }
}

