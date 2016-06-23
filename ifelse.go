package main
import "fmt"

func main() {
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  // a statement can precede the conditional. If a variable is declared this way, it is accessible in all of the statement's branches

  if num:= 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
  // parentheses aren't required around the conditional, but the curly brackets are required
  // no ternaries in go
}
