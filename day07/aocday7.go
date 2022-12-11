package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type dir struct {
  size int
  subdirs map[string]dir
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var fs dir
    var cmd string

    for fs, cmd = process_commands(reader); cmd != ""; fs, cmd = process_commands(reader) {
      fmt.Printf("++++++++ back to top, next command: %s ++++++++\n", cmd)
    }
    fmt.Printf("size of /: %d\n", fs.size)
    free := 70000000 - fs.size
    fmt.Printf("free: %d\n", free)
    need_to_free := 30000000 - free
    fmt.Printf("need to free additional %d\n", need_to_free)
// part 1
//    size := walk_and_find(fs, "/")
//    fmt.Printf("final size: %d\n", size)
// part 2
    label, delta := find_closest_but_not_under(fs, "/", need_to_free)
    size := delta + need_to_free
    fmt.Printf("closest is: %s, size: %d\n", label, size)
}

func find_closest_but_not_under(fs dir, label string, ntf int) (string, int) {
  closest := label
  delta := fs.size - ntf
  fmt.Printf("+++ starting guess: %s, delta: %d\n", closest, delta)

  for d, m := range fs.subdirs {
    next_guess, next_delta := find_closest_but_not_under(m, d, ntf)
    if(next_delta >= 0 && next_delta < delta) {
      closest = next_guess
      delta = next_delta
      fmt.Printf("~~~ found closer guess: %s, delta: %d\n", closest, delta)
    }
  }

  return closest, delta
}

func walk_and_find(fs dir, label string) int {
  size := 0
  for d, m := range fs.subdirs {
    size = size + walk_and_find(m, d)
    fmt.Printf("+++ back from lower call, size now: %d\n", size)
  }

  if(fs.size <= 100000) {
    size = size + fs.size
    fmt.Printf("--- adding directory %s, %d; size now: %d\n", label, fs.size, size)
  }

  return size
}

func process_commands(reader *bufio.Reader) (dir, string) {
    var fs dir
    var next_cmd string
    var down_dir dir
    fs.size = 0
    fs.subdirs = make(map[string]dir)

    fmt.Println("-------- new process_commands --------")
    for {
      line, err := reader.ReadString('\n')

      if err != nil {
          fmt.Printf("exiting p_c due to: %s\n", err)
          return fs, ""
      }

      line = strings.TrimSpace(line)
      fmt.Println(line)

      if(line[0:4] == "$ cd") {
        dir := line[5:]
        if(dir == "/" || dir == "..") {
          return fs, line
        } else {
          down_dir, next_cmd = process_commands(reader)
          fmt.Printf("      ~~~~ back from process_commands, down_dir: %s, next_cmd: %s\n", down_dir, next_cmd)
          fs.subdirs[dir] = down_dir
          fs.size = fs.size + down_dir.size
          fmt.Printf("      ~~~~ local fs now: %s\n", fs)
          if(next_cmd == "$ cd /") {
            return fs, next_cmd
          }
        }
      } else if(line[0:1] != "$") {
        if(line[0:3] != "dir") {
          fmt.Printf("+++ file\n")
          fields := strings.Fields(line)
          s, _ := strconv.Atoi(fields[0])
          fs.size = fs.size + s
          fmt.Printf("      ~~~~ local fs.size: %d\n", fs.size)
        }
      }
    }
}
