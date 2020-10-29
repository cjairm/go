package main

import "fmt"

func areTheyEqual(s1 []int, s2 []int) bool {
 // Write your code here
  var m1 map[int]int = make(map[int]int)
  var m2 map[int]int = make(map[int]int)

  for _, v := range s1 {
    m1[v]++
  }

  for _, v := range s2 {
    m2[v]++
  }

  if len(m1) != len(m2) {
    return false
  }

  for key := range m2 {
    if m1[key] != m2[key] {
      return false
    }
  }

	return true
}

func main() {
  // Call areTheyEqual() with test cases here
  got1 := areTheyEqual([]int{1, 2, 3, 4}, []int{1, 4, 3, 2})
  if got1 != true {
    fmt.Println("Got: ", got1, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 1")
  }

  got2 := areTheyEqual([]int{1, 2, 3, 4}, []int{1, 4, 3, 3})
  if got2 != false {
    fmt.Println("Got: ", got2, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 2")
  }

  got3 := areTheyEqual([]int{}, []int{})
  if got3 != true {
    fmt.Println("Got: ", got3, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 3")
  }

  got4 := areTheyEqual([]int{100000,100000,100000,100000,100000,100000}, []int{100000,100000,100000,100000,100000,100000})
  if got4 != true {
    fmt.Println("Got: ", got4, ". Expected: ", true)
  } else {
    fmt.Println("✔ Test 4")
  }

  got5 := areTheyEqual([]int{100000,100000,100000,100000,100000,2}, []int{1,100000,100000,100000,100000,100000})
  if got5 != false {
    fmt.Println("Got: ", got5, ". Expected: ", false)
  } else {
    fmt.Println("✔ Test 5")
  }
}
