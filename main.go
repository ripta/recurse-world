package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ripta/recurse-world/pkg/server"
	"go.uber.org/zap"
)

var (
	port         = flag.Int("port", 8080, "the port number to listen on")
	serverName   = flag.String("server-name", "", "the name of the server (default: hostname)")
	upstreamHost = flag.String("upstream", "localhost:8080", "the upstream to use")
	withName     = flag.Bool("with-name", false, "include server name in response")
	withTime     = flag.Bool("with-time", false, "include time in response")
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
		*serverName, _ = os.Hostname()
	}

	s := &server.Server{
		Logger:       logger,
		Name:         *serverName,
		UpstreamHost: *upstreamHost,
		WithName:     *withName,
		WithTime:     *withTime,
	}
	if err := s.Serve(":" + strconv.Itoa(*port)); err != nil {
		logger.Error("server ended", zap.Error(err))
	}
}
