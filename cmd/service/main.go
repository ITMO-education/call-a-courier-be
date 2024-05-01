package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/itmo-education/delivery-backend/internal/config"
	"github.com/itmo-education/delivery-backend/internal/data/sqlite"
	"github.com/itmo-education/delivery-backend/internal/transport/rest"
	"github.com/itmo-education/delivery-backend/internal/utils/closer"
	//_transport_imports
)

func main() {
	logrus.Println("starting app")

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		logrus.Fatalf("error reading config %s", err.Error())
	}

	if cfg.AppInfo().StartupDuration == 0 {
		logrus.Fatalf("no startup duration in config")
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.AppInfo().StartupDuration)
	closer.Add(func() error {
		cancel()
		return nil
	})

	sqliteDbConf, err := cfg.Resources().Sqlite(config.Sqlite)
	if err != nil {
		logrus.Fatalf("error loading sqlite database %s", err.Error())
	}
	db, err := sqlite.New(sqliteDbConf)
	if err != nil {
		logrus.Fatalf("can't open database: %s", err.Error())
	}

	restAPI, err := cfg.Api().REST(config.ApiRest)
	if err != nil {
		logrus.Fatalf("error getting rest config %s", err)
	}

	srv := rest.NewServer(cfg, restAPI, db)
	err = srv.Start(ctx)
	if err != nil {
		logrus.Fatalf("error starting web server %s", err)
	}

	waitingForTheEnd()

	logrus.Println("shutting down the app")

	if err = closer.Close(); err != nil {
		logrus.Fatalf("errors while shutting down application %s", err.Error())
	}
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is a LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
