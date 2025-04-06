// Package cmd is used for cli commands
package cmd

import (
	clicommon "github.com/alireza-mht/tax-calculator/cmd/tax-calculator/cmd/common"
	"github.com/spf13/cobra"
)

// RootCmd is the main application command.
var RootCmd = &cobra.Command{
	Use:   "tc",
	Short: "TC - Tax calculator.",
	Long: `TC - Tax calculator.

	Allows calculating the tax.`,
}

// init configures the main Cobra command's flags.
func init() {
	RootCmd.PersistentFlags().StringVarP(
		&clicommon.LogLevel,
		"log-level", "l",
		"info",
		"Set log level. Choices: d[ebug], i[nfo], n[otice], w[arning], e[rror]",
	)
}
