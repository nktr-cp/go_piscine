package main

import "ft"

const CLOSE = 0
const OPEN = 1

type Door struct {
	state int
}

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}
func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == OPEN
}
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
