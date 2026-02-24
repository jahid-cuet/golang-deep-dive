package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	
  /*
  int
  float32
  float64
  string
  bool
  */

//   var a int =10
a := 10
a=11 //No need to use colon 2nd time
fmt.Println(a)

//data types--> Numeric, String, Boolean

// if else and switch case
  age := 1
  if age >= 18 {
    fmt.Println("You are an adult.")
  } else if age >= 13 {
    fmt.Println("You are a minor.")
  } else {
    fmt.Println("You are a child.")
  }


  is_prime := true
  if is_prime {
    fmt.Println("The number is prime.")
  } else {
    fmt.Println("The number is not prime.")
  }

  day := "Monday"
  switch day {
  case "Monday":
    fmt.Println("It's the start of the week.")
  case "Friday":
    fmt.Println("It's the end of the week.")
  default:
    fmt.Println("It's a regular day.")
  }


  //functions


  x:= 5
  y:= 10
  sum := add(x, y)
  fmt.Println("The sum is:", sum)

}
  func add(a int, b int) int{
    return a+b
  }