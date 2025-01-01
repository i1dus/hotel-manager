package start_handler

import tele "gopkg.in/telebot.v4"

type MenuWrapper struct {
	Menu    *tele.ReplyMarkup
	BtnHelp *tele.Btn
}

func NewMenuWrapper() *MenuWrapper {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnHelp := menu.Text("ℹ Помощь по командам")
	menu.Reply(
		menu.Row(btnHelp),
	)

	return &MenuWrapper{
		Menu:    menu,
		BtnHelp: &btnHelp,
	}
}
