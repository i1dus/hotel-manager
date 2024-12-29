package help_usecase

import (
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
)

type HelpUseCase struct{}

func NewHelpUseCase() *HelpUseCase {
	return &HelpUseCase{}
}

func (uc *HelpUseCase) Help(c tele.Context) error {
	return c.Send(domain.HelpMsg)
}
