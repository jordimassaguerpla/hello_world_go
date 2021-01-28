package main

import (
  "fmt"
  "math"
  "os"
)

type Vertex struct {
 X, Y int
}

type Abser interface {
  Abs() float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func main() {
  my_f := func() int {
    return 1
  }
  fmt.Println("hw", my_f)
  fmt.Println("hw", my_f())
  vertex := Vertex{1,2}
  fmt.Println("aaa", vertex.Abs())
  var vertex2 Abser
  vertex2 = vertex
  fmt.Println("bbb", vertex2.Abs())
  do_sth()
}

func do_sth() {
	var r error
	fmt.Println("r is")
	fmt.Println(r)
	fmt.Println("do sth")
	_, err := os.Open("filename.ext")
	r = err
	fmt.Println(r)
}
