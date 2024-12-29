package handler

import tele "gopkg.in/telebot.v4"

var (
	// Universal markup builders.
	menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector = &tele.ReplyMarkup{}

	// Reply buttons.
	btnHelp = menu.Text("ℹ Help")
)
