package main

import (
	"fmt"
	osutils "go-win-admin/internal/os_utils"
)

func main() {
	fmt.Println("hello")

	if osutils.IsAdmin() {
		fmt.Println("The app is running as administrator")
	} else {
		fmt.Println("The app is not running as administrator")
	}

	fmt.Println("\n")
	fmt.Println("Press <ENTER> key to exit...")
	fmt.Scanln()
}
