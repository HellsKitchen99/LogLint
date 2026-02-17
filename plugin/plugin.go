package main

import (
	loglint "github.com/HellsKitchen99/LogLint"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		loglint.Analyzer,
	}, nil
}
