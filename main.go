package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

const SystemExpirationDateAddress uintptr = 0x7ffe02c8

/* GetExpirationTime retrieves the running system's timebomb expiration. It works by retrieving the SystemExpirationDate value as a FileTime from KUSER_SHARED_DATA with some pointer magic. */
func GetExpirationTime() time.Time {
	// who needs safety and 0 vet warnings anyway
	ptr := unsafe.Pointer(SystemExpirationDateAddress)

	if *((*uint)(ptr)) == 0 {
		// if the value is 0, there is no expiration
		return time.Time{}
	}

	expiration := *((*syscall.Filetime)(ptr))
	return time.Unix(0, expiration.Nanoseconds())
}

func main() {
	cv, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println(err)
	}

	buildLabEx, _, err := cv.GetStringValue("BuildLab")
	if err != nil {
		buildLabEx = fmt.Sprintf("(unable to retrieve BuildLab value due to '%s')", err)
	}

	expirationTime := GetExpirationTime()
	switch {
	case expirationTime.IsZero():
		fmt.Printf("Build %s will not expire\n", buildLabEx)
	default:
		fmt.Printf("Build %s will expire on: %s\n", buildLabEx, expirationTime.Format("2006-01-02 15:04:05 MST (Monday)"))
	}
}
