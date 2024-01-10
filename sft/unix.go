// go:build !windows
//go:build !windows
// +build !windows

package main

import (
	"golang.org/x/sys/unix"
	"os"
)

func DiskSpaceLeft() uint64 {
	var stat unix.Statfs_t
	wd, _ := os.Getwd()
	unix.Statfs(wd, &stat)
	return stat.Bavail * uint64(stat.Bsize)
}
