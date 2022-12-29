package main
import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
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

  rounds := 10000
  monkies := read_notes()
  inspections := map[int]int{}

  debug = true

  shared_mod := 1
  for _, m := range monkies {
    shared_mod = shared_mod * m.test
  }

  for r := 0; r < rounds; r++ {
    for n, m := range monkies {
      for i := range m.items {
        inspections[n]++
        if(debug) { fmt.Printf("round: %d, monkey: %d, item: %d(%d)", r, n, i, m.items[i]) }
        m.items[i] = adjust_worry(m, i)
        m.items[i] = adjust_damage(m, i, shared_mod)
        if(debug) { fmt.Printf(" after worry/damage: %d\n", m.items[i]) }

        if(m.items[i] % m.test == 0) {
          monkies[m.t].items = append(monkies[m.t].items, m.items[i])
          if(debug) { fmt.Printf("---> [%d, T] throwing to %d\n", m.test, m.t) }
        } else {
          monkies[m.f].items = append(monkies[m.f].items, m.items[i])
          if(debug) { fmt.Printf("---> [%d, F] throwing to %d\n", m.test, m.f) }
        }

        if(debug) { fmt.Println() }
      }
      // clear items list, monkey threw them all
      monkies[n].items = []int{}
    }

    debug = false

    if(r == 0 || r == 19 || r % 1000 == 999) {
      fmt.Printf("Round %d\n", r + 1)
      for i, _ := range monkies {
        fmt.Printf("Monkey %d: %d\n", i, inspections[i])
      }
      for i, m := range monkies {
        fmt.Printf("Monkey %d: %d\n", i, m.items)
      }
      fmt.Println()
    }
  }

  ii := map[int]int{}
  for k, v := range inspections {
    ii[v] = k
  }

  var a []int
  for k := range ii {
    a = append(a, k)
  }

  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  monkey_business := 0
  for _, k := range a {
    fmt.Printf("Monkey %d inspected items %d times.\n", ii[k], k)
    if(monkey_business == 0) {
      monkey_business = k
    } else {
      monkey_business = monkey_business * k
      break
    }
  }

  fmt.Printf("monkey business: %d\n", monkey_business)

}

func adjust_worry(m Monkey, i int) int {
  switch m.op {
  case "+":
    return m.items[i] + m.arg
  case "*":
    return m.items[i] * m.arg
  case "++":
    return m.items[i] + m.items[i]
  case "**":
    return m.items[i] * m.items[i]
  }
  return m.items[i]
}

func adjust_damage(m Monkey, i int, s int) int {
  return m.items[i] % s
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
