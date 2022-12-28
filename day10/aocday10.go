package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func main() {
  reg_history := make([]int, 1)
  reader := bufio.NewReader(os.Stdin)

  cycle := 0
  reg_history[0] = 1

  for {
    line, err := reader.ReadString('\n')

    if err != nil {
      break
    }

    line = strings.TrimSpace(line)

    instr := line[0:4]
    arg := 0

    if(instr == "addx") {
      // addx, 2 cycles
      arg, _ = strconv.Atoi(line[5:])
      reg_history = append(reg_history, reg_history[cycle])
      cycle++
      reg_history = append(reg_history, reg_history[cycle] + arg)
    } else if(instr == "noop") {
      // noop
      reg_history = append(reg_history, reg_history[cycle])
    }

    cycle++
    fmt.Printf("cycle %d: instr: %s, arg: %d, x now %d\n", cycle, instr, arg, reg_history[cycle])

  }
  sum := sig_strength(reg_history)

  fmt.Printf("signal strength sum: %d\n", sum)
}

func sig_strength(h []int) int {
  start := 19
  offset := 40
  sum := 0

  for i := start; i < len(h); i = i + offset {
    fmt.Printf("%d: %d\n", i, h[i])
    sum = sum + (i + 1) * h[i]
  }

  return sum
}
