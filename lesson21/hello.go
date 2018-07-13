// goroutine : 並行処理
// goroutine　だと returnで返り値を受け取ることができない
// channel　で解決する
// 一度channelに値をためて、そこから値を取り出すことができる

package main

import (
  "fmt"
  "time"
)

func task1(result chan string) {
  time.Sleep(time.Second * 2) // 2秒待つ
  result<- "task1 result"
}

func task2() {
  fmt.Println("task2 finished")
}

func main() {
  result := make(chan string)
  // 並列処理を行う例
  go task1(result)
  go task2()

  fmt.Println(<-result) // resultの結果がなければ、何かが入るまで待つ

  time.Sleep(time.Second * 3)
}
