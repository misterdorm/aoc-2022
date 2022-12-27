package main
import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

type MoveFunc func([2]int, [2]int, int)([2]int, [2]int)

var locs map[string]bool

func main() {
  moves := map[string]MoveFunc{"R":R, "L":L, "U":U, "D":D}
  head := [2]int{0,0}
  tail := [2]int{0,0}
  locs = map[string]bool{}

  record_tail(tail)

  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')

    if err != nil {
      break
    }

    line = strings.TrimSpace(line)

    move := moves[line[0:1]]
    dist, _ := strconv.Atoi(line[2:])

    head, tail = move(head, tail, dist)

    fmt.Printf("move %s by %d, now at h(%d,%d) t(%d,%d)\n", line[0:1], dist, head[0], head[1], tail[0], tail[1])
  }

  fmt.Printf("total locations visited by tail: %d\n", len(locs))
}

func R(h [2]int, t [2]int, d int) ([2]int, [2]int) {
  fmt.Printf("+++ R %d\n", d)
  for i := 0; i < d; i++ {
    h[0] = h[0] + 1
    if(!is_touching(h, t)) { t = move_tail(h, t) }
  }
  return h, t
}

func L(h [2]int, t [2]int, d int) ([2]int, [2]int) {
  fmt.Printf("+++ L %d\n", d)
  for i := 0; i < d; i++ {
    h[0] = h[0] - 1
    if(!is_touching(h, t)) { t = move_tail(h, t) }
  }
  return h, t
}

func U(h [2]int, t [2]int, d int) ([2]int, [2]int) {
  fmt.Printf("+++ U %d\n", d)
  for i := 0; i < d; i++ {
    h[1] = h[1] + 1
    if(!is_touching(h, t)) { t = move_tail(h, t) }
  }
  return h, t
}

func D(h [2]int, t [2]int, d int) ([2]int, [2]int) {
  fmt.Printf("+++ D %d\n", d)
  for i := 0; i < d; i++ {
    h[1] = h[1] - 1
    if(!is_touching(h, t)) { t = move_tail(h, t) }
  }
  return h, t
}

func is_touching(h [2]int, t [2]int) bool {
  fmt.Printf("--- h(%d,%d) t(%d,%d) ", h[0], h[1], t[0], t[1])

  // If we are 1 or less distant on both dimensions, we are touching
  if(math.Abs(float64(h[0] - t[0])) <= 1 && math.Abs(float64(h[1] - t[1])) <= 1) {
    fmt.Println("touching")
    return true
  }
  fmt.Printf("NOT touching ")
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
    // neither x nor y match, diagonal
    // figure out which dimension has the bigger gap
    if(math.Abs(float64(h[0] - t[0])) - math.Abs(float64(h[1] - t[1])) > 0) {
      // x delta is larger than y delta
      // set y's to equal, x to one offset
      t[1] = h[1]
      // if tail is left of head, move tail to just left of head
      if(t[0] < h[0]) { t[0] = h[0] - 1 }

      // else tail is right of head, move tail to just right of head
      if(t[0] > h[0]) { t[0] = h[0] + 1 }
    } else {
      // y delta is larger than x delta
      // set x's to equal, y to one offset
      t[0] = h[0]
      // if tail is under head, move tail up to just under head
      if(t[1] < h[1]) { t[1] = h[1] - 1 }

      // else tail is over head, move tail to just over head
      if(t[1] > h[1]) { t[1] = h[1] + 1 }
    }
  }

  fmt.Printf("--> t(%d, %d)\n", t[0], t[1])
  record_tail(t)
  return t
}

func record_tail(t [2]int) {
  idx := fmt.Sprintf("%d,%d", t[0], t[1])
  fmt.Printf("**** recording index: %s\n", idx)
  locs[idx] = true
}
