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

    visible := 0
    for row, _ := range treemap {
      for col, _ := range treemap[row] {
        if(row == 0 || col == 0 || row == (rows - 1) || col == (cols - 1)) {
          // edge trees
          visible = visible + 1
        } else {
          visible = visible + check_visible(treemap, row, col)
          fmt.Printf("visible count now: %d\n\n", visible)
        }
      }
    }

    fmt.Println("total visible: %d\n", visible)
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
