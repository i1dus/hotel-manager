package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"hotel-management/internal/handlers"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	appCtx := context.Background()

	conn, err := pgx.Connect(appCtx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close(appCtx)

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = handlers.RegisterHandlers(b)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
