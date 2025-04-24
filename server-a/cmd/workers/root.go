package workers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"server/internal/gateway/provider"
	"server/pkg/config"
	"server/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

func cleanup() {
	logger.Log.Info("Running cleanup...")

	if provider.GRPCUserClient != nil {
		provider.GRPCUserClient.Close()
		logger.Log.Info("gRPC client connection closed")
	}

	logger.Log.Info("Cleanup completed")
}

func Start() {
	cfg = config.InitConfig()
	logger.SetZerologLogger(cfg)
	provider.InitProvider(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	defer cleanup()
	rootCmd := &cobra.Command{}
	cmd := []*cobra.Command{
		{
			Use:   "serve-all",
			Short: "Run all",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(ctx)
			},
			PreRun: func(cmd *cobra.Command, _ []string) {
				go func() {
					runQueueWorker(ctx)
				}()
				go func() {
					runGRPCWorker(ctx)
				}()
			},
		},
		{
			Use:   "serve-http",
			Short: "Run HTTP server",
			Run: func(cmd *cobra.Command, _ []string) {
				runHttpWorker(ctx)
			},
		},
		{
			Use:   "serve-worker",
			Short: "Run worker",
			Run: func(cmd *cobra.Command, _ []string) {
				runQueueWorker(ctx)
			},
		},
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		cleanup()
		log.Fatal(err)
	}
}
