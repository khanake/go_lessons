// rang
// 配列やスライスの要素分だけ処理を行う

package main

import (
  "fmt"
)

func main() {
  s := []int{2,3,8}
  for i, v := range s { // vにスライスの値 iに何番目の値かを入れてくれる
    fmt.Println(i,v)
  }
}

