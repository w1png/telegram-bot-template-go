# English

## Language

### Creating a language string
Open `language/language_string.go` file. Add your string to the list of constants.

```go
package language

type LanguageString string

const (
	...
    Greet LanguageString = "Greet"
    ...
)
```

Open all the language files i.e. `language/language_en.go` and add your string to the map.

```go
package language

import "github.com/w1png/telegram-bot-template/errors"

type English struct {
	Values map[LanguageString]string
}

func NewEnglish() *English {
	return &English{
		Values: map[LanguageString]string{
			...
            Greet: "Hello!",
            ...
		},
	}
}

func (e *English) Get(key LanguageString) (string, error) {
	value, ok := e.Values[key]
	if !ok {
		return "", errors.NewLanguageStringError(string(key))
	}
	return value, nil
}
```

## Getting your language string

```go
import "github.com/w1png/telegram-bot-template/language"

func printGreet() {
	s, err := language.LanguageInstance.Get(language.Greet)
	if err != nil {
		return log.Fatal(err)
	}
    log.Println(s)
}
```

## Creating a command handler

Create a file in `handlers/commands/` folder i.e. `handlers/commands/greet.go`.

```go
package commands

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func GreetCommand() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
        s, err := language.LanguageInstance.Get(language.Greet)
        if err != nil {
            return nil, err
        }

		return tg.NewMessage(
			update.FromChat().ID,
            s,
		), nil
    }
}
```

Open `handlers/commands/commands.go` and add your command to the map.

```go
package commands

import "github.com/w1png/telegram-bot-template/types"

var CommandsMap = map[string]func() types.UpdateHandlerFunction{
    "greet": GreetCommand,
}
```

You can now use your /greet command.

![greet](https://github.com/w1png/telegram-bot-template-go/assets/74238629/d47c7c40-46ae-4cf9-b4b2-c485a4c9f3cc)

# Creating a message handler
Message handlers are useful for handling keyboard button inputs.

For example, let's create a `/start` command that will send a greeting message with a keyboard.
```go
func StartCommand() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		text, err := language.LanguageInstance.Get(language.Start)
		if err != nil {
			return tg.MessageConfig{}, err
		}

		replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
		replyMsg.ReplyToMessageID = update.Message.MessageID
		replyMsg.ReplyMarkup = tg.NewReplyKeyboard(
			tg.NewKeyboardButtonRow(
				tg.NewKeyboardButton("Help"),
			),
		)

		return replyMsg, nil
	}
}
```

![help_start](https://github.com/w1png/telegram-bot-template-go/assets/74238629/7debd3b8-08c3-401f-9044-fc66ccf15248)


Create a file in `handlers/messages/` folder i.e. `handlers/messages/help.go`.
```go
package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/types"
)

func HelpMessage() types.UpdateHandlerFunction {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update, user *models.User) (tgbotapi.Chattable, error) {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "This is a help message"), nil
	}
}
```

Open `handlers/messages/messages.go` and add your message handler to the map.

```go
package messages

import "github.com/w1png/telegram-bot-template/types"

var MessagesMap = map[string]func() types.UpdateHandlerFunction{
    ...
    "help": HelpMessage,
    ...
}
```

You can now use your `Help` button.

![help](https://github.com/w1png/telegram-bot-template-go/assets/74238629/2373324f-34c0-49cb-8f49-c6d0e54d96da)


# Callbacks

## Marshaling callback data
Callback data has a call (a key in the CallbackMap) and data that can be anything.
```go
type CallbackData struct {
	Call string `json:"c"`
	Data any    `json:"d"`
}
```

If you, for example, you are making an ecommerce bot and you want to make a button that opens the category with id 1.
```go
...
callback_data, err := types.NewCallbackData("category", "1").Marshal()
...
```

The button callback data will look like this: `{"c":"category","cid":1}`

You can pass any JSON serializable data to the callback data. This can be usefull for creating a `back` button.

```go
...
callback_data, err := types.NewCallbackData(
  "category",
  struct {
    CategoryId uint `json:"cid"`
    OpenedFromCategoryId  uint `json:"from"`
  }{
    CategoryId: 1,
    OpenedFromCategoryId: 2,
  },
).Marshal()
...
```
The button callback data will look like this: `{"c":"category",{"cid":1,"from":2}}`

It's a great idea to use short strings for keys and data because telegram has a character limit for button callback data.


## Creating a callback
For example, let's create a `/start` command that will send a greeting message with an inline keyboard.
```go
func StartCommand() types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		text, err := language.LanguageInstance.Get(language.Start)
		if err != nil {
			return tg.MessageConfig{}, err
		}

		callback_data, err := types.NewCallbackData("category", "1").Marshal()
		if err != nil {
			return nil, err
		}

		replyMsg := tg.NewMessage(update.Message.Chat.ID, text)
		replyMsg.ReplyToMessageID = update.Message.MessageID
		replyMsg.ReplyMarkup = tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(
				tg.NewInlineKeyboardButtonData("Open category", string(callback_data)),
			),
		)

		return replyMsg, nil
	}
}
```

![2024-01-08 2 32 16 PM](https://github.com/w1png/telegram-bot-template-go/assets/74238629/4f2e29c3-93c0-4fe5-8258-ae7499f4daf2)

Create a file in `handlers/callbacks/` folder i.e. `handlers/callbacks/category.go`.
```go
package callbacks

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/w1png/telegram-bot-template/language"
	"github.com/w1png/telegram-bot-template/models"
	"github.com/w1png/telegram-bot-template/storage"
	"github.com/w1png/telegram-bot-template/types"
)

func CategoryCallback(data types.CallbackData) types.UpdateHandlerFunction {
	return func(bot *tg.BotAPI, update tg.Update, user *models.User) (tg.Chattable, error) {
		category_id, ok := data.Data.(uint)
		if !ok {
			return nil, fmt.Errorf("category_id is not string")
		}

		category, err := storage.StorageInstance.GetCategoryById(category_id)
		if err != nil {
			return nil, err
		}

		t, err := language.LanguageInstance.Get(language.CategoryFormat)
		if err != nil {
			return nil, err
		}

		return tg.NewEditMessageText(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			fmt.Sprintf(t, category.Name),
		), nil
	}
}
```

Open `handlers/callbacks/callbacks.go` and add your callback to the map.

```go
package callbacks

import "github.com/w1png/telegram-bot-template/types"

var CallbackMap = map[string]func(types.CallbackData) types.UpdateHandlerFunction{
    "category": CategoryCallback,
}
```

You can now use your `Open category` button.

![2024-01-08 2 45 46 PM](https://github.com/w1png/telegram-bot-template-go/assets/74238629/b945a629-54dd-4948-8ff7-5e6b05845dcf)

# Environment variables
Environment variables can be added to config/config.go.

| **Name**       | **Description**                                                             |
|----------------|-----------------------------------------------------------------------------|
| TELEGRAM_TOKEN | Your Telegram bot token                                                     |
| IS_DEBUG       | Sets tgbotapi.Bot.Debug=true if IS_DEBUG=true                               |
| LANGUAGE       | Sets the language. Possible values are `en` or `ru`                         |
| STORAGE_TYPE   | Database type. Possible value is `sqlite`                                   |
| SQLITE_PATH    | A path to your sqlite database. Example: `data.db`                          |
| LOGGER_TYPE    | Logger type. Possible value is `console`. Defaults to no logger if not set. |

# Development
To start the bot in live reload mode, run the `air` command.

# Building and running
To build the docker container, run `make build-docker`.
To the run the container, run `make run-docker`
