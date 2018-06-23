package conio

import (
	"sync"
	"syscall"
)

var conOut Handle
var conOutOnce sync.Once

// ConOut returns the handle for Console-Output
func ConOut() Handle {
	conOutOnce.Do(func() {
		var err error
		conOut, err = syscall.Open("CONOUT$", syscall.O_RDWR, 0)
		if err != nil {
			panic(err.Error())
		}
	})
	return conOut
}
