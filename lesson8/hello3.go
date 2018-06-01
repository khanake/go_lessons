// 関数２
// 即時関数のようなこともできる
package main

import "fmt"

func main() {
  func(msg string) {
    fmt.Println(msg)
  }("Taro") //ここで引数に値を入れる
}
