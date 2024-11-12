package handlers

import (
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
)

type handlerController struct {
	repo *pgx.Conn
}

func (c *handlerController) RegisterHandlers(bot *tele.Bot) error {
	bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	c.repo.Query()
	return nil
}
