package main

import "fmt"

func getCode(n, s int) string {
  i := s
  c := n

  // 65 - 90 = A - Z
  if n >= 65 && n <= 90 {
     for i > 0 {
       if c >= 90 {
         c = 65
       } else {
         c++
       }
       i--
    }
  // 97 - 122 = a - z
  } else if n >= 97 && n <= 122 {
    for i > 0 {
      if c >= 122 {
        c = 97
      } else {
        c++
      }
      i--
    }
  // 48 - 57 = 0 - 9
  } else if n >= 48 && n <= 57 {
    for i > 0 {
      if c >= 57 {
        c = 48
      } else {
        c++
      }
      i--
    }
  // -
  } else {
    return fmt.Sprintf("%c", rune(n))
  }

  return fmt.Sprintf("%c", rune(c))
}

func rotationalCipher(input string, rotationFactor int) string {
  // Write your code here
  var str string;
  for _, l := range input {
    str += getCode(int(l), rotationFactor)
  }
	return str
}

func main() {
  // Call rotationalCipher() with test cases here
  got1 := rotationalCipher("All-convoYs-9-be:Alert1.", 4)
  if got1 != "Epp-gsrzsCw-3-fi:Epivx5." {
    fmt.Println("Got: ", got1, ". Expected: ", "Epp-gsrzsCw-3-fi:Epivx5.")
  } else {
    fmt.Println("✔ Test 1")
  }

  got2 := rotationalCipher("abcdZXYzxy-999.@", 200)
  if got2 != "stuvRPQrpq-999.@" {
    fmt.Println("Got: ", got2, ". Expected: ", "stuvRPQrpq-999.@")
  } else {
    fmt.Println("✔ Test 2")
  }

  got3 := rotationalCipher("Zebra-493?", 3)
  if got3 != "Cheud-726?" {
    fmt.Println("Got: ", got3, ". Expected: ", "Cheud-726?")
  } else {
    fmt.Println("✔ Test 3")
  }

  got4 := rotationalCipher("abcdefghijklmNOPQRSTUVWXYZ0123456789", 39)
  if got4 != "nopqrstuvwxyzABCDEFGHIJKLM9012345678" {
    fmt.Println("Got: ", got4, ". Expected: ", "nopqrstuvwxyzABCDEFGHIJKLM9012345678")
  } else {
    fmt.Println("✔ Test 4")
  }

  got5 := rotationalCipher("", 0)
  if got5 != "" {
    fmt.Println("Got: ", got5, ". Expected: ", "")
  } else {
    fmt.Println("✔ Test 5")
  }
}
