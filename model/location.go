package model

var food Location

// Location -
type Location struct {
	x, y int
}

// Set -
func (l *Location) Set(x, y int) {
	l.x, l.y = x, y
}

// Get -
func (l *Location) Get() (int, int) {
	return l.x, l.y
}

// GetX -
func (l *Location) GetX() int {
	return l.x
}

// GetY -
func (l *Location) GetY() int {
	return l.y
}

// Up -
func (l *Location) Up() *Location {
	if l.y > 0 {
		l.y--
	}
	return l
}

// Down -
func (l *Location) Down() *Location {
	l.y++
	return l
}

// Left -
func (l *Location) Left() *Location {
	if l.x > 0 {
		l.x--
	}
	return l
}

// Right -
func (l *Location) Right() *Location {
	l.x++
	return l
}
