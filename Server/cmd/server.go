package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	c "github.com/ViPDanger/L0/Server/Internal/config"
	n "github.com/ViPDanger/L0/Server/Internal/nats"
	pg "github.com/ViPDanger/L0/Server/Internal/postgres"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config := c.Read_Config()
	go func() {
		client, _ := pg.NewClient(context.Background(), config)
		rep := pg.NewRepository(client)
		if rep != nil {
		}
		natsConnection, _ := n.ConnectToNATS(config)
		if natsConnection != nil {
		}
	}()
	<-ctx.Done()
	log.Println("Server was killed.")
}
