package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
func main() {
    reader := bufio.NewReader(os.Stdin)

    a := []rune{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    count := 0

    for {
      char, _, err := reader.ReadRune()

      if err != nil {
          log.Fatal(err)
          break
      }

      a = append(a[1:14], 0)
      a[13] = char

      fmt.Printf("%d ", count)
      fmt.Println(a)
      count++

      dupe := false
      for i := 0; i < len(a); i++ {
        for j := i+1; j < len(a); j++ {
          if(a[i] == a[j]) { dupe = true }
        }
      }

      if(count >= 14 && !dupe) {
        fmt.Printf("first marker after %d", count)
        break
      }
    }
}
