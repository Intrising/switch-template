package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Intrising/intri-core/handlers"
	utilsLog "github.com/Intrising/intri-utils/log"
)

func main() {
	c := make(chan os.Signal, 1)
	defer close(c)

	defer utilsLog.Info("main done")

	defer time.Sleep(500 * time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	go new(handlers.Handler).Run(ctx)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
