package main

import "fmt"

func main() {
  fmt.Println("Enter a place name: ")
  var place string
  fmt.Scan(&place)

  fmt.Println("Enter a colour: ")
  var colour string
  fmt.Scan(&colour)

  fmt.Println("Enter a animal: ")
  var animal string
  fmt.Scan(&animal)

  fmt.Println("Enter a food: ")
  var food string
  fmt.Scan(&food)

  fmt.Println("You went to", place, "where you saw a", colour, animal, "and you went home and ate a huge", food)
}
