package main

import (
	"github.com/jeremyrajan/golangci-linters/pkg/loggingmetadata"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		loggingmetadata.Analyzer,
	}
}

var AnalyzerPlugin analyzerPlugin
