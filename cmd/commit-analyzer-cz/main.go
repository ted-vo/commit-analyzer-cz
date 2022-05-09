package main

import (
	defaultAnalyzer "github.com/ted-vo/commit-analyzer-cz/pkg/analyzer"
	"github.com/ted-vo/semantic-release/v3/pkg/analyzer"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		CommitAnalyzer: func() analyzer.CommitAnalyzer {
			return &defaultAnalyzer.DefaultCommitAnalyzer{}
		},
	})
}
