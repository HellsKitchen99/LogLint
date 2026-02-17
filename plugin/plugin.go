package main

import (
	loglint "selectel"

	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin = []*analysis.Analyzer{
	loglint.Analyzer,
}
