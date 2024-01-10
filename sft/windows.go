// go:build windows
// build linux

package main

import "golang.org/x/sys/windows"

func DiskSpaceLeft() uint64 {
	var freeBytes, totalBytes, totalFreeBytes uint64
	_ := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr("C:"), &freeBytes, &totalBytes, &totalFreeBytes)
	return freeBytes
}
