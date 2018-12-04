package model

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/lougazen/greedy-snake/lib"
)

// const
const (
	Width  = 20
	Height = 20
)

// DrawUI -
func DrawUI(p Location, ch byte) {
	x, y := p.Get()
	lib.GotoXY(x*2+4, y+2)
	fmt.Fprintf(os.Stderr, "%c", ch)
}

// Cls -
func Cls() {
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
