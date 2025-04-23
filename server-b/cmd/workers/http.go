package workers

import (
	"context"

	"server/internal/gateway/server"
)

func runHttpWorker(ctx context.Context) {
	srv := server.NewHttpServer(ctx, cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}
