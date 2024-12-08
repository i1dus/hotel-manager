package help_usecase

import (
	tele "gopkg.in/telebot.v4"
	"strings"
)

type HelpUseCase struct{}

func NewHelpUseCase() *HelpUseCase {
	return &HelpUseCase{}
}

func (uc *HelpUseCase) Help(c tele.Context) error {
	msg := strings.Builder{}
	msg.WriteString("Доступные команды:\n")
	msg.WriteString("/add_employee [TG-логин] [Позиция] [Имя] — добавить работника\n")
	msg.WriteString("/remove_employee [TG-логин] — удалить работника\n")
	msg.WriteString("/employees — вывести всех работников\n")
	msg.WriteString("/add_client [Имя] [Фамилия] [Номер паспорта] — добавить клиента\n")
	msg.WriteString("/add_room [Номер] [Категория] [Цена] — добавить комнату\n")
	msg.WriteString("/rooms — вывести все комнаты\n")
	msg.WriteString("/change_room_price [Номер] [Новая цена за сутки] — изменить цену за сутки\n")
	msg.WriteString("/clean_room [Номер] — выставить уборку в комнату\n")
	msg.WriteString("/room_cleaned [Номер] — уборка в комнате выполнена\n")
	msg.WriteString("/add_room [Номер] [Категория] [Цена] — добавить комнату\n")
	msg.WriteString("/add_occupancy [Номер] [Паспорт клиента] [Конец занятости (DD-MM-YYYY)] [Опционально: описание] — добавить занятость\n")
	msg.WriteString("/end_occupancy [ID занятости] — завершить занятость\n")
	msg.WriteString("/occupancies — вывести все занятости\n")
	msg.WriteString("/stats — статистика по занятости номеров")

	return c.Send(msg.String())
}
