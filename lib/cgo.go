package lib

/*
#include <windows.h>
#include <conio.h>

void gotoxy(int x,int y) {
    COORD c;
    c.X=x,c.Y=y;
    SetConsoleCursorPosition(GetStdHandle(STD_OUTPUT_HANDLE),c);
}

int direct() {
    return _getch();
}
*/
import "C"

func GotoXY(x, y int) {
	C.gotoxy(C.int(x), C.int(y))
}

func Direct() int {
	return int(C.direct())
}
