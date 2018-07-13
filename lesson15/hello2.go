// BreakとContinueについて

package main

import (
  "fmt"
)

func main() {

  for i := 0; i < 10; i++ { //初期化、終了条件、ループごとの処理
    if i == 3 { break } //ループを抜ける
    if i == 1 { continue } //処理をスキップ
    fmt.Println(i)
  }
}
