package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/lougazen/greedy-snake/lib"
)

type location struct {
	x, y int
}

var (
	area = [20][20]byte{}
	food bool
	lead byte
	head location
	tail location
	size int
)

func cls() {
	cmd := &exec.Cmd{}
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("cmd", "/c", "clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func place() location {
	k := rand.Int() % 400
	return location{k / 20, k % 20}
}

func draw(p location, c byte) {
	lib.GotoXY(p.x*2+4, p.y+2)
	fmt.Fprintf(os.Stderr, "%c", c)
}

func init() {
	cls()

	head, tail = location{4, 4}, location{4, 4}
	lead, size = 'R', 1
	area[4][4] = 'H'
	rand.Seed(int64(time.Now().Unix()))

	fmt.Fprintln(os.Stderr,
		`
  ┌─────────────────────────────────────────┐
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  │                                         │
  └─────────────────────────────────────────┘
`)

	go func() {
		for {
			switch byte(lib.Direct()) {
			case 72:
				lead = 'U'
			case 75:
				lead = 'L'
			case 77:
				lead = 'R'
			case 80:
				lead = 'D'
			case 32:
				lead = 'P'
			}
		}
	}()
}

func main() {
	for {
		time.Sleep(time.Millisecond * 400)

		if lead == 'P' {
			continue
		}

		if !food {
			give := place()
			if area[give.x][give.y] == 0 {
				area[give.x][give.y] = 'F'
				draw(give, '$')
				food = true
			}
		}

		area[head.x][head.y] = lead

		switch lead {
		case 'U':
			head.y--
		case 'L':
			head.x--
		case 'R':
			head.x++
		case 'D':
			head.y++
		}

		if head.x < 0 || head.x >= 20 || head.y < 0 || head.y >= 20 {
			lib.GotoXY(0, 23)
			break
		}

		eat := area[head.x][head.y]

		if eat == 'F' {
			food = false

			size++
		} else if eat == 0 {
			draw(tail, ' ')

			dir := area[tail.x][tail.y]

			area[tail.x][tail.y] = 0

			switch dir {
			case 'U':
				tail.y--
			case 'L':
				tail.x--
			case 'R':
				tail.x++
			case 'D':
				tail.y++
			}
		} else {
			lib.GotoXY(0, 23)
			break
		}
		draw(head, '*')
	}

	fmt.Fprintf(os.Stderr, "score: %v", size-1)
}
