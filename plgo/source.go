package main

import (
	"go/build"
	"os"
)

//go:generate go-bindata -pkg main -nometadata -modtime=0 -o ./bindata.go ../pl.go

func readPlGoSource() ([]byte, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH // Go 1.8 and later have a default GOPATH
	}

	assetBytes, err := Asset("../pl.go")
	if err != nil {
		return nil, err
	}

	return assetBytes, nil
}
