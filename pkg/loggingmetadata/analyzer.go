package loggingmetadata

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loggingmetadata",
	Doc:  "checks that all logging calls include metadata fields",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		// Look for function calls
		callExpr, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		// Check if it's a selector expression (e.g., log.Info)
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		// Check if the base is an identifier
		ident, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return true
		}

		// Check if it's a logging call
		if ident.Name == "log" {
			methodName := selectorExpr.Sel.Name
			if isLoggingMethod(methodName) {
				// Check if there's only one argument (just the message)
				if len(callExpr.Args) <= 1 {
					pass.Reportf(
						callExpr.Pos(),
						"logging calls must include metadata fields, not just a message",
					)
				}
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}

	return nil, nil
}

// isLoggingMethod returns true if the method name is a logging method
func isLoggingMethod(name string) bool {
	loggingMethods := []string{
		"Debug", "Info", "Warn", "Error", "Fatal",
		"Debugf", "Infof", "Warnf", "Errorf", "Fatalf",
	}

	for _, method := range loggingMethods {
		if strings.EqualFold(name, method) {
			return true
		}
	}
	return false
}
