package main

import (
	"flag"
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

const SystemExpirationDateAddress uintptr = 0x7ffe02c8

// GetExpirationTime retrieves the running system's timebomb expiration. It works by retrieving the SystemExpirationDate value as a FileTime from KUSER_SHARED_DATA with some pointer magic.
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
	emulateWinver := flag.Bool("w", false, "Mimic a winver dialog")
	flag.Parse()
	if *emulateWinver {
		winver()
		return
	}
	buildLab := ""

	// get the correct buildlab for CU builds - if this key does not exist just fall back
	cv, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\Software\Microsoft\BuildLayers\OSClient`, registry.QUERY_VALUE)
	if err == nil {
		buildLab, _, err = cv.GetStringValue("BuildLab")
		if err != nil {
			buildLab = fmt.Sprintf("(unable to retrieve OSClient BuildLab value due to '%s')", err)
		}
	} else {
		cv, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
		if err != nil {
			fmt.Println(err)
		}

		buildLab, _, err = cv.GetStringValue("BuildLab")
		if err != nil {
			buildLab = fmt.Sprintf("(unable to retrieve CurrentVersion BuildLab value due to '%s')", err)
		}
	}

	// ensure we have the correct build number for any enablement package builds (e.g. 26220)
	cv, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println(err)
	}

	buildNumber, _, err := cv.GetStringValue("CurrentBuild")
	if err != nil {
		buildNumber = fmt.Sprintf("(unable to retrieve CurrentVersion CurrentBuild value due to '%s')", err)
	}

	buildLabSplit := strings.Split(buildLab, ".")
	if buildLabSplit[0] != buildNumber {
		// this is an enablement package build
		buildLab = buildNumber + "." + strings.Join(buildLabSplit[1:], ".")
	}

	expirationTime := GetExpirationTime()
	switch {
	case expirationTime.IsZero():
		fmt.Printf("Build %s will not expire\n", buildLab)
	default:
		fmt.Printf("Build %s will expire on: %s\n", buildLab, expirationTime.Format("2006/01/02 15:04:05 MST (Monday)"))
	}
}
