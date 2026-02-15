package loglint

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "finds invalid args in log funs based on 4 rules",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {

}
