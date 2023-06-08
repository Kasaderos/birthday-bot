package main

import (
	"birthday-bot/config"
	dbPg "birthday-bot/internal/adapters/db/pg"
	"birthday-bot/internal/adapters/logger/zap"
	notifier "birthday-bot/internal/adapters/notifier/tgbot"
	repoPg "birthday-bot/internal/adapters/repo/pg"
	"birthday-bot/internal/adapters/server"
	"birthday-bot/internal/adapters/server/rest"
	"birthday-bot/internal/domain/core"
	"birthday-bot/internal/domain/usecases"
	"birthday-bot/pkg/clock"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	dopDbPg "github.com/rendau/dop/adapters/db/pg"
)

func main() {
	var err error

	app := struct {
		lg       *zap.St
		db       *dopDbPg.St
		dbRaw    *dbPg.St
		repo     *repoPg.St
		core     *core.St
		srv      *server.St
		ucs      *usecases.St
		notifier *notifier.St
	}{}

	// load config
	conf := config.Load()

	// logger
	app.lg = zap.New(conf.LogLevel, conf.Debug)

	// db
	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// dbRaw
	app.dbRaw, err = dbPg.New(conf.Debug, app.lg, dbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		app.lg.Fatal(err)
	}

	// repo
	app.repo = repoPg.New(app.lg, app.db, app.dbRaw)

	app.notifier, err = notifier.New(app.lg, conf.BotToken)
	if err != nil {
		app.lg.Fatal(err)
	}

	// core
	app.core = core.New(
		app.lg,
		app.repo,
		app.notifier,
	)

	// usecases
	app.ucs = usecases.New(app.lg, app.dbRaw, app.core)

	// START

	app.lg.Infow("Starting")

	app.srv = server.Start(
		app.lg,
		conf.HttpListen,
		rest.GetHandler(app.lg, app.ucs, conf.HttpCors),
	)

	ctx, cancelCtx := context.WithCancel(context.Background())
	timeInterval, err := clock.NewTimeInterval(conf.NotifyInterval)
	if err != nil {
		app.lg.Fatal(err)
	}
	go app.core.Start(ctx, timeInterval)

	time.Now().Clock()
	// LISTEN

	var exitCode int

	select {
	case <-stopSignal():
	case <-app.srv.Wait():
		exitCode = 1
	}

	// STOP

	cancelCtx()
	app.lg.Infow("Shutting down...")

	if !app.srv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Wait routines...")

	app.core.StopAndWaitJobs()

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}

func stopSignal() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	return ch
}
