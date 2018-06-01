//スライスの便利な使い方

package main

import (
  "fmt"
)

func main() {
  //いきなりスライスを作る方法
  //s := make([]int, 3) //かっこの中は型と長さ [0 0 0]となる
  s := []int{1,2,3} //要素の個数をいれないで、いきなり値を入れることもできる
  // append
  s = append(s, 4, 5, 6)
  // copy
  t := make([]int, len(s)) //引数の中は肩してとCapの数を指定している
  n := copy(t, s) //copyは要素数を返す
  fmt.Println(s)
  fmt.Println(t)
  fmt.Println(n)
}
