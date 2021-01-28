package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	var err error
	err = &os.PathError{Err: syscall.ENOENT}
	switch err.(type) {
	case *os.PathError:
		//do nothing
	default:
		fmt.Println("somethingelse")
	}
}
