package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/ripta/recurse-world/pkg/server"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var (
	port         = flag.Int("port", 8080, "the port number to listen on")
	serverName   = flag.String("server-name", "", "the name of the server (default: random UUID)")
	upstreamHost = flag.String("upstream", "localhost:8080", "the upstream to use")
)

func main() {
	flag.Parse()
	run() // so defer works
}

func run() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Printf("cannot initialize zap logger: %v\n", err)
		return
	}
	defer logger.Sync()

	if *serverName == "" {
		*serverName = uuid.NewV4().String()
	}

	s := &server.Server{
		Logger:       logger,
		Name:         *serverName,
		UpstreamHost: *upstreamHost,
	}
	if err := s.Serve(":" + strconv.Itoa(*port)); err != nil {
		logger.Error("server ended", zap.Error(err))
	}
}
