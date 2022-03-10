package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Intrising/switch-template/hal"
	_ "github.com/Intrising/switch-template/handler"
)

func main() {
	go hal.OrderInit()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	close(c)

	hal.Cancel()
	time.Sleep(50 * time.Millisecond)
}
