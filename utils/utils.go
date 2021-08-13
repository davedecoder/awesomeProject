package utils

import (
	"fmt"
	"os"
)

func SafeFileClose(file *os.File) {
	if err := file.Close(); err != nil {
		println(fmt.Sprintln("Error Closing File", err))
	}
}
