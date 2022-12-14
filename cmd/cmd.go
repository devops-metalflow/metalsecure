package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/devops-metalflow/metalsecure/config"
	"github.com/devops-metalflow/metalsecure/server"
)

var (
	app       = kingpin.New("metalsecure", "metalsecure").Version(config.Version + "-build-" + config.Build)
	listenUrl = app.Flag("listen-url", "Listen URL (host:port)").Required().String()
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c, err := initConfig(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	s, err := initServer(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init server")
	}

	if err := runServer(ctx, s); err != nil {
		return errors.Wrap(err, "failed to run server")
	}

	return nil
}

func initConfig(_ context.Context) (*config.Config, error) {
	c := config.New()
	return c, nil
}

func initServer(ctx context.Context, _ *config.Config) (server.Server, error) {
	c := server.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Addr = *listenUrl

	return server.New(ctx, c), nil
}

func runServer(ctx context.Context, srv server.Server) error {
	if err := srv.Init(ctx); err != nil {
		return errors.New("failed to init")
	}

	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Fatalf("failed to run: %v", err)
		}
	}()

	s := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		<-s
		_ = srv.Deinit(ctx)
		done <- true
	}()

	<-done

	return nil
}
