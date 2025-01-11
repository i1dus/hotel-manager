package start_handler

import (
	tele "gopkg.in/telebot.v4"
)

type HelpUseCase interface {
	Help(c tele.Context) error
}

type StartUseCase interface {
	Start(c tele.Context) error
}

// StartHandler is responsible to establish connection between bot & user
type StartHandler struct {
	bot          *tele.Bot
	menuWrapper  *MenuWrapper
	helpUseCase  HelpUseCase
	startUseCase StartUseCase
}

func NewStartHandler(bot *tele.Bot, menuWrapper *MenuWrapper, helpUseCase HelpUseCase, startUseCase StartUseCase) *StartHandler {
	return &StartHandler{bot: bot, menuWrapper: menuWrapper, helpUseCase: helpUseCase, startUseCase: startUseCase}
}

func (h *StartHandler) RegisterHandlers() {
	h.bot.Handle("/start", h.startUseCase.Start)
	h.bot.Handle(h.menuWrapper.BtnHelp, h.helpUseCase.Help)
}
