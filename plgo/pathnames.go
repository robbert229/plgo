// +build !windows

package main

import (
	"io/ioutil"
	"strings"
)

// getCorrectPath used on Windows, see file pathnames_windows.go
func getCorrectPath(p string) string {
	ret := strings.TrimRight(p, "\n")
	return ret
}

// addOtherIncludesAndLDFLAGS used on Windows, see file pathnames_windows.go
func addOtherIncludesAndLDFLAGS(plgoSource *string, postgresIncludeDir string) {
	return
}

func buildPath() (string, error) {
	return ioutil.TempDir("", plgo)
}
