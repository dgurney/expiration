package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func winver() {
	shell32 := windows.NewLazySystemDLL("shell32.dll")
	shellAbout := shell32.NewProc("ShellAboutW")

	szApp, _ := windows.UTF16PtrFromString("Binbows")
	exp := GetExpirationTime()
	szOtherStuff, _ := windows.UTF16PtrFromString(fmt.Sprintf("Evaluation copy. Expires %s", exp.Format("2006/01/02 03:04")))
	if exp.IsZero() {
		szOtherStuff, _ = windows.UTF16PtrFromString("")
	}
	shellAbout.Call(0, uintptr(unsafe.Pointer(szApp)), uintptr(unsafe.Pointer(szOtherStuff)), 0)
}
