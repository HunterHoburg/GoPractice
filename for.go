package main
import "fmt"

func main() {
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }
  // Go only has the for loop, but it can be used a few different ways

  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  // this will run continuously until you break the loop or return from the function inside
  for {
    fmt.Println("loop")
    break
  }
}
