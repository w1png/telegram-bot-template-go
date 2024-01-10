package callbacks

import (
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/errors"
	"github.com/w1png/telegram-bot-template/maps"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func BackCallback(data types.CallbackData) types.UpdateHandlerFunction {
	var state_back_data types.StateBackData
	if state_back_data_map, ok := data.Data.(map[string]any); ok {
		state_back_data = types.StateBackData{
			Data:        state_back_data_map["data"],
			Destination: state_back_data_map["dest"].(string),
		}
	} else {
		if state_back_data, ok = data.Data.(types.StateBackData); !ok {
			return nil
		}
	}

	f, ok := maps.CallbackMap.GetCallback(state_back_data.Destination)
	if !ok {
		return nil
	}

	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		if strings.HasPrefix(state_back_data.Destination, "_admin") && !user.IsAdmin() {
			return nil, errors.NewUnauthorizedError()
		}

		return f(*types.NewCallbackData(
			state_back_data.Destination,
			state_back_data.Data,
		))(bot, update, user)
	}
}
