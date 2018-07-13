// goroutine : 並行処理

package main

import (
  "fmt"
  "time"
)

func task1() {
  time.Sleep(time.Second * 2) // 2秒待つ
  fmt.Println("task1 finished")
}

func task2() {
  fmt.Println("task2 finished")
}

func main() {
  // 並列処理を行う例
  go task1()
  go task2()
  // これだと何も起こらない
  // なぜならgoroutineが終わる前にMain関数が終わってしまっているから
}
