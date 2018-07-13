// 繰り返しについて
// 繰り返しはforのみしかないwhileなどは存在しないので注意が必要

package main

import (
  "fmt"
)

func main() {

  for i := 0; i < 10; i++ { //初期化、終了条件、ループごとの処理
    fmt.Println(i)
  }
}
