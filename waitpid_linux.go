package dkinit
// +build linux

// #include "waitpid_linux.h"
import "C"
//import "log"

func Regpid() {
	C.regpid()
}

func waitanypid() int {
	return int(C.wait_any_pid())
}
