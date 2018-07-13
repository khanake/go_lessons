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
  // 並列処理を行わない例
  task1()
  task2()
}
