package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/yosa12978/mockPizza/internal/api"
	"github.com/yosa12978/mockPizza/internal/db"
)

func Run() {
	db.GetDB()
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}
	go api.Listen(listener)
	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	<-out
}
