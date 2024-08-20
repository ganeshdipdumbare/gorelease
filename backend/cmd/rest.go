/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"

	"backend/config"
	"backend/internal/api/rest"
	"backend/internal/app"
	"backend/internal/gocodeplayer"

	"github.com/spf13/cobra"
)

var port int

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A command to start rest api server",
	Long: `A command to start rest api server at given port. 
	If the port is not provided, it will start the server at default port 8080.`,
	Run: func(cmd *cobra.Command, args []string) {
		// create a slog logger instance
		opts := &slog.HandlerOptions{
			AddSource: true,
		}
		switch config.Get().LogLevel {
		case "debug":
			opts.Level = slog.LevelDebug
		case "info":
			opts.Level = slog.LevelInfo
		case "warn":
			opts.Level = slog.LevelWarn
		case "error":
			opts.Level = slog.LevelError
		}
		slogJsonHandler := slog.NewJSONHandler(os.Stdout, opts)
		logger := slog.New(slogJsonHandler)

		// create new gocodeplayer instance
		goCodePlayer, err := gocodeplayer.NewPlaygroundClient(config.Get().CodePlayerURL, gocodeplayer.WithLogger(logger))
		if err != nil {
			logger.Error("failed to create new gocodeplayer instance", slog.Any("error", err))
			os.Exit(1)
		}

		// create new app instance
		codePlayerApp := app.NewCodePlayer(goCodePlayer)

		// create a new rest api instance
		api, err := rest.NewApi(codePlayerApp, config.Get().Port)
		if err != nil {
			logger.Error("failed to create new rest api", slog.Any("error", err))
			os.Exit(1)
		}

		// start the rest server
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(restCmd)

	restCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
}
