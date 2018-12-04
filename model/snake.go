package model

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/lougazen/greedy-snake/lib"
)

// Snake -
type Snake struct {
	loc         [Width * Height]Location //each part location
	size, score int                      //length and score
	dir         byte                     //direction
}

func (s *Snake) init() {
	rand.Seed(time.Now().UnixNano())

	s.size = 2
	s.loc[0].Set(Width/2, Height/2)
	s.loc[1].Set(Width/2-1, Height/2)

	// init direction
	s.dir = 'R'

	// init area
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

	// init food
	food = Location{}
	food.Set(rand.Intn(Width), rand.Intn(Height))
	DrawUI(food, '#')

	go func() {
		for {
			switch lib.Direct() {
			case 72, 87, 119: // up
				s.dir = 'U'
			case 65, 97, 75: // left
				s.dir = 'L'
			case 100, 68, 77: // right
				s.dir = 'R'
			case 83, 115, 80: // down
				s.dir = 'D'
			case 32: // pause
				s.dir = 'P'
			}
		}
	}()
}

// Start -
func (s *Snake) Start() int {
	s.init()
	for {
		// refresh time, the same as speed
		time.Sleep(time.Second / 3)

		//pause
		if s.dir == 'P' {
			continue
		}

		// snake head hit the wall
		if s.loc[0].GetX() < 0 || s.loc[0].GetX() >= Width || s.loc[0].GetY() < 0 || s.loc[0].GetY() >= Height {
			lib.GotoXY(0, 23)
			return s.score
		}

		// snake head hit itself
		for i := 1; i < s.size; i++ {
			x0, y0 := s.loc[0].Get()
			xi, yi := s.loc[i].Get()
			if x0 == xi && y0 == yi {
				lib.GotoXY(0, 23)
				return s.score
			}
		}

		// hit the food
		if s.loc[0].GetX() == food.GetX() && s.loc[0].GetY() == food.GetY() {
			s.size++
			s.score++
			food.Set(rand.Intn(Width), rand.Intn(Height))
			DrawUI(food, '#')
		}

		lp := s.loc[s.size-1]
		for i := s.size - 1; i > 0; i-- {
			s.loc[i] = s.loc[i-1]
			DrawUI(s.loc[i], '*')
		}
		DrawUI(lp, ' ')

		// update snake head
		switch s.dir {
		case 'U':
			s.loc[0].Up()
		case 'L':
			s.loc[0].Left()
		case 'R':
			s.loc[0].Right()
		case 'D':
			s.loc[0].Down()
		}
		DrawUI(s.loc[0], '@')
	}

	return s.score
}
