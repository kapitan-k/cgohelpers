package cgohelpers

import (
	"syscall"
)

type Errno uintptr

func ErrnoGet(err error) error {
	eno, ok := err.(syscall.Errno)
	if ok {
		return Errno(eno)
	}
	return err
}

func (errno Errno) Error() string {
	return syscall.Errno(errno).Error()
}
