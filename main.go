package main

import (
	"fmt"
	"os"

	"github.com/lougazen/greedy-snake/model"
)

func main() {
	model.Cls()

	snake := model.Snake{}
	score := snake.Start()

	fmt.Fprintf(os.Stderr, "score: %v", score)
}
