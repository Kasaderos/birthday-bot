/*
	todo docker command
*/

package test

import (
	"birthday-bot/config"
	dbPg "birthday-bot/internal/adapters/db/pg"
	"birthday-bot/internal/adapters/logger/zap"
	notifier "birthday-bot/internal/adapters/notifier/tgbot"
	repoPg "birthday-bot/internal/adapters/repo/pg"
	"birthday-bot/internal/adapters/server"
	"birthday-bot/internal/adapters/server/rest"
	"birthday-bot/internal/domain/core"
	"birthday-bot/internal/domain/entities"
	"birthday-bot/internal/domain/usecases"
	"birthday-bot/pkg/clock"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	dopDbPg "github.com/rendau/dop/adapters/db/pg"
)

var app *TestApp

type TestApp struct {
	lg         *zap.St
	db         *dopDbPg.St
	dbRaw      *dbPg.St
	repo       *repoPg.St
	core       *core.St
	srv        *server.St
	ucs        *usecases.St
	notifier   *notifier.St
	cfg        *config.ConfSt
	testChatID int64
}

func TestBirthdayBot(t *testing.T) {
	setup(t)
	defer stop(t)

	// wait server
	time.Sleep(time.Second * 3)
	testUsersCRUD(t)
	// t.Run("notify user", notifyUser)
}

func notifyUser(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	start := time.Now()
	sendInterval := clock.TimeInterval{
		Start: start,
		End:   start.Add(time.Hour),
	}
	app.core.Start(ctx, sendInterval)
}

func testUsersCRUD(t *testing.T) {
	serverURL := "http://localhost" + app.cfg.HttpListen + "/"

	users := []entities.UserSt{
		{0, "Nate", "River", "1998-08-25", app.testChatID},
		{1, "Alice", "River", "1998-08-25", app.testChatID},
	}
	for _, user := range users {
		// post user
		body, _ := json.Marshal(user)
		resp, err := http.Post(serverURL+"users/", "application/json", bytes.NewBuffer(body))
		assert(t, err == nil, "request post", err)

		defer resp.Body.Close()

		// check id
		respJSON := struct {
			ID int64 `json:"id"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&respJSON)
		assert(t, err == nil, "resp json")

		// update id
		user.ID = respJSON.ID

		// get user
		resp, err = http.Get(fmt.Sprintf("%s/users/%d", serverURL, respJSON.ID))
		assert(t, err == nil, "request get")
		defer resp.Body.Close()

		// check user
		var respUserJson entities.UserSt
		err = json.NewDecoder(resp.Body).Decode(&respUserJson)
		assert(t, err == nil, "resp json")

		assert(t, respUserJson.Birthday == user.Birthday, "birthay doesn't match ")
		assert(t, respUserJson.TelegramChatID == user.TelegramChatID, "chat id doesn't match")

		// update birthday
		user.Birthday = time.Now().Format(time.DateOnly)

		body, _ = json.Marshal(user)
		req, err := http.NewRequest("PUT", serverURL+"users/", bytes.NewBuffer(body))
		assert(t, err == nil, "request put")

		updateResp, err := http.DefaultClient.Do(req)
		assert(t, err == nil, "request put")
		defer updateResp.Body.Close()

		err = json.NewDecoder(updateResp.Body).Decode(&respUserJson)
		assert(t, err == nil, "resp json")

		// check user
		assert(t, respUserJson.Birthday == user.Birthday, "updated birthay doesn't match")
	}
}

func setup(t *testing.T) {
	var err error

	app = &TestApp{}

	// load config
	conf := config.Load()
	app.cfg = conf
	log.Println(conf)

	// logger
	app.lg = zap.New(conf.LogLevel, conf.Debug)

	// db
	app.db, err = dopDbPg.New(conf.Debug, app.lg, dopDbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		t.Fatal("dopdbpg", err)
	}

	// dbRaw
	app.dbRaw, err = dbPg.New(conf.Debug, app.lg, dbPg.OptionsSt{
		Dsn: conf.PgDsn,
	})
	if err != nil {
		t.Fatal("dbpg", err)
	}

	// repo
	app.repo = repoPg.New(app.lg, app.db, app.dbRaw)

	app.notifier, err = notifier.New(app.lg, conf.BotToken)
	if err != nil {
		t.Fatal("notifier", err)
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

}

func stop(t *testing.T) {
	var exitCode int

	app.lg.Infow("Shutting down...")

	if !app.srv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Wait routines...")

	app.core.StopAndWaitJobs()

	app.lg.Infow("Exit")

	if exitCode > 0 {
		t.Fatalf("exited with code %d", exitCode)
	}
}

func stopSignal() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	return ch
}

func assert(t *testing.T, ok bool, msgs ...interface{}) {
	if !ok {
		t.Fatal(msgs...)
	}
}
