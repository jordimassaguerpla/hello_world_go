package main

import (
  "fmt"
  "math"
)

const World = "世界"

func something(first, second int) (int, string) {
  return first + second, "hello"

}

func else_(first, second int) ( x int, y string) {
  x = first
  y = string(y)
  return

}
func main() {
  var t string
  var j int
  k, st := something(1,2)
  j, t = else_(1, 2)
  fmt.Println("something: ", k, st)
  fmt.Println("else: ", j, t)
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
  fmt.Printf("type %T", "aaa")
  fmt.Printf("type %T", 6.66)
  fmt.Printf("type %T", math.Pi)
  fmt.Printf("world ", World)
}
