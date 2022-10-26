package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/upbound/upjet/pkg/pipeline"

	"github.com/upbound/provider-gcp/config"
)

TEST 

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	absRootDir, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path of %s", os.Args[1]))
	}
	pipeline.Run(config.GetProvider(), absRootDir)
}
