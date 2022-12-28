package main
import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type Monkey struct {
  items []int
  op string
  arg int
  test int
  t int
  f int
}

var debug bool

func main() {
  debug = true

//  rounds := 20
  monkies := read_notes()

  for i := range monkies {
    _ = i
  }

}

func read_notes() []Monkey {
  monkies := make([]Monkey, 0)
  var monkey_num int

  reader := bufio.NewReader(os.Stdin)
  for {
    line, err := reader.ReadString('\n')

    if err != nil {
      break
    }

    line = strings.TrimSpace(line)

    if(debug) { fmt.Printf("+++ Line: %s\n", line) }

    // start of monkey line?
    r, _ := regexp.Compile(`^Monkey (\d+):`)
    m := r.FindStringSubmatch(line)
    if(len(m) > 0) {
      monkies = append(monkies, Monkey{})
      monkey_num, _ = strconv.Atoi(m[1])
      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    // starting items?
    r, _ = regexp.Compile(`^Starting items: (.+)$`)
    m = r.FindStringSubmatch(line)
    if(len(m) > 0) {
      s, _ := regexp.Compile(`[, ]+`)
      m = s.Split(m[1], -1)

      for _, i_s := range m {
        i_i, _ := strconv.Atoi(i_s)
        monkies[monkey_num].items = append(monkies[monkey_num].items, i_i)
      }

      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    // operation?
    r, _ = regexp.Compile(`^Operation: new = old (\S+) (\S+)$`)
    m = r.FindStringSubmatch(line)
    if(len(m) > 0) {
      monkies[monkey_num].op = m[1]

      if(m[2] == "old") {
        monkies[monkey_num].op = monkies[monkey_num].op + m[1]
      } else {
        monkies[monkey_num].arg, _ = strconv.Atoi(m[2])
      }

      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    // test?
    r, _ = regexp.Compile(`^Test: divisible by (\d+)$`)
    m = r.FindStringSubmatch(line)
    if(len(m) > 0) {
      monkies[monkey_num].test, _ = strconv.Atoi(m[1])
      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    // true?
    r, _ = regexp.Compile(`^If true: throw to monkey (\d+)$`)
    m = r.FindStringSubmatch(line)
    if(len(m) > 0) {
      monkies[monkey_num].t, _ = strconv.Atoi(m[1])
      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    // false?
    r, _ = regexp.Compile(`^If false: throw to monkey (\d+)$`)
    m = r.FindStringSubmatch(line)
    if(len(m) > 0) {
      monkies[monkey_num].f, _ = strconv.Atoi(m[1])
      if(debug) { dump_monkey(monkey_num, monkies[monkey_num]) }
    }

    if(debug) { fmt.Println() }

  }

  return monkies
}

func dump_monkey(n int, m Monkey) {
  fmt.Printf("-----------------------------------------\n")
  fmt.Printf("Monkey %d:\n", n)
  fmt.Printf("  Items: %d\n", m.items)
  fmt.Printf("  Operation: %s, %d\n", m.op, m.arg)
  fmt.Printf("  Test: %d\n", m.test)
  fmt.Printf("    True: %d\n", m.t)
  fmt.Printf("    False: %d\n", m.f)
  fmt.Printf("-----------------------------------------\n")
}
