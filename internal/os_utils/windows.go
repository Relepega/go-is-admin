//go:build windows

package osutils

import (
	"golang.org/x/sys/windows"
)

func IsAdmin() bool {
	elevated := windows.GetCurrentProcessToken().IsElevated()
	return elevated
}
