package main

import (
	"fmt"
)

func main() {
  lesson2(10)
}

/**
 * The objective of this lesson is to use basic variables and print
 * them.
 */
func lesson1() {
  var a = "hello"
  var b = a
  var c = a + b
  fmt.Println("a = ", a)
  fmt.Println("b = ", b)
  fmt.Println("c = ", c)
}

/**
 * The objective of this lesson is to draw a right angle triangle
 * shape filled with "." characters.
 */
func lesson2(n int) {
  var a = ""
  for i := 0; i < n; i++ {
    a = a + "."
    fmt.Println(a)
  }
  
}