package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"hotel-management/internal/handler"
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

	botSettings := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(botSettings)
	if err != nil {
		log.Fatal(err)
		return
	}

	handlerController := handler.NewHandlerController(b, conn)
	handlerController.RegisterHandlers()
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Start()
}
