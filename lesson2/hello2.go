// 変数
package main

import "fmt"

func main() {
  // var msg string //変数の定義 [var] [変数名] [型]
  //var msg = "Hello World" //いきなり代入もできる、この場合は型の省略が可能
  msg := "Hello World" //宣言(var)と値の紐づけを同時に行っている
  fmt.Println(msg)
  
  // 複数の変数を同時に紐付ける
  var a, b int // aのintは省略可能
  a, b = 10, 15
  fmt.Println(a,b)

  // := での記法も可能
  c, d := 20, 25
  fmt.Println(c,d)

  // 型が異なる変数を定義
  var (
    e int
    //e = 1
    f string
    //f = "hogehoge"
  )
  e = 1
  f = "hogehoge"
  fmt.Println(e,f)
}
