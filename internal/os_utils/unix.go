//go:build !windows
// +build !windows

package osutils

import (
	"os"
)

func IsAdmin() bool {
	if os.Geteuid() == 0 {
		return true
	}

	return false
}
