package main
import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

var locs map[string]bool
var knot_count = 10

type MoveFunc func([][2]int, int)([][2]int)

func main() {
  moves := map[string]MoveFunc{"R":R, "L":L, "U":U, "D":D}
  knots := make([][2]int, knot_count)
  locs = map[string]bool{}

  record_tail(knots[9])

  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')

    if err != nil {
      break
    }

    line = strings.TrimSpace(line)

    move := moves[line[0:1]]
    dist, _ := strconv.Atoi(line[2:])

    knots = move(knots, dist)
    record_tail(knots[knot_count-1])

    fmt.Printf("move %s by %d\n", line[0:1], dist)
    dump_knots(knots)
  }

  fmt.Printf("total locations visited by tail: %d\n", len(locs))
}

func R(k [][2]int, d int) [][2]int {
  fmt.Printf("+++ R %d\n", d)
  for i := 0; i < d; i++ {
    k[0][0] = k[0][0] + 1
    for j := 0; j < knot_count - 1; j++ {
      //fmt.Printf("~~~ checking knots %d and %d\n", j, j+1)
      //dump_knots(k)
      if(!is_touching(k[j], k[j+1])) {
        k[j+1] = move_tail(k[j], k[j+1])
      }
      record_tail(k[knot_count-1])
    }
  }
  return k
}

func L(k [][2]int, d int) [][2]int {
  fmt.Printf("+++ L %d\n", d)
  for i := 0; i < d; i++ {
    k[0][0] = k[0][0] - 1
    for j := 0; j < knot_count - 1; j++ {
      //fmt.Printf("~~~ checking knots %d and %d\n", j, j+1)
      //dump_knots(k)
      if(!is_touching(k[j], k[j+1])) {
        k[j+1] = move_tail(k[j], k[j+1])
      }
      record_tail(k[knot_count-1])
    }
  }
  return k
}

func U(k [][2]int, d int) [][2]int {
  fmt.Printf("+++ U %d\n", d)
  for i := 0; i < d; i++ {
    k[0][1] = k[0][1] + 1
    for j := 0; j < knot_count - 1; j++ {
      //fmt.Printf("~~~ checking knots %d and %d\n", j, j+1)
      //dump_knots(k)
      if(!is_touching(k[j], k[j+1])) {
        k[j+1] = move_tail(k[j], k[j+1])
      }
      record_tail(k[knot_count-1])
    }
  }
  return k
}

func D(k [][2]int, d int) [][2]int {
  fmt.Printf("+++ D %d\n", d)
  for i := 0; i < d; i++ {
    k[0][1] = k[0][1] - 1
    for j := 0; j < knot_count - 1; j++ {
      //fmt.Printf("~~~ checking knots %d and %d\n", j, j+1)
      //dump_knots(k)
      if(!is_touching(k[j], k[j+1])) {
        k[j+1] = move_tail(k[j], k[j+1])
      }
      record_tail(k[knot_count-1])
    }
  }
  return k
}

func is_touching(h [2]int, t [2]int) bool {
  //fmt.Printf("--- h(%d,%d) t(%d,%d) ", h[0], h[1], t[0], t[1])

  // If we are 1 or less distant on both dimensions, we are touching
  if(math.Abs(float64(h[0] - t[0])) <= 1 && math.Abs(float64(h[1] - t[1])) <= 1) {
    //fmt.Println("touching")
    return true
  }
  //fmt.Printf("NOT touching ")
  return false
}

func move_tail(h [2]int, t [2]int) [2]int {
  if(h[0] == t[0]) {
    // same x value, move tail's y

    // if tail is under head, move tail up to just under head
    if(t[1] < h[1]) { t[1] = h[1] - 1 }

    // else tail is over head, move tail to just over head
    if(t[1] > h[1]) { t[1] = h[1] + 1 }
  } else if(h[1] == t[1]) {
    // same y value, move tail's x

    // if tail is left of head, move tail to just left of head
    if(t[0] < h[0]) { t[0] = h[0] - 1 }

    // else tail is right of head, move tail to just right of head
    if(t[0] > h[0]) { t[0] = h[0] + 1 }
  } else {
    // neither x nor y match, move diagonally toward head
    if(h[0] > t[0]) {
      t[0]++
    } else {
      t[0]--
    }

    if(h[1] > t[1]) {
      t[1]++
    } else{
      t[1]--
    }

    // if(math.Abs(float64(h[0] - t[0])) - math.Abs(float64(h[1] - t[1])) > 0) {
    //   // x delta is larger than y delta
    //   // set y's to equal, x to one offset
    //   t[1] = h[1]
    //   // if tail is left of head, move tail to just left of head
    //   if(t[0] < h[0]) { t[0] = h[0] - 1 }
    //
    //   // else tail is right of head, move tail to just right of head
    //   if(t[0] > h[0]) { t[0] = h[0] + 1 }
    // } else {
    //   // y delta is larger than x delta
    //   // set x's to equal, y to one offset
    //   t[0] = h[0]
    //   // if tail is under head, move tail up to just under head
    //   if(t[1] < h[1]) { t[1] = h[1] - 1 }
    //
    //   // else tail is over head, move tail to just over head
    //   if(t[1] > h[1]) { t[1] = h[1] + 1 }
    // }
  }

  //fmt.Printf("--> t(%d, %d)\n", t[0], t[1])
  return t
}

func record_tail(t [2]int) {
  idx := fmt.Sprintf("%d,%d", t[0], t[1])
  fmt.Printf("**** recording index: %s\n", idx)
  locs[idx] = true
}

func dump_knots(k [][2]int) {
  x_max := 0
  x_min := 0
  y_max := 0
  y_min := 0

  for i := 0; i < len(k); i++ {
    if(k[i][0] > x_max) { x_max = k[i][0] }
    if(k[i][0] < x_min) { x_min = k[i][0] }
    if(k[i][1] > y_max) { y_max = k[i][1] }
    if(k[i][1] < y_min) { y_min = k[i][1] }
  }

  p := make([][]string, y_max - y_min + 1)
  for i := range p {
    p[i] = make([]string, x_max - x_min + 1)
  }

  for i := 0; i < len(k); i++ {
    p[k[i][1] - y_min][k[i][0] - x_min] = fmt.Sprintf("%d", i)
  }

  fmt.Println("=============================")
  for i := len(p) - 1; i >= 0; i-- {
    for j := range p[i] {
      if(p[i][j] == "") { p[i][j] = "." }
      fmt.Printf("%s", p[i][j])
    }
    fmt.Println("")
  }
  fmt.Println("=============================")
}
