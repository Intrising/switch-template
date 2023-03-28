package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	utilsLog "github.com/Intrising/intri-utils/log"
	"github.com/Intrising/switch-template/handlers"
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
