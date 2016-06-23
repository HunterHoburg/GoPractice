package main
import "fmt"

func main() {
  // var declares one or more variables, so var b, c int = 1, 2 printed out would print out 1 2
  var a string = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  // vars declared without an intialization value are 'zero-valued', so var e int would set the variable e to zero
  // := is shorthand for declaring and initializing a value, like this:
  h := "hunter"
  fmt.Println(h)

  // const declares constants
  // constants perform arithmetic with arbitrary precision, and don't have a type until one is 'cast' (I think by using something like math.Sin(2e20), which would make it type float64)
}
