// Package for starting the application
package main

import (
	"github.com/alireza-mht/tax-calculator/cmd/tax-calculator/cmd"
	clicommon "github.com/alireza-mht/tax-calculator/cmd/tax-calculator/cmd/common"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		clicommon.Logger.Fatalln(err.Error())
	}
}
