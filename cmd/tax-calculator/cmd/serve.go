package cmd

import (
	"net/http"
	"os"
	"strconv"
	"time"

	clicommon "github.com/alireza-mht/tax-calculator/cmd/tax-calculator/cmd/common"
	"github.com/alireza-mht/tax-calculator/internal/log"
	"github.com/alireza-mht/tax-calculator/internal/server"
	"github.com/alireza-mht/tax-calculator/internal/server/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"srv"},
	Short:   "Serve REST API",
	Long:    "Run commands .",
	Run: func(cmd *cobra.Command, args []string) {
		clicommon.ServerMode = true
		serve(cmd, args)
		os.Exit(1)
	},
}

func serve(_ *cobra.Command, _ []string) {
	log.InitLogger(clicommon.LogLevel)

	// Create the router
	router := gin.Default()
	apiGroupMainV1 := router.Group("/v1")

	// Create the server
	server := server.NewServer()

	// Register the handler
	api.RegisterHandlersWithOptions(apiGroupMainV1, server, api.GinServerOptions{})

	serveAddress := clicommon.Host + ":" + strconv.Itoa(clicommon.RestPort)
	httpServer := &http.Server{
		Handler:           router,
		Addr:              serveAddress,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Info("Serving HTTP REST API on " + serveAddress)
	log.Info(httpServer.ListenAndServe().Error())
}

func init() {
	serveCmd.PersistentFlags().IntVarP(
		&clicommon.RestPort,
		"port", "p",
		8383,
		"Set REST interface port (default: 8383)",
	)

	serveCmd.Flags().StringVarP(
		&clicommon.Host,
		"host", "",
		"localhost",
		"Set server host address (default: localhost)",
	)
	RootCmd.AddCommand(serveCmd)
}
