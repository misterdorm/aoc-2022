package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    treemap := read_map()

    rows := len(treemap)
    cols := len(treemap[0])

    fmt.Printf("dimensions: %d x %d\n", rows, cols)

    high_score := 0
    for row, _ := range treemap {
      for col, _ := range treemap[row] {
        if(row == 0 || col == 0 || row == (rows - 1) || col == (cols - 1)) {
          // edge trees
        } else {
          score := scenic_score(treemap, row, col)
          if(score > high_score) {
            high_score = score
          }
          fmt.Printf("high score now: %d\n\n", high_score)
        }
      }
    }

    fmt.Printf("final high score: %d\n", high_score)
}

func scenic_score(m [][]int, r int, c int) int {
  // check in all four directions, here to the edge or a tree
  // equal or higher to us.  count how far we can see in each direction.
  // scenic score is the product of all four viewing distances
  h := m[r][c]

  fmt.Printf("---- checking tree at [%d, %d], height = %d\n", r, c, h)

  // look up
  d1 := 0
  for row := r - 1 ; row >= 0 ; row-- {
    d1 = d1 + 1
    if(m[row][c] >= h) {
      fmt.Printf("[UP] tree at [%d, %d] is %d, viewing distance: %d\n", row, c, m[row][c], d1)
      break
    }
  }
  fmt.Printf("[UP] distance: %d\n", d1)

  // look down
  d2 := 0
  for row := r + 1 ; row < len(m) ; row++ {
    d2 = d2 + 1
    if(m[row][c] >= h) {
      fmt.Printf("[DOWN] tree at [%d, %d] is %d, viewing distance: %d\n", row, c, m[row][c], d2)
      break
    }
  }
  fmt.Printf("[DOWN] distance: %d\n", d1)

  // look left
  d3 := 0
  for col := c - 1 ; col >= 0 ; col-- {
    d3 = d3 + 1
    if(m[r][col] >= h) {
      fmt.Printf("[LEFT] tree at [%d, %d] is %d, viewing distance: %d\n", r, col, m[r][col], d3)
      break
    }
  }
  fmt.Printf("[LEFT] distance: %d\n", d1)

  // look right
  d4 := 0
  for col := c + 1 ; col < len(m[0]) ; col++ {
    d4 = d4 + 1
    if(m[r][col] >= h) {
      fmt.Printf("[RIGHT] tree at [%d, %d] is %d, viewing distance: %d\n", r, col, m[r][col], d4)
      break
    }
  }
  fmt.Printf("[RIGHT] distance: %d\n", d1)

  fmt.Printf("scenic score: %d\n", (d1 * d2 * d3 * d4))
  return (d1 * d2 * d3 * d4)
}


func check_visible(m [][]int, r int, c int) int {
  // check in all four directions, here out to each edge
  // if at least one tree along the way is at least as tall
  // as us, we are invisible
  // if all trees are shorter than us, we are visible
  h := m[r][c]

  fmt.Printf("---- checking tree at [%d, %d], height = %d\n", r, c, h)

  // look up
  v := true
  for row := r - 1 ; row >= 0 ; row-- {
    if(m[row][c] >= h) {
      v = false
      fmt.Printf("tree at [%d, %d] is %d, we are invisible up\n", row, c, m[row][c])
      break
    }
  }
  if(v) {
    // visible up
    return 1
  }

  // look down
  v = true
  for row := r + 1 ; row < len(m) ; row++ {
    if(m[row][c] >= h) {
      v = false
      fmt.Printf("tree at [%d, %d] is %d, we are invisible down\n", row, c, m[row][c])
      break
    }
  }
  if(v) {
    // visible down
    return 1
  }

  // look left
  v = true
  for col := c - 1 ; col >= 0 ; col-- {
    if(m[r][col] >= h) {
      v = false
      fmt.Printf("tree at [%d, %d] is %d, we are invisible left\n", r, col, m[r][col])
      break
    }
  }
  if(v) {
    // visible left
    return 1
  }

  // look right
  v = true
  for col := c + 1 ; col < len(m[0]) ; col++ {
    if(m[r][col] >= h) {
      v = false
      fmt.Printf("tree at [%d, %d] is %d, we are invisible right\n", r, col, m[r][col])
      break
    }
  }
  if(v) {
    // visible right
    return 1
  }

  fmt.Println("invisible from all directions")
  return 0
}

func read_map() [][]int {
    reader := bufio.NewReader(os.Stdin)

    var treemap [][]int
    row := 0
    // col := 0

    for {
      line, err := reader.ReadString('\n')

      if err != nil {
          return treemap
      }

      line = strings.TrimSpace(line)
      fmt.Println(line)

      treemap = append(treemap, []int{})
      for _, x := range line {
        h, _ := strconv.Atoi(string(x))
        treemap[row] = append(treemap[row], h)
      }

      fmt.Println(treemap[row])
      row = row + 1
    }
}
