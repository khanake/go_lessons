package main

import (
  "fmt"
)

func main() {
  // 配列を代入を一緒に書くことができる
  //b := [3]int{1,2,3} //こうする
  b := [...]int{1,2,3,4} //個数は自明なので、こう書くこともできる
  //fmt.Println(b)
  fmt.Println(len(b)) //数を知る

}
