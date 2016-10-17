package utils

import (
	"fmt"
)

// CheckError a convenience for checking error.
func CheckError(e error, file string) {
	if e != nil {
		panic(fmt.Sprintf("Got error in file: %s, error: %s", file, e))
	}
}
