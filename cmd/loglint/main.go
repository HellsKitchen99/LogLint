package main

import (
	loglint "github.com/HellsKitchen99/LogLint"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(loglint.Analyzer)
}
